package handler

import (
	"context"
	"github.com/qyh794/microMall/cart/common"
	"github.com/qyh794/microMall/cart/domain/model"
	"github.com/qyh794/microMall/cart/domain/service"
	cart "github.com/qyh794/microMall/cart/proto/cart"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c *Cart) AddCart(ctx context.Context, request *cart.RequestAddCartInfo, response *cart.ResponseAddCart) (err error) {
	//TODO implement me
	//panic("implement me")
	cartObj := &model.Cart{}
	common.SwapTo(request, cartObj)
	response.CartId, err = c.CartDataService.AddCart(cartObj)
	return err
}

func (c *Cart) CleanCart(ctx context.Context, request *cart.RequestCleanCart, response *cart.Response) error {
	//TODO implement me
	//panic("implement me")
	err := c.CartDataService.CleanCart(request.UserId)
	if err != nil {
		return err
	}
	response.Msg = "clean cart succeed"
	return nil
}

func (c *Cart) Incr(ctx context.Context, request *cart.RequestItem, response *cart.Response) error {
	//TODO implement me
	//panic("implement me")
	err := c.CartDataService.IncrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "increment succeed"
	return nil
}

func (c *Cart) Decr(ctx context.Context, request *cart.RequestItem, response *cart.Response) error {
	//TODO implement me
	//panic("implement me")
	err := c.CartDataService.DecrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "decrement succeed"
	return nil
}

func (c *Cart) DeleteItemByID(ctx context.Context, request *cart.RequestCartID, response *cart.Response) error {
	//TODO implement me
	//panic("implement me")
	err := c.CartDataService.DeleteCart(request.Id)
	if err != nil {
		return err
	}
	response.Msg = "delete cart succeed"
	return nil
}

func (c *Cart) GetAll(ctx context.Context, request *cart.RequestCartFindAll, response *cart.ResponseCartAll) error {
	//TODO implement me
	//panic("implement me")
	allCart, err := c.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}
	for i, _ := range allCart {
		cartObj := &cart.RequestAddCartInfo{}
		err := common.SwapTo(allCart[i], cartObj)
		if err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cartObj)
	}
	return nil
}
