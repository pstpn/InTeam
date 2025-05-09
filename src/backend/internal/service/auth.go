package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-faster/errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/logger"
)

type IAuthService interface {
	Register(ctx context.Context, req *dto.RegisterReq) (string, error)
	Login(ctx context.Context, req *dto.LoginReq) (string, error)
}

type AuthService struct {
	Logger   logger.Interface
	UserRepo storage.IUserStorage

	signingKey     string
	accessTokenTTL time.Duration
}

func NewAuthService(
	l logger.Interface,
	repo storage.IUserStorage,
	signingKey string,
	accessTokenTTL time.Duration,
) *AuthService {
	return &AuthService{
		Logger:         l,
		UserRepo:       repo,
		signingKey:     signingKey,
		accessTokenTTL: accessTokenTTL,
	}
}

func (s *AuthService) Register(ctx context.Context, req *dto.RegisterReq) (string, error) {
	_, err := s.UserRepo.GetUserByEmail(ctx, req.Email)
	if !errors.Is(err, storage.ErrNotFound) {
		s.Logger.Errorf("get user by email")
		return "", errors.New("get user by email")
	}

	hashed, err := hashAndSalt([]byte(req.Password))
	if err != nil {
		s.Logger.Errorf("hash pasword: %s", err.Error())
		return "", fmt.Errorf("hash password: %w", err)
	}
	user := &model.User{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Role:     model.RoleUser,
		Password: hashed,
	}

	err = s.UserRepo.Create(ctx, user)
	if err != nil {
		s.Logger.Errorf("create user: %s", err.Error())
		return "", fmt.Errorf("create user: %w", err)
	}

	return generateToken(user.ID, string(user.Role), s.signingKey, s.accessTokenTTL)
}

func (s *AuthService) Login(ctx context.Context, req *dto.LoginReq) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.Logger.Errorf("get user by email: %s", err.Error())
		return "", fmt.Errorf("get user by email: %w", err)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password))
	if err != nil {
		s.Logger.Errorf("wrong password")
		return "", errors.New("wrong password")
	}

	return generateToken(user.ID, string(user.Role), s.signingKey, s.accessTokenTTL)
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int
	Role   string
}

func generateToken(userID int, role string, signingKey string, ttl time.Duration) (string, error) {
	expiresAt := &jwt.NumericDate{
		Time: time.Now().Add(ttl),
	}

	issuedAt := &jwt.NumericDate{
		Time: time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
		userID,
		role,
	})

	return token.SignedString([]byte(signingKey))
}

func hashAndSalt(pass []byte) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
