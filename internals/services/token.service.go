package services

import (
	"fmt"
	"testMedods2/config"
	"testMedods2/internals/repository"
	"testMedods2/x/interfacesx"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/rand"
)

type TokenService interface {
	GenerateJwtToken(userIp string) (*string, error)
	GetClaimsFromToken(tokenCookie string) (*interfacesx.Claims, error)
	NewRefreshToken(email string) (*string, error)
}

type tokenService struct {
	userRepo repository.UserRepository
}

func NewTokenService(userRepo repository.UserRepository) TokenService {
	return &tokenService{
		userRepo: userRepo,
	}
}

func (ts *tokenService) GenerateJwtToken(userIp string) (*string, error) {
	expirationTime := time.Now().UTC().Add(time.Hour * 2)

	claims := &interfacesx.Claims{
		UserIp: userIp,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtKey))

	if err != nil {
		logrus.Errorf("jwt token not signed: %s", err)
		return nil, err
	}

	return &tokenString, nil

}

func (ts *tokenService) GetClaimsFromToken(tokenCookie string) (*interfacesx.Claims, error) {
	token, err := jwt.Parse(tokenCookie, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})
	if err != nil {
		logrus.Errorf("token parse error: %s", err)
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &interfacesx.Claims{
			UserIp: claims["user_ip"].(string),
		}, nil
	} else {
		logrus.Errorf("invalid token: %s", err)
		return nil, err
	}

}

func (ts *tokenService) NewRefreshToken(email string) (*string, error) {
	b := make([]byte, 32)

	n := rand.NewSource(uint64(time.Now().Unix()))
	r := rand.New(n)

	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}
	refreshToken := fmt.Sprintf("%x", b)

	if err := ts.userRepo.UpdateRefreshTokenDb(&email, &refreshToken); err != nil {
		return nil, err

	}

	return &refreshToken, nil
}
