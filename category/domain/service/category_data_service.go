package service2

import (
	"github.com/qyh794/microMall/category/domain/model"
	"github.com/qyh794/microMall/category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]*model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByParent(int64) ([]*model.Category, error)
	FindCategoryByLevel(uint32) ([]*model.Category, error)
}

// 创建
func NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{categoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

func (c *CategoryDataService) FindAllCategory() ([]*model.Category, error) {
	//TODO implement me
	return c.CategoryRepository.FindAll()
	panic("implement me")
}

func (c *CategoryDataService) FindCategoryByName(name string) (*model.Category, error) {
	//TODO implement me
	return c.CategoryRepository.FindCategoryByName(name)
	panic("implement me")
}

func (c *CategoryDataService) FindCategoryByParent(i int64) ([]*model.Category, error) {
	//TODO implement me
	return c.CategoryRepository.FindCategoryByParent(i)
	panic("implement me")
}

func (c *CategoryDataService) FindCategoryByLevel(level uint32) ([]*model.Category, error) {
	//TODO implement me
	return c.FindCategoryByLevel(level)
	panic("implement me")
}

// 插入
func (c *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return c.CategoryRepository.CreateCategory(category)
}

// 删除
func (c *CategoryDataService) DeleteCategory(categoryID int64) error {
	return c.CategoryRepository.DeleteCategoryByID(categoryID)
}

// 更新
func (c *CategoryDataService) UpdateCategory(category *model.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}

// 查找
func (c *CategoryDataService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByID(categoryID)
}

////查找
//func (u *CategoryDataService) FindAllCategory() ([]model.Category, error) {
//	return u.CategoryRepository.FindAll()
//}
