package data

import (
	"SistemVaksinAPI/features/user"
	"SistemVaksinAPI/middleware"
	"fmt"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) user.Data {
	return &mysqlUserRepository{
		Conn: conn,
	}
}

func (ur *mysqlUserRepository) InsertUser(data user.UserCore) (resp user.UserCore, err error) {
	record := fromCore(data)

	if err := ur.Conn.Create(&record).Error; err != nil {
		return user.UserCore{}, err
	}
	return toCore(&record), nil
}

func (ur *mysqlUserRepository) Login(data user.UserCore) (resp user.UserCore, err error) {

	record := fromCore(data)

	if err := ur.Conn.Model(&User{}).Where("username = ? AND password = ?", data.Username, data.Password).First(&record).Error; err != nil {
		return user.UserCore{}, err
	}
	record.Token, _ = middleware.CreateToken(int(record.ID))

	if err != nil {
		return user.UserCore{}, err
	}

	if err := ur.Conn.Model(&User{}).Where("id = ?", data.UserID).Updates(&record).Error; err != nil {
		return user.UserCore{}, err
	}

	fmt.Println(record.Token)
	return toCore(&record), err
}
