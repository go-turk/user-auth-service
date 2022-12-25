package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-turk/user-auth-service/internal/core/handlers"
	"github.com/go-turk/user-auth-service/pkg/postgres"
	"log"
	"net/http"
)

// @title User Auth Service API
// @description This is a sample server for a user auth service.
// @version 0.1.0
func main() {
	logger := log.New(log.Writer(), "user-auth-service", log.LstdFlags)
	db, err := postgres.Connection()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("Connected to database: %s", db)
	r := gin.Default()
	healthChecker(r)
	userHandler := handlers.NewUserHandler(r, logger)
	if err := userHandler.Run(); err != nil {
		logger.Fatal(err)
	}
	if err := r.Run(":8080"); err != nil {
		logger.Fatal(err)
	}
}

// @Summary Health Checker
// @Description Health Checker for User Auth Service
// @Accept  json
// @Produce  json
// @Success 200 {object} HealthCheckResponse
// @Router /healthcheck [get]

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func healthChecker(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) {
		responseForHealthCheck := HealthCheckResponse{
			Status: "OK",
		}
		c.JSON(http.StatusOK, responseForHealthCheck)
	})
}
