package controllers

import (
	"net/http"
	"go_todoapp/config"
)

func StartMainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

