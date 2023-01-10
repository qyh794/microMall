package handler

import (
	"context"
	"fmt"
	"github.com/qyh794/microMall/product/common"
	"github.com/qyh794/microMall/product/domain/model"
	"github.com/qyh794/microMall/product/domain/service"
	. "github.com/qyh794/microMall/product/proto/product"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (p *Product) AddProduct(ctx context.Context, request *ProductInfo, response *ResponseProduct) error {
	//TODO implement me
	productObj := &model.Product{}
	err := common.SwapTo(request, productObj)
	if err != nil {
		return err
	}
	fmt.Println(productObj)
	id, err := p.ProductDataService.AddProduct(productObj)
	if err != nil {
		return err
	}
	response.ProductId = id

	return nil
	//panic("implement me")
}

func (p *Product) FindProductByID(ctx context.Context, request *RequestID, response *ProductInfo) error {
	//TODO implement me
	productObj, err := p.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err = common.SwapTo(productObj, response); err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateProduct(ctx context.Context, request *ProductInfo, response *Response) error {
	//TODO implement me
	productObj := &model.Product{}
	err := common.SwapTo(request, productObj)
	if err != nil {
		return err
	}
	err = p.ProductDataService.UpdateProduct(productObj)
	if err != nil {
		return err
	}
	response.Msg = "update product succeed"
	return nil
}

func (p *Product) DeleteProductByID(ctx context.Context, request *ProductInfo, response *Response) error {
	//TODO implement me
	err := p.ProductDataService.DeleteProduct(request.ProductCategoryId)
	if err != nil {
		return err
	}
	response.Msg = "delete product succeed"
	return nil
}

func (p *Product) FindAllProduct(ctx context.Context, request *RequestAll, response *AllProduct) error {
	//TODO implement me
	products, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	for i, _ := range products {
		info := &ProductInfo{}
		err = common.SwapTo(products[i], info)
		if err != nil {
			return err
		}
		response.ProductInfo = append(response.ProductInfo, info)
	}
	return nil
	//panic("implement me")
}
