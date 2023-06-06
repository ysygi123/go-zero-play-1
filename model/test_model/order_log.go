package test_model

import (
	"context"
	"go-zero-play-1/common/symysql"
)

type OrderLog struct {
	ID             int    `gorm:"column:id;primary_key:auto_increment"`
	OrderMasterID  int    `gorm:"column:order_master_id;not null"`
	CID            int    `gorm:"column:cid;not null"`
	PayUserID      int    `gorm:"column:pay_user_id;not null"`
	OrderPayTime   int64  `gorm:"column:order_pay_time;not null"`
	OrderStartTime int64  `gorm:"column:order_start_time;not null"`
	Product        []byte `gorm:"column:product;type:json"`
}

func (OrderLog) TableName() string {
	return "order_log"
}

var OrderLogModel *OrderLog = &OrderLog{}

func (o *OrderLog) Insert(ctx context.Context, data *OrderLog) (err error) {
	err = symysql.GetDbSession(ctx).Table(o.TableName()).Create(data).Error
	return
}

func (o *OrderLog) InsertMany(ctx context.Context, datas []*OrderLog) (err error) {
	err = symysql.GetDbSession(ctx).Table(o.TableName()).Create(datas).Error
	return
}
