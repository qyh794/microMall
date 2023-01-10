package repository

import (
	"github.com/qyh794/microMall/cart/domain/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)
	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

// 创建cartRepository
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

// CleanCart clean cart by user id
func (c *CartRepository) CleanCart(userID int64) error {
	//TODO implement me
	//panic("implement me")
	// delete from cart where user_id = ?
	return c.mysqlDb.Where("user_id = ?", userID).Delete(model.Cart{}).Error
}

func (c *CartRepository) IncrNum(cartID int64, num int64) error {
	//TODO implement me
	//panic("implement me")
	//  UPDATE "cart" SET "num" = "num" - num WHERE "id" = cartID;
	cartObj := &model.Cart{ID: cartID}
	return c.mysqlDb.Model(cartObj).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (c *CartRepository) DecrNum(cartID int64, num int64) error {
	//TODO implement me
	//panic("implement me")
	cartObj := &model.Cart{ID: cartID}
	// UPDATE "cart" SET "num"="num" - num WHERE "cartID"=cartID AND "num">=num;
	db := c.mysqlDb.Model(cartObj).Where("num = ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("decr failed")
	}
	return nil
}

// 初始化表
func (c *CartRepository) InitTable() error {
	return c.mysqlDb.CreateTable(&model.Cart{}).Error
}

// 根据ID查找Cart信息
func (c *CartRepository) FindCartByID(cartID int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, c.mysqlDb.First(cart, cartID).Error
}

// 创建Cart信息
func (c *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := c.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("create cart failed")
	}
	return cart.ID, nil
}

// 根据ID删除Cart信息
func (c *CartRepository) DeleteCartByID(cartID int64) error {
	return c.mysqlDb.Where("id = ?", cartID).Delete(&model.Cart{}).Error
}

// 更新Cart信息
func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDb.Model(cart).Update(cart).Error
}

// 获取结果集
func (c *CartRepository) FindAll(userID int64) (cartAll []model.Cart, err error) {
	return cartAll, c.mysqlDb.Where("user_id = ?", userID).Find(&cartAll).Error
}
