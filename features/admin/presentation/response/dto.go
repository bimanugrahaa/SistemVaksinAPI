package response

import "SistemVaksinAPI/features/admin"

type Admin struct {
	AdminID  int    `json:"id"`
	Username string `json:"Username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func FromCore(core admin.AdminCore) Admin {
	return Admin{
		AdminID:  core.AdminID,
		Username: core.Username,
		Token:    core.Token,
	}
}

func FromCoreLogin(core admin.AdminCore) Admin {
	return Admin{
		AdminID:  core.AdminID,
		Username: core.Username,
		Token:    core.Token,
	}
}
