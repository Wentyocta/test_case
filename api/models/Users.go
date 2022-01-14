package models

import (
	"errors"
	"time"
	"html"
	"strings"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    ID      uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name    string    `gorm:"size:255;not null;unique" json:"name"`
    Username    string    `gorm:"size:255;not null;unique" json:"username"`
	Password    string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy uint32 `json:"created_by"`
    UpdatedBy uint32 `json:"updated_by"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	if u.Password == "" {
        return errors.New("Required Password")
    }
    if u.Username == "" {
        return errors.New("Required Username")
    }
    return nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}