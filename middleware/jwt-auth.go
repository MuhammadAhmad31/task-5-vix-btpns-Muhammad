package middleware

import (
	"WebApiGo/helpers"
	"WebApiGo/router"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthorizeJWT(jwtHelpers helpers.JWTHelpers) gin.HandlerFunc {
	return func(c *gin.Context) {
		{
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				response := router.BuildErrorResponse("Failed to process request", "no token found", nil)
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
				return
			}
			token, err := jwtHelpers.ValidateToken(authHeader)
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				log.Println("Claim[user_id]: ", claims["user_id"])
				log.Println("Claim[issuer]: ", claims["issuer"])
			} else {
				log.Println(err)
				response := router.BuildErrorResponse("Token is not valid", err.Error(), nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			}
		}
	}
}
