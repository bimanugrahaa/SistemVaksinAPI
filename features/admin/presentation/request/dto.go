package request

import "SistemVaksinAPI/features/admin"

type Admin struct {
	AdminID  int    `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func ToCore(req Admin) admin.AdminCore {
	return admin.AdminCore{
		// ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
	}
}
