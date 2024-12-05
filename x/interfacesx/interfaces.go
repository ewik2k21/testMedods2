package interfacesx

import "github.com/gin-gonic/gin"

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
