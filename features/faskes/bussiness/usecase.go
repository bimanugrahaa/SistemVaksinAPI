package bussiness

import (
	"SistemVaksinAPI/features/faskes"
	"SistemVaksinAPI/features/vaksin"
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
	// if err := du.validate.Struct(data); err != nil {
	// 	return donation.Core{}, err
	// }

	resp, err = fu.faskesData.InsertFaskes(data)
	// user, _ := fu.userData.GetUserById(resp.AuthorID)
	// resp.Author.ID = user.ID
	// resp.Author.Name = user.Name

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
	vaksin, _ := fu.vaksinData.GetVaksinByID(ID)
	resp.Vaksin.ID = vaksin.ID
	resp.Vaksin.Jenisvaksin = vaksin.Jenisvaksin
	resp.Vaksin.Jadwal = vaksin.Jadwal
	resp.Vaksin.Waktu = vaksin.Waktu
	resp.Vaksin.Stokvaksin = vaksin.Stokvaksin

	return
}
