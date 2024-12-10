package routes

import (
	"testMedods2/cmd/server"
	"testMedods2/internals/handler"
	"testMedods2/x/interfacesx"
)

func RegistterUserRoutes(server server.GinServer, userHandler *handler.UserHandler) {
	server.RegisterGroupRoute("/api", []interfacesx.RouteDefinition{
		{Method: "POST", Path: "/sign_up", Handler: userHandler.SignUpUser},
		{Method: "POST", Path: "/sign_in", Handler: userHandler.SignInUser},
	})
}
