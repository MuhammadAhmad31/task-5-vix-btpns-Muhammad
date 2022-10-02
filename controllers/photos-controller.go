package controller

import (
	"WebApiGo/app"
	"WebApiGo/helpers"
	"WebApiGo/models"
	"WebApiGo/router"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PhotoController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type photoController struct {
	photoHelpers helpers.PhotoHelpers
	jwtHelpers   helpers.JWTHelpers
}

func NewPhotoController(photoHelp helpers.PhotoHelpers, jwtHelp helpers.JWTHelpers) PhotoController {
	return &photoController{
		photoHelpers: photoHelp,
		jwtHelpers:   jwtHelp,
	}
}

func (c *photoController) All(context *gin.Context) {
	var photos []models.Photos = c.photoHelpers.All()
	res := router.BuildResponse(true, "OK", photos)
	context.JSON(http.StatusOK, res)
}

func (c *photoController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := router.BuildErrorResponse("No param id was found", err.Error(), router.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var photo models.Photos = c.photoHelpers.FindByID(id)
	if (photo == models.Photos{}) {
		res := router.BuildErrorResponse("Data not Found", "No data with given id", router.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := router.BuildResponse(true, "OK", photo)
		context.JSON(http.StatusOK, res)
	}
}

func (c *photoController) Insert(context *gin.Context) {
	var photoCreateApp app.PhotosCreateApp
	errApp := context.ShouldBind(&photoCreateApp)
	if errApp != nil {
		res := router.BuildErrorResponse("Failed to process request", errApp.Error(), router.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			photoCreateApp.UserID = convertedUserID
		}
		result := c.photoHelpers.Insert(photoCreateApp)
		response := router.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *photoController) Update(context *gin.Context) {
	var photoUpdateApp app.PhotosUpdateApp
	errApp := context.ShouldBind(&photoUpdateApp)
	if errApp != nil {
		res := router.BuildErrorResponse("Failed to process request", errApp.Error(), router.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtHelpers.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.photoHelpers.IsAllowedToEdit(userID, photoUpdateApp.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			photoUpdateApp.UserID = id
		}
		result := c.photoHelpers.Update(photoUpdateApp)
		response := router.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := router.BuildErrorResponse("You dont have permission", "you are not the owner", router.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *photoController) Delete(context *gin.Context) {
	var photo models.Photos
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := router.BuildErrorResponse("Failed to get id", "No param id were found", router.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	photo.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtHelpers.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.photoHelpers.IsAllowedToEdit(userID, photo.ID) {
		c.photoHelpers.Delete(photo)
		res := router.BuildResponse(true, "Deleted", router.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := router.BuildErrorResponse("You dont have permission", "you are not the owner", router.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *photoController) getUserIDByToken(token string) string {
	aToken, err := c.jwtHelpers.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
