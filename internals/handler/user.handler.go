package handler

import (
	"net/http"
	"testMedods2/internals/services"
	"testMedods2/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	tokenService services.TokenService
	userService  services.UserService
	validate     *validator.Validate
}

func NewUserHandler(userService services.UserService, tokenService services.TokenService) *UserHandler {
	return &UserHandler{
		userService:  userService,
		tokenService: tokenService,
		validate:     validator.New(),
	}
}

func (h *UserHandler) SignUpUser(c *gin.Context) {
	var userRequest interfacesx.UserRegistrationRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	if err := h.validate.Struct(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})
		return
	}

	userData, err := h.userService.CreateUserAccount(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})
		return

	}
	c.JSON(http.StatusOK, interfacesx.UserResponse{
		Message: "User create successfully",
		Status:  interfacesx.StatusSucces,
		Code:    http.StatusOK,
		Data:    *userData,
	})

}

func (h *UserHandler) SignInUser(c *gin.Context) {
	var userCredentials interfacesx.UserCredentials
	if err := c.ShouldBindJSON(&userCredentials); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}

	passwordOK, err := h.userService.CheckPassword(userCredentials.Email, userCredentials.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})
		return
	}

	if !passwordOK {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: "Email or password incorrect",
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})
		return
	}
	userIP := h.userService.ReadUserIp(c)

	tokenString, expirationTime, err := h.tokenService.GenerateJwtToken(*userIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})
		return
	}

	refreshToken, err := h.tokenService.NewRefreshToken(userCredentials.Email, userIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})
		return
	}
	http.SetCookie(c.Writer,
		&http.Cookie{
			Name:     "token",
			Value:    *tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
			Path:     "/",
		})

	c.JSON(http.StatusOK, interfacesx.UserSignInResponse{
		AccessToken:  *tokenString,
		RefreshToken: *refreshToken,
	})

}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var refreshTokenRequest interfacesx.RefreshTokenRequest
	if err := c.ShouldBindJSON(&refreshTokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})

		return
	}
	newDataForToken, err := h.tokenService.GetNewDataForTokens(refreshTokenRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})
		return
	}
	userIP := h.userService.ReadUserIp(c)

	tokenString, expirationTime, err := h.tokenService.GenerateJwtToken(*userIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})
		return
	}

	refreshToken, err := h.tokenService.NewRefreshToken(newDataForToken.Email, userIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusInternalServerError,
		})
		return
	}

	http.SetCookie(c.Writer,
		&http.Cookie{
			Name:     "token",
			Value:    *tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
			Path:     "/",
		})

	c.JSON(http.StatusOK, interfacesx.UserSignInResponse{
		AccessToken:  *tokenString,
		RefreshToken: *refreshToken,
	})

}
