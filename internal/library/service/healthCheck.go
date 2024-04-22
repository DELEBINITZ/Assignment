package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Service) HealthCheck(ctx *gin.Context) error {
	fmt.Println("health check!")
	return nil
}
