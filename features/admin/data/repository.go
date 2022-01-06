package data

import (
	"SistemVaksinAPI/features/admin"

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

func (ur *mysqlAdminRepository) InsertAdmin(data admin.AdminCore) (resp admin.AdminCore, err error) {
	record := fromCore(data)

	if err := ur.Conn.Create(&record).Error; err != nil {
		return admin.AdminCore{}, err
	}
	return toCore(&record), nil
}
