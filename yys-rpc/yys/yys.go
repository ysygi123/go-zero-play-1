// Code generated by goctl. DO NOT EDIT!
// Source: yys.proto

package yys

import (
	"context"

	"go-zero-play-1/yys-rpc/types/yys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseResponse           = yys.BaseResponse
	CalculateInventoryReq  = yys.CalculateInventoryReq
	CalculateInventoryResp = yys.CalculateInventoryResp
	Inventory              = yys.Inventory
	RepeatedInventory      = yys.RepeatedInventory
	SecondAttr             = yys.SecondAttr
	SetYysRequest          = yys.SetYysRequest

	Yys interface {
		SetYys(ctx context.Context, in *SetYysRequest, opts ...grpc.CallOption) (*BaseResponse, error)
		CalculateInventory(ctx context.Context, in *CalculateInventoryReq, opts ...grpc.CallOption) (*CalculateInventoryResp, error)
	}

	defaultYys struct {
		cli zrpc.Client
	}
)

func NewYys(cli zrpc.Client) Yys {
	return &defaultYys{
		cli: cli,
	}
}

func (m *defaultYys) SetYys(ctx context.Context, in *SetYysRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	client := yys.NewYysClient(m.cli.Conn())
	return client.SetYys(ctx, in, opts...)
}

func (m *defaultYys) CalculateInventory(ctx context.Context, in *CalculateInventoryReq, opts ...grpc.CallOption) (*CalculateInventoryResp, error) {
	client := yys.NewYysClient(m.cli.Conn())
	return client.CalculateInventory(ctx, in, opts...)
}
