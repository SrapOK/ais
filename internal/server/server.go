package server

import (
	"kis/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(cfg config.HTTPserver, handler *gin.Engine) *http.Server {
	//Todo: добавить auth middleware
	server := &http.Server{
		Handler:      handler,
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return server
}
