package logic

import (
	"context"
	"micro-mall-server/common/errx"

	"micro-mall-server/app/user/rpc/internal/svc"
	"micro-mall-server/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNoExistsError = errx.NewErrMsg("用户不存在")

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserInfoResp{}, nil
}
