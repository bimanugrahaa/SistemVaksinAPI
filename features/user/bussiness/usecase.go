package bussiness

import (
	"SistemVaksinAPI/features/faskes"
	"SistemVaksinAPI/features/requestvaksin"
	"SistemVaksinAPI/features/user"
	"SistemVaksinAPI/features/vaksin"
	"errors"
)

type userUsecase struct {
	userData          user.Data
	vaksinData        vaksin.Bussiness
	faskesData        faskes.Bussiness
	requestvaksinData requestvaksin.Bussiness
}

func NewUserBussiness(userData user.Data, vaksinData vaksin.Bussiness, faskesData faskes.Bussiness, requestvaksinData requestvaksin.Bussiness) user.Bussiness {
	return &userUsecase{
		userData:          userData,
		vaksinData:        vaksinData,
		faskesData:        faskesData,
		requestvaksinData: requestvaksinData,
	}
}

func (uu *userUsecase) CreateUser(data user.UserCore) (resp user.UserCore, err error) {

	email, _ := uu.userData.SelectUserEmail(data)

	if email.Email == data.Email {
		err = errors.New("data is available")
		return user.UserCore{}, err
	}

	resp, err = uu.userData.InsertUser(data)
	if err != nil {
		return user.UserCore{}, err
	}

	return resp, nil
}

func (uu *userUsecase) Login(data user.UserCore) (resp user.UserCore, err error) {
	resp, err = uu.userData.Login(data)

	if err != nil {
		err = errors.New("ada yg ngawur")
		return user.UserCore{}, err
	}

	return resp, nil
}

func (uu *userUsecase) GetUserByID(ID int) (resp user.UserCore, err error) {
	resp, err = uu.userData.SelectUserByID(ID)
	requestvaksin, _ := uu.requestvaksinData.GetRequestvaksinByUserID(ID)

	for _, value := range requestvaksin {
		vaksin, _ := uu.vaksinData.GetVaksinByID(value.VaksinID_satu)
		vaksinDua, _ := uu.vaksinData.GetVaksinByID(value.VaksinID_dua)
		faskes, _ := uu.faskesData.GetFaskesByID(vaksin.FaskesID)
		faskesDua, _ := uu.faskesData.GetFaskesByID(vaksinDua.FaskesID)

		resp.RequestVaksin = append(resp.RequestVaksin, user.RequestvaksinCore{
			ID:            value.ID,
			Nama:          value.Nama,
			NIK:           value.NIK,
			JenisKelamin:  value.JenisKelamin,
			TanggalLahir:  value.TanggalLahir,
			Nomor:         value.Nomor,
			UserID:        value.UserID,
			VaksinID_dua:  value.VaksinID_dua,
			Status_dua:    value.Status_dua,
			VaksinID_satu: value.VaksinID_satu,
			Status_satu:   value.Status_satu,
			Vaksin_satu: user.VaksinCore{
				ID:          vaksin.ID,
				Jenisvaksin: vaksin.Jenisvaksin,
				Jadwal:      vaksin.Jadwal,
				Waktu:       vaksin.Waktu,
				FaskesID:    vaksin.FaskesID,
				Faskes: user.FaskesCore{
					ID:   faskes.ID,
					Nama: faskes.Nama,
				},
			},
			Vaksin_dua: user.VaksinCore{
				ID:          vaksinDua.ID,
				Jenisvaksin: vaksinDua.Jenisvaksin,
				Jadwal:      vaksinDua.Jadwal,
				Waktu:       vaksinDua.Waktu,
				FaskesID:    vaksinDua.FaskesID,
				Faskes: user.FaskesCore{
					ID:   faskesDua.ID,
					Nama: faskesDua.Nama,
				},
			},
		})
	}
	return
}

func (uu *userUsecase) UpdateUser(data user.UserCore) (resp user.UserCore, err error) {

	userID, _ := uu.userData.SelectUserByID(data.UserID)

	resp, err = uu.userData.EditUser(data)
	resp.UserID = userID.UserID

	if err != nil {
		return user.UserCore{}, err
	}

	return resp, nil
}
