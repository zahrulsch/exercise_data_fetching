package services

import (
	"backend/database"
	"math/rand"
	"time"

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

func (ser *Service) addDelay() {
	min := 19
	max := 26
	randomNumber := rand.Intn(max-min+1) + min

	time.Sleep(time.Duration(randomNumber*100) * time.Millisecond)
}

func (ser *Service) sendJson(ctx *gin.Context, status int, data interface{}) {
	ser.addDelay()
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, data)
}

func (ser *Service) sendError(ctx *gin.Context, status int, errorMessage string) {
	ser.addDelay()
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"status_code":   status,
		"error_message": errorMessage,
	})
}
