package data

import (
	"SistemVaksinAPI/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      int
	Email       string
	Namalengkap string
	NIK         int
	Password    string
	Token       string
}

func toCore(u *User) user.UserCore {
	return user.UserCore{
		UserID:      int(u.ID),
		Email:       u.Email,
		Password:    u.Password,
		Namalengkap: u.Namalengkap,
		NIK:         u.NIK,
		Token:       u.Token,
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
		Email:       core.Email,
		Password:    core.Password,
		Namalengkap: core.Namalengkap,
		NIK:         core.NIK,
	}
}
