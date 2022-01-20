package request

import "SistemVaksinAPI/features/admin"

type Admin struct {
	AdminID  int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToCore(req Admin) admin.AdminCore {
	return admin.AdminCore{
		AdminID:  int(req.AdminID),
		Username: req.Username,
		Password: req.Password,
	}
}
