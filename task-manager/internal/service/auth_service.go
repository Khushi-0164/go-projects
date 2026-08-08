package service

import (
	"errors"
	"time"

	"task-manager/config"
	"task-manager/internal/models"
	"task-manager/internal/repository"
	"task-manager/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrEmailTaken         = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

// Claims mirrors middleware.Claims — duplicated here so the service
// doesn't depend on the middleware package (services should only
// depend on repository/models, never on HTTP-layer packages).
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Signup(email, name, password string) (*models.User, error) {
	if _, err := s.repo.FindByEmail(email); err == nil {
		return nil, ErrEmailTaken
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		Name:         name,
		PasswordHash: hash,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// Login returns the authenticated user and a signed JWT.
func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, "", ErrInvalidCredentials
	}

	claims := Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(config.JWTSecret())
	if err != nil {
		return nil, "", err
	}

	return user, signed, nil
}
