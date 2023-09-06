package userbusiness

import (
	"context"
	"fmt"
	usermodel "todo-app/module/user/model"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
	) (*usermodel.User, error)
}

type loginBiz struct {
	store LoginStorage
}

func NewLoginBiz(store LoginStorage) *findBiz {
	return &findBiz{store: store}
}

func (biz *findBiz) Login(ctx context.Context, condition map[string]interface{}, data *usermodel.LoginPayload) (string, error) {
	user, err := biz.store.FindUser(ctx, condition)

	if err != nil {
		return "", err
	}

	fmt.Println(user)

	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if compareErr != nil {
		return "", compareErr
	}

	// Set custom claims
	claims := &usermodel.JwtCustomClaims{
		Email: data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
