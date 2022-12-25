package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-turk/user-auth-service/internal/core/interfaces"
	"log"
)

type userHandler struct {
	router *gin.Engine
	logger *log.Logger
	// userService

}

func NewUserHandler(router *gin.Engine, logger *log.Logger) interfaces.UserHandler {
	return &userHandler{
		router: router,
		logger: logger,
	}
}

func (u *userHandler) Run() error {
	u.router.POST("/login", u.Login)
	u.router.POST("/register", u.Register)
	u.router.POST("/refresh-token", u.RefreshToken)
	u.router.POST("/user", u.CreateUser)
	u.router.GET("/user/:id", u.GetUser)
	u.router.DELETE("/user/:id", u.DeleteUser)
	return nil
}

// Login is a function that logs in a user to the system.
// @Summary Login
// @Description Login to User Auth Service
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Login"
func (u *userHandler) Login(c *gin.Context) {
	u.logger.Println("Login INFO")
	c.JSON(200, "Login")
	return
}

// Register is a function that registers a new user to the system.
// @Summary Register
// @Description Register to User Auth Service
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Register"
// @Router /register [post]
func (u *userHandler) Register(c *gin.Context) {
	u.logger.Println("Register INFO")
	c.JSON(200, "Register")
	return
}

// RefreshToken is a function that refreshes the token of a user.
// @Summary Refresh Token
// @Description Refresh Token for User Auth Service
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Refresh
// @Router /refresh-token [post]
// @Success 200 {string} string	"Refresh
func (u *userHandler) RefreshToken(c *gin.Context) {
	u.logger.Println("RefreshToken INFO")
	c.JSON(200, "RefreshToken")
	return
}
func (u *userHandler) CreateUser(c *gin.Context) {
	u.logger.Println("CreateUser INFO")
	c.JSON(200, "CreateUser")
	return
}
func (u *userHandler) GetUser(c *gin.Context) {
	u.logger.Println("GetUser INFO")
	c.JSON(200, "GetUser")
	return
}

func (u *userHandler) DeleteUser(c *gin.Context) {
	u.logger.Println("DeleteUser INFO")
	c.JSON(200, "DeleteUser")
	return
}
