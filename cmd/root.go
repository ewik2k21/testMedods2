package cmd

import (
	"context"
	"os"
	"testMedods2/cmd/server"
	"testMedods2/config"
	"testMedods2/provider"
	"time"

	"github.com/sirupsen/logrus"
)

func Execute() {
	builder := server.NewGinServerBuilder()
	server := builder.Build()

	ctx := context.Background()
	config.LoadEnviroment()

	db, err := config.SetUpDatabase()
	if err != nil {
		logrus.Fatalf("Error setting up database %s", err)
	}

	provider.NewProvider(db, server)

	go func() {
		if err := server.Start(ctx, os.Getenv(config.AppPort)); err != nil {
			logrus.Errorf("Error starting server: %s", err)
		}

	}()

	<-ctx.Done()
	logrus.Info("Server stoped")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("Error shutting down server %s", err)
	}

}
