package migrate

import (
	"SistemVaksinAPI/config"
	m_admin "SistemVaksinAPI/features/admin/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(&m_admin.Admin{})
}
