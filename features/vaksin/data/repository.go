package data

import (
	"SistemVaksinAPI/features/vaksin"

	"gorm.io/gorm"
)

type mysqlVaksinRepository struct {
	Conn *gorm.DB
}

func NewVaksinRepository(conn *gorm.DB) vaksin.Data {
	return &mysqlVaksinRepository{
		Conn: conn,
	}
}

func (dr *mysqlVaksinRepository) InsertVaksin(data vaksin.VaksinCore) (resp vaksin.VaksinCore, err error) {

	record := fromCore(data)
	if err := dr.Conn.Create(&record).Error; err != nil {

		return vaksin.VaksinCore{}, err
	}

	return toCore(&record), nil
}

func (vr *mysqlVaksinRepository) SelectVaksinByFaskesID(id int) (resp []vaksin.VaksinCore, err error) {

	var record []Vaksin

	if err := vr.Conn.Model(&Vaksin{}).Where("faskes_id = ?", id).Find(&record).Error; err != nil {
		return []vaksin.VaksinCore{}, err
	}

	return toVaksinList(record), nil
}

func (vr *mysqlVaksinRepository) SelectVaksinByID(ID int) (resp vaksin.VaksinCore, err error) {
	var record Vaksin

	if err := vr.Conn.First(&record, ID).Error; err != nil {
		return vaksin.VaksinCore{}, err
	}

	return toCore(&record), nil
}

func (vr *mysqlVaksinRepository) SelectVaksinByFaskesID(id int) (resp []vaksin.VaksinCore, err error) {

	var record []Vaksin

	if err := vr.Conn.Model(&Vaksin{}).Where("faskes_id = ?", id).Find(&record).Error; err != nil {
		return []vaksin.VaksinCore{}, err
	}

	return toVaksinList(record), nil
}

func (vr *mysqlVaksinRepository) EditVaksinByID(data vaksin.VaksinCore) (resp vaksin.VaksinCore, err error) {
	record := fromCore(data)

	if err := vr.Conn.Model(&Vaksin{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		return vaksin.VaksinCore{}, err
	}

	return toCore(&record), nil
}
