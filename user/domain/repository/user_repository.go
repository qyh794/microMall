package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/qyh794/microMall/user/domain/model"
)

type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	FindAll() ([]model.User, error)
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

func (u *UserRepository) InitTable() error {
	//TODO implement me
	return u.mysqlDB.CreateTable(&model.User{}).Error
	//panic("implement me")
}

func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	//TODO implement me
	user := &model.User{}
	return user, u.mysqlDB.Where("UserName = ?", name).Find(user).Error
	//panic("implement me")
}

func (u *UserRepository) FindUserByID(id int64) (*model.User, error) {
	//TODO implement me
	user := &model.User{}
	return user, u.mysqlDB.First(user, id).Error
	//panic("implement me")
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	//TODO implement me
	//panic("implement me")
	return user.ID, u.mysqlDB.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(id int64) error {
	//TODO implement me
	return u.mysqlDB.Delete(&model.User{}, id).Error
	//panic("implement me")
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	//TODO implement me
	return u.mysqlDB.Model(user).Update(&user).Error
	//panic("implement me")
}

func (u *UserRepository) FindAll() ([]model.User, error) {
	//TODO implement me
	users := make([]model.User, 0)
	return users, u.mysqlDB.Find(&users).Error
	//panic("implement me")
}
