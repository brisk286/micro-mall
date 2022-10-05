package user

import (
	"context"
	"micro-mall-server/app/user/rpc/userrpc"

	"micro-mall-server/app/user/api/internal/svc"
	"micro-mall-server/app/user/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userrpc.LoginReq{
		AuthType: "",
		AuthKey:  req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	_ = copier.Copy(resp, loginResp)

	return
}
