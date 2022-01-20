package bussiness

import (
	"SistemVaksinAPI/features/faskes"
	"SistemVaksinAPI/features/vaksin"
	"errors"
)

type faskesUsecase struct {
	faskesData faskes.Data
	vaksinData vaksin.Bussiness
}

func NewFaskesBussiness(faskesData faskes.Data, vaksinData vaksin.Bussiness) faskes.Bussiness {
	return &faskesUsecase{
		faskesData: faskesData,
		vaksinData: vaksinData,
	}
}

func (fu *faskesUsecase) CreateFaskes(data faskes.FaskesCore) (resp faskes.FaskesCore, err error) {
	resp, err = fu.faskesData.InsertFaskes(data)

	if err != nil {
		return faskes.FaskesCore{}, err
	}

	return resp, nil
}

func (fu *faskesUsecase) GetAllFaskes() (resp []faskes.FaskesCore) {
	resp = fu.faskesData.SelectAllFaskes()

	return resp
}

func (fu *faskesUsecase) GetFaskesByID(ID int) (resp faskes.FaskesCore, err error) {
	resp, err = fu.faskesData.SelectFaskesByID(ID)
	vaksin, _ := fu.vaksinData.GetVaksinByFaskesID(ID)

	for _, value := range vaksin {
		resp.Vaksin = append(resp.Vaksin, faskes.VaksinCore{
			ID:          value.ID,
			Jenisvaksin: value.Jenisvaksin,
			Jadwal:      value.Jadwal,
			Waktu:       value.Waktu,
			Stokvaksin:  value.Stokvaksin,
			FaskesID:    value.FaskesID,
		})

	}
	return
}

func (fu *faskesUsecase) GetFaskesByName(data faskes.FaskesCore) (resp faskes.FaskesCore, err error) {
	resp, err = fu.faskesData.SelectFaskesByName(data)

	vaksin, _ := fu.vaksinData.GetVaksinByFaskesID(resp.ID)

	for _, value := range vaksin {
		resp.Vaksin = append(resp.Vaksin, faskes.VaksinCore{
			ID:          value.ID,
			Jenisvaksin: value.Jenisvaksin,
			Jadwal:      value.Jadwal,
			Waktu:       value.Waktu,
			Stokvaksin:  value.Stokvaksin,
			FaskesID:    value.FaskesID,
		})

	}

	if err != nil {
		err = errors.New("Ada yg ngawur")
		return faskes.FaskesCore{}, err
	}

	return resp, nil
}
