package cmd

import (
	"github.com/spf13/cobra"
	"go-clean-architecture/pkg/log"
	"os"
)

var logger = log.GetLogger()

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "application",
		Long:  "application",
		Run:  func(_ *cobra.Command, args []string) {},
	}
	rootCmd.AddCommand(httpService)

	err := rootCmd.Execute()
	if err != nil {
		logger.Errorf("rootCmd: %v", err)
		os.Exit(1)
	}
	return err
}

