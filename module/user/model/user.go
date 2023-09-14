package usermodel

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrEmailCannotBeBlank    = errors.New("email can not be blank")
	ErrPasswordCannotBeBlank = errors.New("password can not be blank")
	ErrUserNotFound          = errors.New("user not found")
)

type User struct {
	Id        int        `json:"id" gorm:"column:id" swaggerignore:"true"`
	Email     string     `json:"email" gorm:"column:email"`
	Password  string     `json:"password" gorm:"column:password"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;" swaggerignore:"true"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;" swaggerignore:"true"`
}

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string { return "users" }

func (user User) Validate() error {
	if user.Email == "" {
		return ErrEmailCannotBeBlank
	}
	if user.Password == "" {
		return ErrPasswordCannotBeBlank
	}
	return nil
}
