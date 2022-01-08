package data

import (
	"SistemVaksinAPI/features/admin"
	"SistemVaksinAPI/middleware"
	"fmt"

	"gorm.io/gorm"
)

type mysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewAdminRepository(conn *gorm.DB) admin.Data {
	return &mysqlAdminRepository{
		Conn: conn,
	}
}

func (ar *mysqlAdminRepository) InsertAdmin(data admin.AdminCore) (resp admin.AdminCore, err error) {
	record := fromCore(data)

	if err := ar.Conn.Create(&record).Error; err != nil {
		return admin.AdminCore{}, err
	}
	return toCore(&record), nil
}

func (ar *mysqlAdminRepository) Login(data admin.AdminCore) (resp admin.AdminCore, err error) {

	record := fromCore(data)

	if err := ar.Conn.Model(&Admin{}).Where("username = ? AND password = ?", data.Username, data.Password).First(&record).Error; err != nil {
		return admin.AdminCore{}, err
	}
	record.Token, _ = middleware.CreateToken(int(record.ID))

	if err != nil {
		return admin.AdminCore{}, err
	}

	if err := ar.Conn.Model(&Admin{}).Where("id = ?", data.AdminID).Updates(&record).Error; err != nil {
		return admin.AdminCore{}, err
	}

	fmt.Println(record.Token)
	return toCore(&record), err
}
