package yys_model

import (
	"context"
	"go-zero-play-1/common/symysql"
	"go-zero-play-1/common/utils"
	"strconv"
	"strings"
)

type InventoryTable struct {
	Id              int         `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	InventoryId     string      `gorm:"column:inventory_id;type:varchar(64);NOT NULL" json:"inventory_id"`
	Ordersn         string      `gorm:"column:ordersn;type:varchar(128);NOT NULL" json:"ordersn"`
	EquipName       string      `gorm:"column:equip_name;type:varchar(255);NOT NULL" json:"equip_name"`
	Pos             int         `gorm:"column:pos;type:tinyint(4);default:0;comment:位置;NOT NULL" json:"pos"`
	Qua             int         `gorm:"column:qua;type:tinyint(4);default:0;comment:星星;NOT NULL" json:"qua"`
	Name            string      `gorm:"column:name;type:varchar(16);comment:御魂名称;NOT NULL" json:"name"`
	AttrsMainName   string      `gorm:"column:attrs_main_name;type:varchar(16);comment:主属性名称;NOT NULL" json:"attrs_main_name"`
	AttrsMainValue  string      `gorm:"column:attrs_main_value;type:varchar(16);comment:主属性数值;NOT NULL" json:"attrs_main_value"`
	SingleAttrName  string      `gorm:"column:single_attr_name;type:varchar(16);comment:特殊御魂固定属性;NOT NULL" json:"single_attr_name"`
	SingleAttrValue string      `gorm:"column:single_attr_value;type:varchar(16);comment:特殊御魂固定属性;NOT NULL" json:"single_attr_value"`
	AttrsMainTable  interface{} `gorm:"column:attrs_main_table;type:JSON;comment:主属性数值;NOT NULL" json:"attrs_main_table"`
}

type InventoryTableExt struct {
	InventoryTable
	AttrsMainTableJson     []*InventorySecondAttr `gorm:"-" json:"attrs_main_table_json"`
	SingleAttrValueFloat64 float64                `gorm:"-" json:"single_attr_value_float_64"`
}

type InventorySecondAttr struct {
	Name  string
	Value float64
}

var InventoryTableModel *InventoryTable = &InventoryTable{}

func (m *InventoryTable) TableName() string {
	return "inventory_table"
}

func (m *InventoryTable) InsertMany(ctx context.Context, datas []*InventoryTable) (err error) {
	err = symysql.GetDbSession(ctx).Table(m.TableName()).Create(datas).Error
	return
}

func (m *InventoryTable) GetAllDataByOrdersn(ctx context.Context, ordersn string) (list []*InventoryTableExt, err error) {
	list = make([]*InventoryTableExt, 0)
	err = symysql.GetDbSession(ctx).Table(m.TableName()).Where("ordersn=?", ordersn).Find(&list).Error
	if err != nil {
		return
	}

	for _, v := range list {
		binterface, ok := v.AttrsMainTable.(*interface{})
		if !ok {
			continue
		}
		bb, ok := (*binterface).([]byte)
		if !ok {
			continue
		}
		v.AttrsMainTableJson = make([]*InventorySecondAttr, 0)
		rs := make([][]string, 0)
		utils.TrueUnmarshal(bb, &rs)
		for _, vv := range rs {
			f, _ := strconv.ParseFloat(strings.Trim(vv[1], "%"), 64)
			if len(v.SingleAttrValue) != 0 {
				v.SingleAttrValueFloat64, _ = strconv.ParseFloat(strings.Trim(v.SingleAttrValue, "%"), 64)
			}
			tp := &InventorySecondAttr{
				Name:  vv[0],
				Value: f,
			}
			v.AttrsMainTableJson = append(v.AttrsMainTableJson, tp)
		}
	}
	return
}
