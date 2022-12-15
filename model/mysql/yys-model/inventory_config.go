package yys_model

import (
	"context"
	"go-zero-play-1/common/symysql"
)

const (
	CritPower          = "暴击伤害"
	CritRate           = "暴击"
	AttackAdditionRate = "攻击加成"
)

type InventoryConfig struct {
	Id            int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	InventoryName string `gorm:"column:inventory_name;type:varchar(255);NOT NULL" json:"inventory_name"`
	InventoryAttr string `gorm:"column:inventory_attr;type:varchar(255);NOT NULL" json:"inventory_attr"`
}

func (m *InventoryConfig) TableName() string {
	return "inventory_config"
}

var InventoryConfigModel *InventoryConfig = &InventoryConfig{}

func (m *InventoryConfig) GetAllConfig(ctx context.Context) (list []*InventoryConfig, err error) {
	list = make([]*InventoryConfig, 0)
	err = symysql.GetDbSession(ctx).Table(m.TableName()).Where("inventory_name<>''").Find(&list).Error
	return
}
