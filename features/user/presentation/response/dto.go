package response

import "SistemVaksinAPI/features/user"

type User struct {
	UserID      int    `json:"id"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Namalengkap string `json:"namalengkap"`
	NIK         int    `json:"nik"`
	Token       string `json:"token"`
}

func FromCore(core user.UserCore) User {
	return User{
		UserID:      core.UserID,
		Namalengkap: core.Namalengkap,
		Email:       core.Email,
		Token:       core.Token,
	}
}

func FromCoreLogin(core user.UserCore) User {
	return User{
		UserID:      core.UserID,
		Namalengkap: core.Namalengkap,
		Email:       core.Email,
		Token:       core.Token,
	}
}
