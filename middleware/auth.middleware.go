package middleware

import (
	"net/http"
	"testMedods2/config"
	"testMedods2/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtTokenValidation() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		tokenCookie, err := ctx.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				ctx.JSON(http.StatusUnauthorized, interfacesx.ErrorMessage{
					Message: err.Error(),
					Status:  interfacesx.StatusError,
					Code:    http.StatusUnauthorized,
				})
				return
			}

			ctx.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
				Message: err.Error(),
				Status:  interfacesx.StatusError,
				Code:    http.StatusBadRequest,
			})
			return
		}

		token, err := jwt.Parse(tokenCookie, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.JwtKey), nil
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
				Message: err.Error(),
				Status:  interfacesx.StatusError,
				Code:    http.StatusInternalServerError,
			})
			return
		}

		if token.Valid {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	})
}
