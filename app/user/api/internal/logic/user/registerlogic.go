package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"micro-mall-server/app/user/rpc/userrpc"

	"micro-mall-server/app/user/api/internal/svc"
	"micro-mall-server/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &userrpc.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		AuthKey:  req.Mobile,
		AuthType: "",
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	_ = copier.Copy(&resp, registerResp)

	return
}
