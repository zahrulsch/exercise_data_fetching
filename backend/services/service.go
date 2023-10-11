package services

import (
	"backend/database"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *database.Database
}

func NewService() (*Service, error) {
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	return &Service{db: db}, nil
}

func (ser *Service) sendJson(ctx *gin.Context, status int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, data)
}

func (ser *Service) sendError(ctx *gin.Context, status int, errorMessage string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"status_code":   status,
		"error_message": errorMessage,
	})
}
