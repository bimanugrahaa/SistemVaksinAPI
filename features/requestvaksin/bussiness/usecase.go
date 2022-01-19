package bussiness

import (
	"SistemVaksinAPI/features/requestvaksin"
)

type requestvaksinUsecase struct {
	requestvaksinData requestvaksin.Data
}

func NewRequestvaksinBussiness(requestvaksinData requestvaksin.Data) requestvaksin.Bussiness {
	return &requestvaksinUsecase{
		requestvaksinData: requestvaksinData,
	}
}

func (ru *requestvaksinUsecase) CreateRequestvaksin(data requestvaksin.RequestvaksinCore) (resp requestvaksin.RequestvaksinCore, err error) {

	resp, err = ru.requestvaksinData.InsertRequestvaksin(data)

	if err != nil {
		return requestvaksin.RequestvaksinCore{}, err
	}

	return resp, nil
}

func (ru *requestvaksinUsecase) GetAllRequestvaksin() (resp []requestvaksin.RequestvaksinCore) {
	resp = ru.requestvaksinData.SelectAllRequestvaksin()

	return resp
}

func (ru *requestvaksinUsecase) GetRequestvaksinByID(ID int) (resp requestvaksin.RequestvaksinCore, err error) {
	resp, err = ru.requestvaksinData.SelectRequestvaksinByID(ID)

	return
}
