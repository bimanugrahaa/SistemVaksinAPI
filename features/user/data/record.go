package data

import (
	"SistemVaksinAPI/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      int
	Username    string
	Email       string
	Namalengkap string
	NIK         int
	Password    string
	Token       string
}

func toCore(u *User) user.UserCore {
	return user.UserCore{
		UserID:   int(u.ID),
		Username: u.Username,
		Password: u.Password,
		Token:    u.Token,
	}
}

func toCoreList(resp []User) []user.UserCore {
	u := []user.UserCore{}
	for _, value := range resp {
		u = append(u, toCore(&value))
	}

	return u
}

func fromCore(core user.UserCore) User {
	return User{
		Username: core.Username,
		Password: core.Password,
	}
}
