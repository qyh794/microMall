package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/qyh794/microMall/product/domain/model"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

// 创建productRepository
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

// 初始化表
func (u *ProductRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Product{}, &model.ProductSeo{}, &model.ProductImage{}, &model.ProductSize{}).Error
}

// 根据ID查找Product信息
func (u *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	return product, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").First(product, productID).Error
}

// 创建Product信息
func (u *ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	return product.ID, u.mysqlDb.Create(product).Error
}

// 根据ID删除Product信息
func (u *ProductRepository) DeleteProductByID(productID int64) error {
	tx := u.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	// delete form Product where id = ?
	err := tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// delete from Product_size where images_product_id = ?
	err = tx.Unscoped().Where("images_product_id = ?", productID).Delete(&model.ProductImage{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.Product{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.Product{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

// 更新Product信息
func (u *ProductRepository) UpdateProduct(product *model.Product) error {
	return u.mysqlDb.Model(product).Update(product).Error
}

// 获取结果集
func (u *ProductRepository) FindAll() (productAll []model.Product, err error) {
	return productAll, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&productAll).Error
}
