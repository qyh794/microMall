package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"log"
	"product/common"
	"product/proto/product"
)

func main() {
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	tracer, closer, err := common.NerTracer("go.micro.service.product.client", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		//添加注册中心
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := go_micro_service_product.NewProductService("go.micro.service.product", service.Client())
	productAdd := &go_micro_service_product.ProductInfo{
		ProductName:        "HUAWEI MATE60",
		ProductSku:         "HUAWEI MATE60",
		ProductPrice:       6999.0,
		ProductDescription: "HUAWEI MATE60 PRO",
		ProductCategoryId:  1,
		ProductImage: []*go_micro_service_product.ProductImage{
			{
				ImageName: "HUAWEI MATE60 PRO",
				ImageCode: "HUAWEI MATE60 PRO",
				ImageUrl:  "HUAWEI MATE60 PRO.url",
			},
			{
				ImageName: "HUAWEI MATE60 PRO image02",
				ImageCode: "HUAWEI MATE60 PRO image02",
				ImageUrl:  "HUAWEI MATE60 PRO image02",
			},
		},
		ProductSize: []*go_micro_service_product.ProductSize{
			{
				SizeName: "HUAWEI MATE60 PRO",
				SizeCode: "HUAWEI MATE60 PRO",
			},
		},
		ProductSeo: &go_micro_service_product.ProductSeo{
			SeoTitle:       "HUAWEI MATE60 PRO-seo",
			SeoKeywords:    "HUAWEI MATE60 PRO-seo",
			SeoDescription: "HUAWEI MATE60 PRO-seo",
			SeoCode:        "HUAWEI MATE60 PRO-seo",
		},
	}
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
