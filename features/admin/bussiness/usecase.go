package bussiness

import "SistemVaksinAPI/features/admin"

type adminUsecase struct {
	adminData admin.Data
}

func NewAdminBussiness(adminData admin.Data) admin.Bussiness {
	return &adminUsecase{
		adminData: adminData,
	}
}

func (au *adminUsecase) CreateAdmin(data admin.AdminCore) (resp admin.AdminCore, err error) {

	resp, err = au.adminData.InsertAdmin(data)
	if err != nil {
		return admin.AdminCore{}, err
	}

	return resp, nil
}
