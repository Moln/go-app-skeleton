package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

var UserTable = "users"

type User struct {
	Id        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time `gorm:"type:datetime"`
	Username  string
	Email     string
	Password  string `json:"-"`
}

func (u *User) PasswordVerify(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) SetPasswordHash(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
