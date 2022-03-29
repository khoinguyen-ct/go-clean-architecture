package main

import (
	"go-clean-architecture/cmd"
	"go-clean-architecture/pkg/log"
)

var logger = log.GetLogger()

func main() {
	err := cmd.Execute()
	if err != nil {
		logger.Errorf("application execute err: %s", err)
	}
}
