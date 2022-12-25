package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Handler
	Login(c *gin.Context)
	Register(c *gin.Context)
	RefreshToken(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
