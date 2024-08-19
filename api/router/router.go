package router

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc controller.UserController) *gin.Engine {
	r := gin.Default()

	r.POST("/signup", uc.Signup)
	r.POST("/login", uc.Login)

	return r
}
