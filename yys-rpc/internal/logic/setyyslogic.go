package logic

import (
	"context"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go-zero-play-1/common/utils"
	yys_model "go-zero-play-1/model/mysql/yys-model"
	"go-zero-play-1/yys-rpc/model/outmodel"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"go-zero-play-1/yys-rpc/internal/svc"
	"go-zero-play-1/yys-rpc/types/yys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetYysLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetYysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetYysLogic {
	return &SetYysLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetYysLogic) SetYys(in *yys.SetYysRequest) (*yys.BaseResponse, error) {
	orderSp := strings.Split(in.GetOrdersn(), "-")
	if len(orderSp) != 3 {
		return &yys.BaseResponse{Code: 100}, errors.New("ordersn 错误")
	}

	resp, err := http.PostForm("https://yys.cbg.163.com/cgi/api/get_equip_detail", url.Values{"serverid": []string{orderSp[1]}, "ordersn": []string{in.Ordersn}})
	if err != nil {
		fmt.Println("请求http err", err)
		return &yys.BaseResponse{Code: 100}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读io err", err)
		return &yys.BaseResponse{Code: 100}, err
	}
	rd := new(outmodel.YysAccountInfoResponse)
	err = jsoniter.Unmarshal(body, rd)
	if err != nil {
		fmt.Println("json : ", err)
		return &yys.BaseResponse{Code: 100}, err
	}
	detailNeedUn := new(outmodel.DetailNeedUn)
	err = jsoniter.Unmarshal(utils.S2B(rd.Equip.EquipDesc), detailNeedUn)
	if err != nil {
		fmt.Println("2222222json : ", err)
		return &yys.BaseResponse{Code: 100}, err
	}
	datas := make([]*yys_model.InventoryTable, 0)
	for id, v := range detailNeedUn.Inventory {
		md := &yys_model.InventoryTable{
			Id:             0,
			Ordersn:        in.Ordersn,
			EquipName:      rd.Equip.EquipName,
			InventoryId:    id,
			Pos:            v.Pos,
			Qua:            v.Qua,
			Name:           v.Name,
			AttrsMainName:  v.Attrs[0][0],
			AttrsMainValue: v.Attrs[0][1],
			AttrsMainTable: utils.B2S(utils.TrueMarshal(v.Attrs[1:])),
		}
		if len(v.SingleAttr) != 0 {
			md.SingleAttrValue = v.SingleAttr[1]
			md.SingleAttrName = v.SingleAttr[0]
		}
		datas = append(datas, md)
	}

	err = yys_model.InventoryTableModel.InsertMany(l.ctx, datas)
	if err != nil {
		return &yys.BaseResponse{
			Code: 100,
		}, err
	}
	return &yys.BaseResponse{
		Code: 0,
	}, nil
}
