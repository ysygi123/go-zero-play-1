package logic

import (
	"context"
	yys_model "go-zero-play-1/model/mysql/yys-model"
	"go-zero-play-1/yys-rpc/internal/svc"
	"go-zero-play-1/yys-rpc/types/yys"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalculateInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCalculateInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalculateInventoryLogic {
	return &CalculateInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CalculateInventory 计算bjbs
func (l *CalculateInventoryLogic) CalculateInventory(in *yys.CalculateInventoryReq) (*yys.CalculateInventoryResp, error) {
	res, err := yys_model.InventoryConfigModel.GetAllConfig(l.ctx)
	if err != nil {
		return nil, err
	}
	inventoryConfigMap := make(map[string]string)
	for _, v := range res {
		inventoryConfigMap[v.InventoryName] = v.InventoryAttr
	}

	theOrderInventoryList, err := yys_model.InventoryTableModel.GetAllDataByOrdersn(l.ctx, in.Ordersn)
	if err != nil {
		return nil, err
	}
	sixInventoryList := l.inventoryClassPos(theOrderInventoryList)

	returnRes := new(yys.CalculateInventoryResp)

	l.inventoryClass(sixInventoryList, returnRes)

	return returnRes, nil
}

// inventoryClassPos 按照位置分组
func (l *CalculateInventoryLogic) inventoryClassPos(theOrderInventoryList []*yys_model.InventoryTableExt) (sixInventoryList [6][]*yys_model.InventoryTableExt) {
	sixInventoryList = [6][]*yys_model.InventoryTableExt{}
	for _, v := range theOrderInventoryList {
		//去6号除非爆伤 去除 2 4 不是攻击加成
		if (v.Pos == 6 && v.AttrsMainName != yys_model.CritPower) || ((v.Pos == 2 || v.Pos == 4) && v.AttrsMainName != yys_model.AttackAdditionRate) {
			continue
		}
		index := v.Pos - 1
		if len(sixInventoryList[index]) == 0 {
			sixInventoryList[index] = make([]*yys_model.InventoryTableExt, 0)
		}
		sixInventoryList[index] = append(sixInventoryList[index], v)
	}
	return
}

// inventoryClass 分类
func (l *CalculateInventoryLogic) inventoryClass(sixInventoryList [6][]*yys_model.InventoryTableExt, returnRes *yys.CalculateInventoryResp) {
	returnRes.Rst = make([]*yys.RepeatedInventory, 6)
	for idx := 0; idx < 6; idx++ {
		sort.Slice(sixInventoryList[idx], func(i, j int) bool {
			return l.getCritRate(sixInventoryList[idx][i]) > l.getCritRate(sixInventoryList[idx][j])
		})
		returnRes.Rst[idx] = &yys.RepeatedInventory{Rst: make([]*yys.Inventory, 0)}

		for _, v := range sixInventoryList[idx] {
			secondAttr := make([]*yys.SecondAttr, 0, len(v.AttrsMainTableJson))
			for _, vv := range v.AttrsMainTableJson {
				secondAttr = append(secondAttr, &yys.SecondAttr{
					Name:  vv.Name,
					Value: vv.Value,
				})
			}

			returnRes.Rst[idx].Rst = append(returnRes.Rst[idx].Rst, &yys.Inventory{
				Name:            v.Name,
				Pos:             int32(v.Pos),
				Attr:            v.AttrsMainName,
				InventoryId:     v.InventoryId,
				SecondAttr:      secondAttr,
				SingleAttrName:  v.SingleAttrName,
				SingleAttrValue: v.SingleAttrValueFloat64,
			})
		}
	}
}

// getCritRate 获取初始bjl
func (l *CalculateInventoryLogic) getCritRate(yysNode *yys_model.InventoryTableExt) float64 {
	for _, v := range yysNode.AttrsMainTableJson {
		if v.Name == yys_model.CritRate {
			return v.Value
		}
	}
	return 0
}

// findAllClass 尝试进行计算
func (l *CalculateInventoryLogic) findAllClass(defaultCritRate, defaultCritPower float64, sixInventoryList [6][]*yys_model.InventoryTableExt, inventoryConfigMap map[string]string) {
	for _, sixPos := range sixInventoryList[5] {
		numMap := make(map[string]int)
		numMap[sixPos.Name] = 1
		_, firstIsCrit := inventoryConfigMap[sixPos.Name]
		initCritRate := l.getCritRate(sixPos) + defaultCritRate
		if sixPos.SingleAttrName == yys_model.CritRate {
			initCritRate += sixPos.SingleAttrValueFloat64
		}
		innerRts := make([]*yys.Inventory, 0, 6)
		innerRts = append(innerRts, l.changeInventoryExtTableToGrpc(sixPos))

		for idx := 0; idx < 5; idx++ {
			for _, inventory := range sixInventoryList[idx] {
				thisCritRate := l.getCritRate(inventory)
				//不要浪费下次的计算
				if initCritRate >= 100 && thisCritRate != 0 {
					continue
				}
				//首个是bj的情况下
				if firstIsCrit {
					if _, okn := numMap[inventory.Name]; okn {
						numMap[inventory.Name]++
					} else {
						numMap[inventory.Name] = 1
					}
					if l.isCritRate(inventoryConfigMap, inventory.Name) && (numMap[inventory.Name] == 2 || numMap[inventory.Name] == 6) {
						initCritRate += 15
					}
				} else { //不含有bj

				}
			}
		}
	}
}

// changeInventoryExtTableToGrpc 转换
func (l *CalculateInventoryLogic) changeInventoryExtTableToGrpc(table *yys_model.InventoryTableExt) *yys.Inventory {
	return &yys.Inventory{
		Name:            table.Name,
		Pos:             int32(table.Pos),
		Attr:            table.AttrsMainName,
		InventoryId:     table.InventoryId,
		SingleAttrName:  table.SingleAttrName,
		SingleAttrValue: table.SingleAttrValueFloat64,
		SecondAttr:      l.changeSecondAttrModelToGrpc(table.AttrsMainTableJson),
	}
}

func (l *CalculateInventoryLogic) changeSecondAttrModelToGrpc(attr []*yys_model.InventorySecondAttr) []*yys.SecondAttr {
	x := make([]*yys.SecondAttr, 0, len(attr))
	for _, v := range attr {
		x = append(x, &yys.SecondAttr{
			Name:  v.Name,
			Value: v.Value,
		})
	}
	return x
}

func (l *CalculateInventoryLogic) isCritRate(cfMap map[string]string, name string) bool {
	_, ok := cfMap[name]
	return ok
}
