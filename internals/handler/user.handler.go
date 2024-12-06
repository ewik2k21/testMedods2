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