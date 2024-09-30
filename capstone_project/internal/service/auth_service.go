package service

import (
	"capstone_project/internal/models"
	"capstone_project/internal/repository"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthServiceInterface interface {
	CreateUser(user *models.User) error
	LoginUser(identifier, password string) (*models.User, string, error)
}

type AuthService struct {
	repo      *repository.UserRepository
	jwtSecret string
}

func NewAuthService(repo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{repo: repo, jwtSecret: jwtSecret}
}

func (s *AuthService) CreateUser(user *models.User) error {
	existingUser, err := s.repo.GetUserByUsername(user.Username)

	if err == nil && existingUser != nil {
		return errors.New("username already exists")
	}

	existingUser, err = s.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	fmt.Println(hashedPassword)

	err = s.repo.Create(user)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (s *AuthService) LoginUser(identifier, password string) (*models.User, string, error) {
	user, err := s.repo.GetUserByEmailOrUsername(identifier)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := s.generateJWT(user)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

func (s *AuthService) generateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	return token.SignedString([]byte(s.jwtSecret))
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
