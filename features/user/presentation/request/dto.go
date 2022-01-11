package request

import (
	"SistemVaksinAPI/features/user"
)

type User struct {
	UserID      int    `json:"id"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Namalengkap string `json:"namalengkap"`
	NIK         int    `json:"nik"`
}

func ToCore(req User) user.UserCore {
	return user.UserCore{
		// ID:       req.ID,
		Password:    req.Password,
		Email:       req.Email,
		Namalengkap: req.Namalengkap,
		NIK:         req.NIK,
	}
}
