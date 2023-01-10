package handler

import (
	"context"
	"github.com/qyh794/microMall/user/domain/model"
	"github.com/qyh794/microMall/user/domain/service"
	user "github.com/qyh794/microMall/user/proto/user"
)

type User struct {
	UserDataServer service.IUserDataService
}

// Return a new handler
func New() *User {
	return &User{}
}

func (u *User) Register(ctx context.Context, request *user.UserRegisterRequest, response *user.UserRegisterResponse) error {
	user := &model.User{
		UserName:  request.UserName,
		FirstName: request.FirstName,
		Password:  request.Password,
	}
	if _, err := u.UserDataServer.AddUser(user); err != nil {
		return err
	}
	response.Message = "register success"
	return nil
	//TODO implement me
	//panic("implement me")
}

func (u *User) Login(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	ok, err := u.UserDataServer.CheckPassword(request.UserName, request.Password)
	if err != nil {
		return nil
	}
	response.IsSuccess = ok
	return nil
	//TODO implement me
	//panic("implement me")
}

func (u *User) GetUserInfo(ctx context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	user, err := u.UserDataServer.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response = UserForResponse(user)
	return nil
	//TODO implement me
	panic("implement me")
}

func UserForResponse(userObj *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserID = userObj.ID
	response.UserName = userObj.UserName
	response.FirstName = userObj.FirstName
	return response
}
