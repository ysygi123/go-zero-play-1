package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-play-1/user-api/internal/svc"
	"go-zero-play-1/user-api/internal/types"
	user2 "go-zero-play-1/user-rpc/types/user"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserReply, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user2.IdRequest{
		Id: int32(req.Id),
	})
	if err != nil {
		return
	}

	l.Logger.Infow("新浦斯顿", logx.Field("啊哈", "黑油"))
	resp = &types.UserReply{
		Id:   int(user.Id),
		Name: user.Name,
	}
	return
}
