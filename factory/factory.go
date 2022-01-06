package factory

import (
	"SistemVaksinAPI/config"
	_admin_bussiness "SistemVaksinAPI/features/admin/bussiness"
	_admin_data "SistemVaksinAPI/features/admin/data"
	_admin_presentation "SistemVaksinAPI/features/admin/presentation"
)

type Presenter struct {
	AdminPresentation *_admin_presentation.AdminHandler
}

func Init() Presenter {

	//Admin
	adminData := _admin_data.NewAdminRepository(config.DB)
	adminBussiness := _admin_bussiness.NewAdminBussiness(adminData)
	adminPresentation := _admin_presentation.NewAdminHandler(adminBussiness)

	return Presenter{
		AdminPresentation: adminPresentation,
	}
}
