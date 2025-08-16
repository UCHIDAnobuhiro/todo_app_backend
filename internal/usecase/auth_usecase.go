package usecase

import (
	"errors"
	"log"
	"os"
	"time"

	"todo_backend/internal/domain"
	"todo_backend/internal/interface/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Signup(email, password string) error
	Login(email, password string) (string, error) // returns JWT
}

type authUsecase struct {
	users repository.UserRepository
}

func NewAuthUsecase(users repository.UserRepository) AuthUsecase {
	return &authUsecase{users: users}
}

func (u *authUsecase) Signup(email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &domain.User{Email: email, Password: string(hashed)}
	return u.users.Create(user)
}

func (u *authUsecase) Login(email, password string) (string, error) {
	user, err := u.users.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// 第1引数が「ハッシュ」、第2引数が「平文」 ← 順序厳守
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("[LOGIN] bcrypt NG: %v", err)
		return "", errors.New("invalid email or password")
	}
	log.Printf("[LOGIN] bcrypt OK for id=%d", user.ID)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("server misconfigured: JWT_SECRET missing")
	}

	// アクセストークン（短め）
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"exp":   time.Now().Add(24 * time.Hour).Unix(), // MVP: 24h
		"iat":   time.Now().Unix(),
		"email": user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	log.Printf("[LOGIN] success id=%d", user.ID)
	return signed, nil
}
