package server

import (
	"context"
	"net/http"
	"testMedods2/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GinServer interface {
	Start(ctx context.Context, httpAddress string) error
	Shutdown(ctx context.Context) error
	RegisterGroupRoute(path string, routes []interfacesx.RouteDefinition, middleWare ...gin.HandlerFunc)
}

type GinServerBuilder struct {
}

type ginServer struct {
	engine *gin.Engine
	server *http.Server
}

func NewGinServerBuilder() *GinServerBuilder {
	return &GinServerBuilder{}
}

func (b *GinServerBuilder) Build() GinServer {
	engine := gin.Default()
	return &ginServer{engine: engine}
}

func (gs *ginServer) Start(ctx context.Context, httpAddress string) error {
	gs.server = &http.Server{
		Addr:    httpAddress,
		Handler: gs.engine,
	}

	go func() {
		if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Listening %s \n", err)
		}

	}()

	logrus.Infof("Http server is running on Port %s", httpAddress)
	return nil
}

func (gs *ginServer) Shutdown(ctx context.Context) error {
	if err := gs.server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shutdown %s", err)
	}

	logrus.Info("Server is exiting")

	return nil
}

func (gs *ginServer) RegisterGroupRoute(path string, routes []interfacesx.RouteDefinition, middleWare ...gin.HandlerFunc) {
	group := gs.engine.Group(path)
	group.Use(middleWare...)
	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Path, route.Handler)
		case "POST":
			group.POST(route.Path, route.Handler)
		case "PUT":
			group.PUT(route.Path, route.Handler)
		case "DELETE":
			group.DELETE(route.Path, route.Handler)
		case "PATCH":
			group.PATCH(route.Path, route.Handler)

		default:
			logrus.Errorf("Invalid https method")
		}
	}
}
