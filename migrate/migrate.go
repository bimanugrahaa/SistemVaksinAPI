package migrate

import (
	"SistemVaksinAPI/config"
	m_admin "SistemVaksinAPI/features/admin/data"
	m_user  "SistemVaksinAPI/features/user/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(&m_admin.Admin{})
	config.DB.AutoMigrate(&m_user.User{})
}
