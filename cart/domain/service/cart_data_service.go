package service

import (
	"github.com/qyh794/microMall/cart/domain/model"
	"github.com/qyh794/microMall/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)
	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

// 创建
func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{cartRepository}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func (c *CartDataService) CleanCart(cartID int64) error {
	//TODO implement me
	//panic("implement me")
	return c.CartRepository.CleanCart(cartID)
}

func (c *CartDataService) IncrNum(cartID int64, num int64) error {
	//TODO implement me
	//panic("implement me")
	return c.CartRepository.IncrNum(cartID, num)
}

func (c *CartDataService) DecrNum(cartID int64, num int64) error {
	//TODO implement me
	//panic("implement me")
	return c.CartRepository.DecrNum(cartID, num)
}

// 插入
func (c *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return c.CartRepository.CreateCart(cart)
}

// 删除
func (c *CartDataService) DeleteCart(cartID int64) error {
	return c.CartRepository.DeleteCartByID(cartID)
}

// 更新
func (c *CartDataService) UpdateCart(cart *model.Cart) error {
	return c.CartRepository.UpdateCart(cart)
}

// 查找
func (c *CartDataService) FindCartByID(cartID int64) (*model.Cart, error) {
	return c.CartRepository.FindCartByID(cartID)
}

// 查找
func (c *CartDataService) FindAllCart(userID int64) ([]model.Cart, error) {
	return c.CartRepository.FindAll(userID)
}
