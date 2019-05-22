package main

import (
	"github.com/dongshimou/golib/logger"

)
func main(){


	logger.New("test")


	logger.Debug("debug")

	logger.Error("error")

	logger.Warn("warning")

	logger.Info("info")

	defer func() {
		logger.Fatal()
	}()

}
