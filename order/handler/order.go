package handler

import (
	context "context"
	"github.com/qyh794/cart/common"
	"github.com/qyh794/microMall/order/domain/model"
	"github.com/qyh794/microMall/order/domain/service"
	order "github.com/qyh794/microMall/order/proto/order"
)

type Order struct {
	OrderDataService service.IOrderDataService
}

func (o *Order) GetOrderByID(ctx context.Context, request *order.OrderID, response *order.OrderInfo) error {
	//TODO implement me
	//panic("implement me")
	orderObj, err := o.OrderDataService.FindOrderByID(request.OrderId)
	if err != nil {
		return err
	}
	err = common.SwapTo(orderObj, response)
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) GetAllOrder(ctx context.Context, request *order.RequestAllOrder, response *order.ResponseAllOrder) error {
	//TODO implement me
	//panic("implement me")
	allOrder, err := o.OrderDataService.FindAllOrder()
	if err != nil {
		return err
	}
	for i, _ := range allOrder {
		orderObj := &order.OrderInfo{}
		err := common.SwapTo(allOrder[i], orderObj)
		if err != nil {
			return err
		}
		response.OrderInfo = append(response.OrderInfo, orderObj)
	}
	return nil
}

func (o *Order) CreateOrder(ctx context.Context, request *order.OrderInfo, response *order.OrderID) error {
	//TODO implement me
	//panic("implement me")
	orderObj := &model.Order{}
	err := common.SwapTo(request, orderObj)
	if err != nil {
		return err
	}
	orderID, err := o.OrderDataService.AddOrder(orderObj)
	if err != nil {
		return err
	}
	response.OrderId = orderID
	return nil
}

func (o *Order) DeleteOrderByID(ctx context.Context, request *order.OrderID, response *order.ResponseMessage) error {
	//TODO implement me
	//panic("implement me")
	err := o.OrderDataService.DeleteOrder(request.OrderId)
	if err != nil {
		return err
	}
	response.Msg = "delete order succeed"
	return nil
}

func (o *Order) UpdateOrderPayStatus(ctx context.Context, request *order.RequestPayStatus, response *order.ResponseMessage) error {
	//TODO implement me
	//panic("implement me")
	err := o.OrderDataService.UpdatePayStatus(request.OrderId, request.PayStatus)
	if err != nil {
		return err
	}
	response.Msg = "update order pay status succeed"
	return nil
}

func (o *Order) UpdateOrderShipStatus(ctx context.Context, request *order.RequestShipStatus, response *order.ResponseMessage) error {
	//TODO implement me
	//panic("implement me")
	err := o.OrderDataService.UpdateShipStatus(request.OrderId, request.ShipStatus)
	if err != nil {
		return err
	}
	response.Msg = "update order ship status succeed"
	return nil
}

func (o *Order) UpdateOrder(ctx context.Context, request *order.OrderInfo, response *order.ResponseMessage) error {
	//TODO implement me
	//panic("implement me")
	orderObj := &model.Order{}
	err := common.SwapTo(request, orderObj)
	if err != nil {
		return err
	}
	err = o.OrderDataService.UpdateOrder(orderObj)
	if err != nil {
		return err
	}
	response.Msg = "update order succeed"
	return nil
}
