package service

import (
	"github.com/qyh794/microMall/order/domain/model"
	"github.com/qyh794/microMall/order/domain/repository"
)

type IOrderDataService interface {
	AddOrder(*model.Order) (int64, error)
	DeleteOrder(int64) error
	UpdateOrder(*model.Order) error
	FindOrderByID(int64) (*model.Order, error)
	FindAllOrder() ([]model.Order, error)
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64, int32) error
}

// 创建
func NewOrderDataService(orderRepository repository.IOrderRepository) IOrderDataService {
	return &OrderDataService{orderRepository}
}

type OrderDataService struct {
	OrderRepository repository.IOrderRepository
}

func (o *OrderDataService) UpdateShipStatus(orderID int64, shipStatus int32) error {
	//TODO implement me
	//panic("implement me")
	return o.OrderRepository.UpdateShipStatus(orderID, shipStatus)
}

func (o *OrderDataService) UpdatePayStatus(orderID int64, payStatus int32) error {
	//TODO implement me
	//panic("implement me")
	return o.OrderRepository.UpdatePayStatus(orderID, payStatus)
}

// 插入
func (o *OrderDataService) AddOrder(order *model.Order) (int64, error) {
	return o.OrderRepository.CreateOrder(order)
}

// 删除
func (o *OrderDataService) DeleteOrder(orderID int64) error {
	return o.OrderRepository.DeleteOrderByID(orderID)
}

// 更新
func (o *OrderDataService) UpdateOrder(order *model.Order) error {
	return o.OrderRepository.UpdateOrder(order)
}

// 查找
func (o *OrderDataService) FindOrderByID(orderID int64) (*model.Order, error) {
	return o.OrderRepository.FindOrderByID(orderID)
}

// 查找
func (o *OrderDataService) FindAllOrder() ([]model.Order, error) {
	return o.OrderRepository.FindAll()
}
