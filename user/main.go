package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/qyh794/microMall/user/domain/repository"
	service2 "github.com/qyh794/microMall/user/domain/service"
	"github.com/qyh794/microMall/user/handler"
	user "github.com/qyh794/microMall/user/proto/user"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)
	// init service
	srv.Init()
	// init database
	db, err := gorm.Open("mysql", "root:123456@/db3?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.SingularTable(true)
	// init database table
	rp := repository.NewUserRepository(db)
	rp.InitTable()
	// create service obj
	dataServiceObj := service2.NewUserDataService(repository.NewUserRepository(db))
	// Register handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataServer: dataServiceObj})
	if err != nil {
		fmt.Println(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
