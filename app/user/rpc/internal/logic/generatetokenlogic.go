package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"micro-mall-server/common/ctxdata"
	"time"

	"micro-mall-server/app/user/rpc/internal/svc"
	"micro-mall-server/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	// 当前时间
	now := time.Now().Unix()
	// 过期时间
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	// 生成token
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", in.UserId, err)
	}

	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {

	// jwt第二部分claims是一个字典，储存各种信息
	claims := make(jwt.MapClaims)
	// 过期时间
	claims["exp"] = iat + seconds
	// 当前时间
	claims["iat"] = iat
	// userid
	claims[ctxdata.CtxKeyJwtUserId] = userId

	// 加密方法
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	// 使用密钥处理得到token
	return token.SignedString([]byte(secretKey))
}
