package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/qyh794/microMall/order/domain/model"
)

type IOrderRepository interface {
	InitTable() error
	FindOrderByID(int64) (*model.Order, error)
	CreateOrder(*model.Order) (int64, error)
	DeleteOrderByID(int64) error
	UpdateOrder(*model.Order) error
	FindAll() ([]model.Order, error)
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64, int32) error
}

// 创建orderRepository
func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{mysqlDb: db}
}

type OrderRepository struct {
	mysqlDb *gorm.DB
}

func (o *OrderRepository) UpdateShipStatus(orderId int64, shipStatus int32) error {
	//TODO implement me
	//panic("implement me")
	// update order set ship_status = ? where id = ?
	db := o.mysqlDb.Model(&model.Order{}).Where("id = ?", orderId).UpdateColumn("ship_status", shipStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("update ship status failed")
	}
	return nil
}

func (o *OrderRepository) UpdatePayStatus(orderId int64, payStatus int32) error {
	//TODO implement me
	//panic("implement me")
	// update order set pay_status = ? where id = ?
	db := o.mysqlDb.Model(&model.Order{}).Where("id = ?", orderId).UpdateColumn("pay_status", payStatus)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("update pay status failed")
	}
	return nil
}

// 初始化表
func (o *OrderRepository) InitTable() error {
	return o.mysqlDb.CreateTable(&model.Order{}).Error
}

// 根据ID查找Order信息
func (o *OrderRepository) FindOrderByID(orderID int64) (order *model.Order, err error) {
	order = &model.Order{}
	// select * from order, order_detail left join order_detail on order.orderID = order_detail.orderID
	// db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;
	return order, o.mysqlDb.First(order, orderID).Error
}

// 创建Order信息
func (o *OrderRepository) CreateOrder(order *model.Order) (int64, error) {
	return order.ID, o.mysqlDb.Create(order).Error
}

// 根据ID删除Order信息
func (o *OrderRepository) DeleteOrderByID(orderID int64) error {
	return o.mysqlDb.Where("id = ?", orderID).Delete(&model.Order{}).Error
}

// 更新Order信息
func (o *OrderRepository) UpdateOrder(order *model.Order) error {
	return o.mysqlDb.Model(order).Update(order).Error
}

// 获取结果集
func (o *OrderRepository) FindAll() (orderAll []model.Order, err error) {
	return orderAll, o.mysqlDb.Find(&orderAll).Error
}
