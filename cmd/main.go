package main

import (
	"github.com/jmechavez/email-registration/infrastructure/http"
	"github.com/jmechavez/email-registration/infrastructure/logger"
)

func main() {
	logger.Info("Starting the application")
	http.Start()
}
