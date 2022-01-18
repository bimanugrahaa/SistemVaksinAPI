package data

import (
	"SistemVaksinAPI/features/requestvaksin"

	"gorm.io/gorm"
)

type mysqlRequestvaksinRepository struct {
	Conn *gorm.DB
}

func NewRequestvaksinRepository(conn *gorm.DB) requestvaksin.Data {
	return &mysqlRequestvaksinRepository{
		Conn: conn,
	}
}

func (vr *mysqlRequestvaksinRepository) InsertRequestvaksin(data requestvaksin.RequestvaksinCore) (resp requestvaksin.RequestvaksinCore, err error) {

	record := fromCore(data)
	if err := vr.Conn.Create(&record).Error; err != nil {

		return requestvaksin.RequestvaksinCore{}, err
	}

	return toCore(&record), nil
}

func (vr *mysqlRequestvaksinRepository) SelectRequestvaksinByID(ID int) (resp requestvaksin.RequestvaksinCore, err error) {
	var record Requestvaksin

	if err := vr.Conn.First(&record, ID).Error; err != nil {
		return requestvaksin.RequestvaksinCore{}, err
	}

	return toCore(&record), nil
}

func (dr *mysqlRequestvaksinRepository) SelectAllRequestvaksin() (resp []requestvaksin.RequestvaksinCore) {

	var record []Requestvaksin

	if err := dr.Conn.Model(&Requestvaksin{}).Find(&record).Error; err != nil {
		return []requestvaksin.RequestvaksinCore{}
	}

	return toList(record)
}
