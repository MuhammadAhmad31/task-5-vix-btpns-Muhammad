package controller

import (
	"WebApiGo/app"
	"WebApiGo/helpers"
	"WebApiGo/models"
	"WebApiGo/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authHelpers helpers.AuthHelpers
	jwtHelpers  helpers.JWTHelpers
}

func NewAuthController(authHelpers helpers.AuthHelpers, jwtHelpers helpers.JWTHelpers) AuthController {
	return &authController{
		authHelpers: authHelpers,
		jwtHelpers:  jwtHelpers,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginApp app.LoginApp
	errApp := ctx.ShouldBind(&loginApp)
	if errApp != nil {
		response := router.BuildErrorResponse("Failed to process request", errApp.Error(), router.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authHelpers.VerifyCredential(loginApp.Email, loginApp.Password)
	if v, ok := authResult.(models.User); ok {
		generatedToken := c.jwtHelpers.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := router.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := router.BuildErrorResponse("Please check again your credential", "Invalid Credential", router.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerApp app.RegisterApp
	errApp := ctx.ShouldBind(&registerApp)
	if errApp != nil {
		response := router.BuildErrorResponse("Failed to process request", errApp.Error(), router.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authHelpers.IsDuplicateEmail(registerApp.Email) {
		response := router.BuildErrorResponse("Failed to process request", "Duplicate email", router.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authHelpers.CreateUser(registerApp)
		token := c.jwtHelpers.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := router.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
