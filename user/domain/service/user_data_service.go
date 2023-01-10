package service

import (
	"errors"
	"github.com/qyh794/microMall/user/domain/model"
	"github.com/qyh794/microMall/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(*model.User, bool) error
	FindUserByName(string) (*model.User, error)
	CheckPassword(string, string) (bool, error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

func (u *UserDataService) AddUser(user *model.User) (int64, error) {
	//TODO implement me
	encryptedPassword, err := GeneratePassword(user.Password)
	if err != nil {
		return user.ID, err
	}
	user.Password = string(encryptedPassword)
	return u.UserRepository.CreateUser(user)
	//panic("implement me")
}

func (u *UserDataService) DeleteUser(id int64) error {
	return u.UserRepository.DeleteUserByID(id)
	//TODO implement me
	//panic("implement me")
}

func (u *UserDataService) UpdateUser(user *model.User, isChangePassword bool) error {
	if isChangePassword {
		encryptedPassword, err := GeneratePassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(encryptedPassword)
	}
	return u.UserRepository.UpdateUser(user)
	//TODO implement me
	//panic("implement me")
}

func (u *UserDataService) FindUserByName(userName string) (*model.User, error) {
	return u.UserRepository.FindUserByName(userName)
	//TODO implement me
	//panic("implement me")
}

func (u *UserDataService) CheckPassword(userName string, pwd string) (bool, error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.Password)
	//TODO implement me
	//panic("implement me")
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false, errors.New("wrong password")
	}
	return true, nil
}
