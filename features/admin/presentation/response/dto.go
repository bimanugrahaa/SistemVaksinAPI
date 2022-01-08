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
		// ID:       req.ID,
		Username: core.Username,
		Password: core.Password,
		Token:    core.Token,
	}
}
