package logic

import (
	"context"
	"go-zero-play-1/common"
	"time"

	"go-zero-play-1/user-api/internal/svc"
	"go-zero-play-1/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJWTLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJWTLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJWTLogic {
	return &GetJWTLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJWTLogic) GetJWT() (resp *types.JwtReply, err error) {
	tnw := time.Now().Unix()
	jstToken, err := common.GetJWT(l.svcCtx.Config.Auth.AccessSecret, tnw, l.svcCtx.Config.Auth.AccessExpire, 1)
	if err != nil {
		return
	}
	resp = new(types.JwtReply)
	resp.Token = jstToken
	resp.Expire = tnw + l.svcCtx.Config.Auth.AccessExpire
	return
}
