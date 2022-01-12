package factory

import (
	"SistemVaksinAPI/config"
	_admin_bussiness "SistemVaksinAPI/features/admin/bussiness"
	_admin_data "SistemVaksinAPI/features/admin/data"
	_admin_presentation "SistemVaksinAPI/features/admin/presentation"
	_faskes_bussiness "SistemVaksinAPI/features/faskes/bussiness"
	_faskes_data "SistemVaksinAPI/features/faskes/data"
	_faskes_presentation "SistemVaksinAPI/features/faskes/presentation"
	_user_bussiness "SistemVaksinAPI/features/user/bussiness"
	_user_data "SistemVaksinAPI/features/user/data"
	_user_presentation "SistemVaksinAPI/features/user/presentation"
)

type Presenter struct {
	AdminPresentation  *_admin_presentation.AdminHandler
	UserPresentation   *_user_presentation.UserHandler
	FaskesPresentation *_faskes_presentation.FaskesHandler
}

func Init() Presenter {

	//Admin
	adminData := _admin_data.NewAdminRepository(config.DB)
	adminBussiness := _admin_bussiness.NewAdminBussiness(adminData)
	adminPresentation := _admin_presentation.NewAdminHandler(adminBussiness)

	//User
	userData := _user_data.NewUserRepository(config.DB)
	userBussiness := _user_bussiness.NewUserBussiness(userData)
	userPresentation := _user_presentation.NewUserHandler(userBussiness)

	//Faskes
	faskesData := _faskes_data.NewFaskesRepository(config.DB)
	faskesBussiness := _faskes_bussiness.NewFaskesBussiness(faskesData)
	faskesPresentation := _faskes_presentation.NewFaskesHandler(faskesBussiness)

	return Presenter{
		AdminPresentation:  adminPresentation,
		UserPresentation:   userPresentation,
		FaskesPresentation: faskesPresentation,
	}
}
