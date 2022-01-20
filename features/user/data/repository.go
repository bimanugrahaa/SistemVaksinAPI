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

	if err := ur.Conn.Model(&User{}).Where("email = ? AND password = ?", data.Email, data.Password).First(&record).Error; err != nil {
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

func (ur *mysqlUserRepository) SelectUserEmail(data user.UserCore) (resp user.UserCore, err error) {
	record := fromCore(data)

	if err := ur.Conn.Model(&User{}).Where("email = ?", data.Email).First(&record).Error; err != nil {
		return user.UserCore{}, err
	}

	return toCore(&record), err
}

func (ur *mysqlUserRepository) SelectUserByID(ID int) (resp user.UserCore, err error) {
	var record User

	if err := ur.Conn.First(&record, ID).Error; err != nil {
		return user.UserCore{}, err
	}

	return toCore(&record), nil
}

func (ur *mysqlUserRepository) EditUser(data user.UserCore) (resp user.UserCore, err error) {
	record := fromCore(data)

	if err := ur.Conn.Model(&User{}).Where("id = ?", data.UserID).Updates(&record).Error; err != nil {
		return user.UserCore{}, err
	}

	return toCore(&record), nil

}
