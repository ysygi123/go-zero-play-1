package logic

import (
	"context"
	user_model "go-zero-play-1/model/mysql/user-model"
	"go-zero-play-1/user-rpc/internal/svc"
	"go-zero-play-1/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	userData, err := user_model.ScCorpUserModel.GetUser(l.ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	g := "难"
	if userData.Gender == 1 {
		g = "女"
	}
	return &user.UserResponse{
		Id:     in.GetId(),
		Name:   userData.Name,
		Gender: g,
	}, nil
}
