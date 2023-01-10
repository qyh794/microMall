package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/cart/common"
	cart "github.com/qyh794/cart/proto/cart"
	"github.com/qyh794/microMall/cartApi/handler"
	cartApi "github.com/qyh794/microMall/cartApi/proto/cartApi"
	"net"
	"net/http"
)

func main() {
	consulObj := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	tracer, closer, err := common.NewTracer("go.micro.api.cartapi", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	streamHandler := hystrix.NewStreamHandler()
	streamHandler.Start()

	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", ""), streamHandler)
		if err != nil {
			log.Error(err)
			fmt.Println("1")
		}
	}()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8086"),
		micro.Registry(consulObj),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(NewClientHystrixWrapper()),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	cartService := cart.NewCartService("go.micro.service.cart", service.Client())
	cartService.AddCart(context.Background(), &cart.RequestAddCartInfo{
		UserId:    3,
		ProductId: 4,
		SizeId:    5,
		Num:       5,
	})

	err = cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartApiDataService: cartService})
	if err != nil {
		fmt.Println("2")
		log.Error(err)
	}
	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
		fmt.Println("3")

	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//run 正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
