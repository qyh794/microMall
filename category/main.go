package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/qyh794/microMall/category/common"
	"github.com/qyh794/microMall/category/domain/repository"
	service2 "github.com/qyh794/microMall/category/domain/service"
	"github.com/qyh794/microMall/category/handler"
	category "github.com/qyh794/microMall/category/proto/category"
)

func main() {
	// config center
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}

	// register center
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		// set IP and port
		micro.Address("127.0.0.1:8082"),
		// add consul as registry center
		micro.Registry(consulRegistry),
	)

	// get mysql config
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	fmt.Printf("user:%s\npassword:%s\ndatabase:%s\nport:%d\nhost:%s\n", mysqlInfo.User, mysqlInfo.Pwd, mysqlInfo.Database, mysqlInfo.Port, mysqlInfo.Host)
	//"mysql",root+:123456@/db3?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	// file=category/main.go:48 level= error Error 1045:
	// Access denied for user ''@'172.17.0.1' (using password: NO)
	// file=category/main.go:48 level= error Error 1045:
	// Access denied for user ''@'172.17.0.1' (using password: NO)

	db.SingularTable(true)
	//  sql: unknown driver "mysql" (forgotten import?)
	rp := repository.NewCategoryRepository(db)
	err = rp.InitTable()
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// Initialise service
	service.Init()
	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))

	// Register Handler
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
