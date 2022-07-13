package service

import (
	"context"

	"github.com/ClintYJQ/name_list_proj/config"
	"github.com/gin-gonic/gin"
)

type GlobalHandler interface {
	Run() error
}

func NewGlobalHandler(config *config.Config) *GlobalHandlerImpl {
	return &GlobalHandlerImpl{
		conf:   config,
		router: gin.Default(),
		ctx:    context.Background(),
	}
}
