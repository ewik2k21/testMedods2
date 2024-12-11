package interfacesx

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
)

type ResponseStatus string

const (
	StatusSucces ResponseStatus = "success"
	StatusError  ResponseStatus = "error"
)

type ErrorMessage struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Status  ResponseStatus `json:"status"`
}

type RouteDefinition struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

type Claims struct {
	UserIp string `json:"user_ip"`
	jwt.StandardClaims
}

type UserRegistrationRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Status  ResponseStatus `json:"status"`
	Data    UserData       `json:"data"`
}

type UserData struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"user_name"`
	Email    string    `json:"email"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDataForClaims struct {
	UserIP string `json:"user_ip"`
}

type UserSignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
