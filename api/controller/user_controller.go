package controller

import (
	"api/model"
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu}
}

func (uc *userController) Signup(c *gin.Context) {
	user := model.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	resUser, err := uc.uu.Signup(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, resUser)
}

func (uc *userController) Login(c *gin.Context) {
	user := model.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	tokenString, err := uc.uu.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
