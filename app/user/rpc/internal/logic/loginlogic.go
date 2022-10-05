package logic

import (
	"context"
	"github.com/pkg/errors"
	"micro-mall-server/app/user/model"
	"micro-mall-server/app/user/rpc/userrpc"
	"micro-mall-server/common/errx"
	"micro-mall-server/common/tool"

	"micro-mall-server/app/user/rpc/internal/svc"
	"micro-mall-server/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = errx.NewErrMsg("生成token失败")
var ErrUsernamePwdError = errx.NewErrMsg("账号或密码不正确")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var userId int64
	var err error

	userId, err = l.loginByMobile(in.AuthKey, in.Password)
	if err != nil {
		return nil, err
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&userrpc.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return &userrpc.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil

}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "mobile:%s", mobile)
	}

	if !(tool.Md5ByString(password) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	return user.Id, nil
}
