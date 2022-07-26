// Code generated by goctl. DO NOT EDIT!
// Source: yys.proto

package server

import (
	"context"

	"go-zero-play-1/yys-rpc/internal/logic"
	"go-zero-play-1/yys-rpc/internal/svc"
	"go-zero-play-1/yys-rpc/types/yys"
)

type YysServer struct {
	svcCtx *svc.ServiceContext
	yys.UnimplementedYysServer
}

func NewYysServer(svcCtx *svc.ServiceContext) *YysServer {
	return &YysServer{
		svcCtx: svcCtx,
	}
}

func (s *YysServer) SetYys(ctx context.Context, in *yys.SetYysRequest) (*yys.BaseResponse, error) {
	l := logic.NewSetYysLogic(ctx, s.svcCtx)
	return l.SetYys(in)
}

func (s *YysServer) CalculateInventory(ctx context.Context, in *yys.CalculateInventoryReq) (*yys.CalculateInventoryResp, error) {
	l := logic.NewCalculateInventoryLogic(ctx, s.svcCtx)
	return l.CalculateInventory(in)
}
