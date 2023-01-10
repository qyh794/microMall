package handler

import (
	context "context"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/qyh794/microMall/category/common"
	"github.com/qyh794/microMall/category/domain/model"
	service2 "github.com/qyh794/microMall/category/domain/service"
	category "github.com/qyh794/microMall/category/proto/category"
)

type Category struct {
	CategoryDataService service2.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	//TODO implement me
	categoryObj := &model.Category{}
	if err := common.SwapTo(request, categoryObj); err != nil {
		return err
	}
	categoryID, err := c.CategoryDataService.AddCategory(categoryObj)
	if err != nil {
		return err
	}
	response.CategoryId = categoryID
	response.Message = "add category succeed"
	return nil
	//panic("implement me")
}

func (c *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	//TODO implement me
	categoryObj := &model.Category{}
	if err := common.SwapTo(request, categoryObj); err != nil {
		return err
	}
	err := c.CategoryDataService.UpdateCategory(categoryObj)
	if err != nil {
		return err
	}
	response.Message = "update category succeed"
	//panic("implement me")
	return nil
}

func (c *Category) DeleteCategoryByID(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	//TODO implement me
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "delete category succeed"
	return nil
	//panic("implement me")
}

func (c *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
	//TODO implement me
	categoryObj, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(categoryObj, response)
	//panic("implement me")
}

func (c *Category) FindCategoryByID(ctx context.Context, request *category.FindByIDRequest, response *category.CategoryResponse) error {
	//TODO implement me
	categoryObj, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(categoryObj, response)
	//panic("implement me")
}

func (c *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.FindAllResponse) error {
	//TODO implement me
	categories, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	appendResponse(categories, response)
	return nil
	//panic("implement me")
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.FindAllResponse) error {
	//TODO implement me
	categories, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	appendResponse(categories, response)
	return nil
	//panic("implement me")
}

func (c *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.FindAllResponse) error {
	//TODO implement me
	allCategory, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	appendResponse(allCategory, response)
	panic("implement me")
}

func appendResponse(categories []*model.Category, response *category.FindAllResponse) {
	for i, _ := range categories {
		res := &category.CategoryResponse{}
		err := common.SwapTo(categories[i], res)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, res)
	}
}
