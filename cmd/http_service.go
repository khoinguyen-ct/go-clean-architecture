package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"go-clean-architecture/api/handler/http"
	"os"
	"os/signal"
	"syscall"
)

var httpService = &cobra.Command{
	Use:   "service",
	Short: "API Command of service",
	Long:  "API Command of service",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof("Starting service...")
		ctx := context.Background()
		err := startService(ctx)
		if err != nil {
			logger.Errorf("Failed to start http service: %v", err)
		}
	},
}

func startService(ctx context.Context) error {
	httpServer := http.NewHttpServer()
	errChan := make(chan error)
	go func() {
		errChan <- httpServer.Run(ctx)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	err := <-errChan
	if err != nil {
		httpServer.ShutDown(ctx)
		logger.Infof("Service is stopped: %v", err)
	}
	return nil
}
