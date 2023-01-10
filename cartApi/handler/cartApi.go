package handler

import (
	"context"
	"encoding/json"
	"errors"
	cart "github.com/qyh794/cart/proto/cart"
	cartApi "github.com/qyh794/microMall/cartApi/proto/cartApi"
	"strconv"
)

type CartApi struct {
	CartApiDataService cart.CartService
}

func (c *CartApi) FindAll(ctx context.Context, request *cartApi.Request, response *cartApi.Response) error {
	//TODO implement me
	//panic("implement me")
	if _, ok := request.Get["user_id"]; !ok {
		return errors.New("args wrong")
	}
	userIDString := request.Get["user_id"].Values[0]
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return err
	}
	all, err := c.CartApiDataService.GetAll(context.Background(), &cart.RequestCartFindAll{UserId: userID})
	bytes, err := json.Marshal(all)
	if err != nil {
		return err
	}
	response.StatusCode = 200
	response.Body = string(bytes)
	return nil
}
