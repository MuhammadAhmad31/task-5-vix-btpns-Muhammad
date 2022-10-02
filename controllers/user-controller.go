package controller

import (
	"WebApiGo/app"
	"WebApiGo/helpers"
	"WebApiGo/router"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userHelpers helpers.UserHelpers
	jwtHelpers  helpers.JWTHelpers
}

func NewUserController(userHelpers helpers.UserHelpers, jwtHelpers helpers.JWTHelpers) UserController {
	return &userController{
		userHelpers: userHelpers,
		jwtHelpers:  jwtHelpers,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateApp app.UserUpdateApp
	errApp := context.ShouldBind(&userUpdateApp)
	if errApp != nil {
		res := router.BuildErrorResponse("Failed to process request", errApp.Error(), router.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtHelpers.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateApp.ID = id
	u := c.userHelpers.Update(userUpdateApp)
	res := router.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtHelpers.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userHelpers.Profile(id)
	res := router.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}
