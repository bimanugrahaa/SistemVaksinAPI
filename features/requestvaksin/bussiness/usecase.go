package bussiness

import (
	"SistemVaksinAPI/features/faskes"
	"SistemVaksinAPI/features/requestvaksin"
	"SistemVaksinAPI/features/vaksin"
	"errors"
)

type requestvaksinUsecase struct {
	requestvaksinData requestvaksin.Data
	vaksinData        vaksin.Bussiness
	faskesData        faskes.Bussiness
}

func NewRequestvaksinBussiness(requestvaksinData requestvaksin.Data, vaksinData vaksin.Bussiness, faskesData faskes.Bussiness) requestvaksin.Bussiness {
	return &requestvaksinUsecase{
		requestvaksinData: requestvaksinData,
		vaksinData:        vaksinData,
		faskesData:        faskesData,
	}
}

func (ru *requestvaksinUsecase) CreateRequestvaksin(data requestvaksin.RequestvaksinCore) (resp requestvaksin.RequestvaksinCore, err error) {

	resp, err = ru.requestvaksinData.InsertRequestvaksin(data)
	vaksin, _ := ru.vaksinData.GetVaksinByID(data.VaksinID_satu)
	faskes, _ := ru.faskesData.GetFaskesByID(vaksin.FaskesID)

	if err != nil {
		return requestvaksin.RequestvaksinCore{}, err
	}

	resp.Vaksin_satu = requestvaksin.VaksinCore{
		ID:          vaksin.ID,
		Jenisvaksin: vaksin.Jenisvaksin,
		Jadwal:      vaksin.Jadwal,
		Waktu:       vaksin.Waktu,
		FaskesID:    vaksin.FaskesID,
	}

	resp.Vaksin_satu.Faskes = requestvaksin.FaskesCore{
		ID:   faskes.ID,
		Nama: faskes.Nama,
	}

	return resp, nil
}

func (ru *requestvaksinUsecase) GetAllRequestvaksin() (resp []requestvaksin.RequestvaksinCore) {
	resp = ru.requestvaksinData.SelectAllRequestvaksin()

	for key, value := range resp {
		vaksin, _ := ru.vaksinData.GetVaksinByID(value.VaksinID_satu)
		vaksinDua, _ := ru.vaksinData.GetVaksinByID(value.VaksinID_dua)
		faskes, _ := ru.faskesData.GetFaskesByID(vaksin.FaskesID)
		faskesDua, _ := ru.faskesData.GetFaskesByID(vaksinDua.FaskesID)

		value.Vaksin_satu = requestvaksin.VaksinCore{
			ID:          vaksin.ID,
			Jenisvaksin: vaksin.Jenisvaksin,
			Jadwal:      vaksin.Jadwal,
			Waktu:       vaksin.Waktu,
			FaskesID:    vaksin.FaskesID,
			Faskes: requestvaksin.FaskesCore{
				ID:   faskes.ID,
				Nama: faskes.Nama,
			},
		}

		value.Vaksin_dua = requestvaksin.VaksinCore{
			ID:          vaksinDua.ID,
			Jenisvaksin: vaksinDua.Jenisvaksin,
			Jadwal:      vaksinDua.Jadwal,
			Waktu:       vaksinDua.Waktu,
			FaskesID:    vaksinDua.FaskesID,
			Faskes: requestvaksin.FaskesCore{
				ID:   faskesDua.ID,
				Nama: faskesDua.Nama,
			},
		}

		resp[key] = value
	}

	return resp
}

func (ru *requestvaksinUsecase) GetRequestvaksinByID(ID int) (resp requestvaksin.RequestvaksinCore, err error) {
	resp, err = ru.requestvaksinData.SelectRequestvaksinByID(ID)
	vaksin, _ := ru.vaksinData.GetVaksinByID(resp.VaksinID_satu)
	vaksinDua, _ := ru.vaksinData.GetVaksinByID(resp.VaksinID_dua)
	faskes, _ := ru.faskesData.GetFaskesByID(vaksin.FaskesID)
	faskesDua, _ := ru.faskesData.GetFaskesByID(vaksinDua.FaskesID)

	if err != nil {
		return requestvaksin.RequestvaksinCore{}, err
	}

	resp.Vaksin_satu = requestvaksin.VaksinCore{
		ID:          vaksin.ID,
		Jenisvaksin: vaksin.Jenisvaksin,
		Jadwal:      vaksin.Jadwal,
		Waktu:       vaksin.Waktu,
		FaskesID:    vaksin.FaskesID,
		Faskes: requestvaksin.FaskesCore{
			ID:   faskes.ID,
			Nama: faskes.Nama,
		},
	}

	resp.Vaksin_dua = requestvaksin.VaksinCore{
		ID:          vaksinDua.ID,
		Jenisvaksin: vaksinDua.Jenisvaksin,
		Jadwal:      vaksinDua.Jadwal,
		Waktu:       vaksinDua.Waktu,
		FaskesID:    vaksinDua.FaskesID,
		Faskes: requestvaksin.FaskesCore{
			ID:   faskesDua.ID,
			Nama: faskesDua.Nama,
		},
	}

	return
}

func (ru *requestvaksinUsecase) GetRequestvaksinByUserID(ID int) (resp []requestvaksin.RequestvaksinCore, err error) {
	resp, err = ru.requestvaksinData.SelectRequestvaksinByUserID(ID)
	return
}

func (ru *requestvaksinUsecase) Login(data requestvaksin.RequestvaksinCore) (resp requestvaksin.RequestvaksinCore, err error) {
	resp, err = ru.requestvaksinData.Login(data)

	vaksin, _ := ru.vaksinData.GetVaksinByID(resp.VaksinID_satu)
	vaksinDua, _ := ru.vaksinData.GetVaksinByID(resp.VaksinID_dua)
	faskes, _ := ru.faskesData.GetFaskesByID(vaksin.FaskesID)
	faskesDua, _ := ru.faskesData.GetFaskesByID(vaksinDua.FaskesID)

	if err != nil {
		return requestvaksin.RequestvaksinCore{}, err
	}

	resp.Vaksin_satu = requestvaksin.VaksinCore{
		ID:          vaksin.ID,
		Jenisvaksin: vaksin.Jenisvaksin,
		Jadwal:      vaksin.Jadwal,
		Waktu:       vaksin.Waktu,
		FaskesID:    vaksin.FaskesID,
		Faskes: requestvaksin.FaskesCore{
			ID:   faskes.ID,
			Nama: faskes.Nama,
		},
	}

	resp.Vaksin_dua = requestvaksin.VaksinCore{
		ID:          vaksinDua.ID,
		Jenisvaksin: vaksinDua.Jenisvaksin,
		Jadwal:      vaksinDua.Jadwal,
		Waktu:       vaksinDua.Waktu,
		FaskesID:    vaksinDua.FaskesID,
		Faskes: requestvaksin.FaskesCore{
			ID:   faskesDua.ID,
			Nama: faskesDua.Nama,
		},
	}

	if err != nil {
		err = errors.New("ada yg ngawur")
		return requestvaksin.RequestvaksinCore{}, err
	}

	return
}

func (ru *requestvaksinUsecase) UpdateRequestVaksinSatu(data requestvaksin.RequestvaksinCore) (resp requestvaksin.RequestvaksinCore, err error) {
	resp, err = ru.requestvaksinData.EditRequestVaksinSatu(data)

	vaksin, _ := ru.vaksinData.GetVaksinByID(resp.VaksinID_satu)
	faskes, _ := ru.faskesData.GetFaskesByID(vaksin.FaskesID)

	if err != nil {
		return requestvaksin.RequestvaksinCore{}, err
	}

	resp.Vaksin_satu = requestvaksin.VaksinCore{
		ID:          vaksin.ID,
		Jenisvaksin: vaksin.Jenisvaksin,
		Jadwal:      vaksin.Jadwal,
		Waktu:       vaksin.Waktu,
		FaskesID:    vaksin.FaskesID,
		Faskes: requestvaksin.FaskesCore{
			ID:   faskes.ID,
			Nama: faskes.Nama,
		},
	}

	return
}

func (ru *requestvaksinUsecase) UpdateRequestVaksinDua(data requestvaksin.RequestvaksinCore) (resp requestvaksin.RequestvaksinCore, err error) {
	resp, err = ru.requestvaksinData.EditRequestVaksinDua(data)

	vaksin, _ := ru.vaksinData.GetVaksinByID(resp.VaksinID_dua)
	faskes, _ := ru.faskesData.GetFaskesByID(vaksin.FaskesID)

	if err != nil {
		return requestvaksin.RequestvaksinCore{}, err
	}

	resp.Vaksin_dua = requestvaksin.VaksinCore{
		ID:          vaksin.ID,
		Jenisvaksin: vaksin.Jenisvaksin,
		Jadwal:      vaksin.Jadwal,
		Waktu:       vaksin.Waktu,
		FaskesID:    vaksin.FaskesID,
		Faskes: requestvaksin.FaskesCore{
			ID:   faskes.ID,
			Nama: faskes.Nama,
		},
	}

	return
}
