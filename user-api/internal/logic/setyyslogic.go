package logic

import (
	"context"
	"go-zero-play-1/yys-rpc/types/yys"

	"go-zero-play-1/user-api/internal/svc"
	"go-zero-play-1/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetYysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetYysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetYysLogic {
	return &SetYysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetYysLogic) SetYys(req *types.YysReq) (resp *types.BaseReq, err error) {
	// todo: add your logic here and delete this line
	b, err := l.svcCtx.YysRpc.SetYys(l.ctx, &yys.SetYysRequest{Ordersn: req.Ordersn})
	if err != nil || b == nil {
		return
	}

	resp = &types.BaseReq{Code: int(b.Code)}
	return
}
