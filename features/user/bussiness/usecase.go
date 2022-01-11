package bussiness

import (
	"SistemVaksinAPI/features/user"
	"errors"
)

type userUsecase struct {
	userData user.Data
}

func NewUserBussiness(userData user.Data) user.Bussiness {
	return &userUsecase{
		userData: userData,
	}
}

func (uu *userUsecase) CreateUser(data user.UserCore) (resp user.UserCore, err error) {

	email, _ := uu.userData.SelectUserEmail(data)

	if email.Email == data.Email {
		err = errors.New("data is available")
		return user.UserCore{}, err
	}

	resp, err = uu.userData.InsertUser(data)
	if err != nil {
		return user.UserCore{}, err
	}

	return resp, nil
}

func (uu *userUsecase) Login(data user.UserCore) (resp user.UserCore, err error) {
	resp, err = uu.userData.Login(data)

	if err != nil {
		err = errors.New("Ada yg ngawur")
		return user.UserCore{}, err
	}

	return resp, nil
}
