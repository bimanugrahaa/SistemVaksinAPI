package bussiness

import (
	"SistemVaksinAPI/features/faskes"
)

type faskesUsecase struct {
	faskesData faskes.Data
}

func NewFaskesBussiness(faskesData faskes.Data) faskes.Bussiness {
	return &faskesUsecase{
		faskesData: faskesData,
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
