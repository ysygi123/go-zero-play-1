package logic

import (
	"context"
	"go-zero-play-1/user-api/internal/svc"
	"go-zero-play-1/user-api/internal/types"
	"go-zero-play-1/yys-rpc/types/yys"
	"golang.org/x/sync/singleflight"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalculateInventoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCalculateInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalculateInventoryLogic {
	return &CalculateInventoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var ordersnSingleFlight singleflight.Group

func (l *CalculateInventoryLogic) CalculateInventory(req *types.CalculateInventoryReq) (resp *types.CalculateInventoryResp, err error) {
	res, err := l.svcCtx.YysRpc.CalculateInventory(l.ctx, &yys.CalculateInventoryReq{
		Ordersn:          req.Ordersn,
		DefaultCritRate:  req.DefaultCritRate,
		DefaultCritPower: req.DefaultCritPower,
	})
	if err != nil {
		return
	}
	resp = new(types.CalculateInventoryResp)
	resp.Rst = make([]types.InventoryBig, 0)

	for i := 0; i < len(res.Rst); i++ {
		resp.Rst = append(resp.Rst, types.InventoryBig{})
		resp.Rst[i].Rst = make([]types.Inventory, 0)
		for _, v := range res.Rst[i].Rst {
			tp := types.Inventory{
				Name:            v.Name,
				Pos:             int(v.Pos),
				Attr:            v.Attr,
				InventoryId:     v.InventoryId,
				SingleAttrName:  v.SingleAttrName,
				SingleAttrValue: v.SingleAttrValue,
				SecondAttr:      make([]types.SecondAttr, 0, len(v.SecondAttr)),
			}
			for _, sa := range v.SecondAttr {
				tp.SecondAttr = append(tp.SecondAttr, types.SecondAttr{
					Name:  sa.Name,
					Value: sa.Value,
				})
			}
			resp.Rst[i].Rst = append(resp.Rst[i].Rst, tp)
		}
	}
	return
}
