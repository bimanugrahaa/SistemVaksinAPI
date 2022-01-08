package data

import (
	"SistemVaksinAPI/features/admin"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string
	Password string
	Token    string
}

func toCore(a *Admin) admin.AdminCore {
	return admin.AdminCore{
		AdminID:  int(a.ID),
		Username: a.Username,
		Password: a.Password,
		Token:    a.Token,
	}
}

func toCoreList(resp []Admin) []admin.AdminCore {
	a := []admin.AdminCore{}
	for _, value := range resp {
		a = append(a, toCore(&value))
	}

	return a
}

func fromCore(core admin.AdminCore) Admin {
	return Admin{
		Username: core.Username,
		Password: core.Password,
	}
}
