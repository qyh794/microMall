package model

import "time"

type Order struct {
	ID           int64         `gorm:"primary_key;not_null;auto_increment"`
	OrderCode    string        `gorm:"unique_index;not_null" json:"order_code"`
	PayStatus    int32         `json:"pay_status"`
	ShipStatus   int32         `json:"ship_status"`
	Price        float64       `json:"price"`
	OrderDetails []OrderDetail `gorm:"ForeignKey:OrderID" json:"order_details"`
	CreateAt     time.Time
	UpdateAt     time.Time
}
