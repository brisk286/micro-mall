package logic

import (
	"context"
	"micro-mall-server/service/order/api/internal/svc"
	"micro-mall-server/service/order/api/internal/types"
	"micro-mall-server/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.OrderRpc.Create(l.ctx, &order.CreateRequest{
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}