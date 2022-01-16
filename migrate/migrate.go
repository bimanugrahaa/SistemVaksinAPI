package migrate

import (
	"SistemVaksinAPI/config"
	m_admin "SistemVaksinAPI/features/admin/data"
	m_faskes "SistemVaksinAPI/features/faskes/data"
	m_user "SistemVaksinAPI/features/user/data"
	m_vaksin "SistemVaksinAPI/features/vaksin/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(&m_admin.Admin{})
	config.DB.AutoMigrate(&m_user.User{})
	config.DB.AutoMigrate(&m_faskes.Faskes{})
	config.DB.AutoMigrate(&m_vaksin.Vaksin{})
}
