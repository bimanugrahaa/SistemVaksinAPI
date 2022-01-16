package bussiness

import (
	"SistemVaksinAPI/features/vaksin"
)

type vaksinUsecase struct {
	vaksinData vaksin.Data
}

func NewVaksinBussiness(vaksinData vaksin.Data) vaksin.Bussiness {
	return &vaksinUsecase{
		vaksinData: vaksinData,
	}
}

func (vu *vaksinUsecase) CreateVaksin(data vaksin.VaksinCore) (resp vaksin.VaksinCore, err error) {

	resp, err = vu.vaksinData.InsertVaksin(data)

	if err != nil {
		return vaksin.VaksinCore{}, err
	}

	return resp, nil
}

func (vu *vaksinUsecase) GetVaksinByID(id int) (resp vaksin.VaksinCore, err error) {
	resp, err = vu.vaksinData.SelectVaksinByID(id)
	return
}
