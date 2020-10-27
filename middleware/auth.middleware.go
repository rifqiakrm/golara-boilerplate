package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rifqiakrm/golara-boilerplate/controllers"
	"github.com/rifqiakrm/golara-boilerplate/models"
	"github.com/rifqiakrm/golara-boilerplate/utils/responses"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type TokenClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
	jwt.StandardClaims
}

func Auth(c *gin.Context) {
	publicKey, err := ioutil.ReadFile("./configs/rsa-key/oauth-public.key")
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "failed to initialized auth"))
		c.Abort()
		return
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "failed to initialized auth"))
		c.Abort()
		return
	}

	tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

	if len(tokenString) > 1 {
		tokenString := tokenString[1]
		token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("RS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return key, nil
		})

		if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid && err == nil {
			subject, error := strconv.Atoi(claims.Subject)

			if err != nil {
				c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, error.Error()))
				c.Abort()
				return
			}

			oauth, err := models.GetTokenById(c, controllers.DBPool, claims.Id, int64(subject))
			if err != nil {
				c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
				c.Abort()
				return
			}

			if oauth.Revoked {
				c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
			}
			controllers.UserID = int64(subject)
		}
	} else {
		c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
		c.Abort()
		return
	}
}

func AuthOptional(c *gin.Context) {
	publicKey, err := ioutil.ReadFile("./configs/rsa-key/oauth-public.key")
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "failed to initialized auth"))
		c.Abort()
		return
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "failed to initialized auth"))
		c.Abort()
		return
	}

	tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

	if len(tokenString) > 1 {
		tokenString := tokenString[1]
		token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("RS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return key, nil
		})

		fmt.Println("masuk")

		if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid && err == nil {
			subject, error := strconv.Atoi(claims.Subject)

			if err != nil {
				c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, error.Error()))
				c.Abort()
				return
			}

			oauth, err := models.GetTokenById(c, controllers.DBPool, claims.Id, int64(subject))
			if err != nil {
				c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
				c.Abort()
				return
			}

			if oauth.Revoked {
				c.JSON(http.StatusUnauthorized, responses.ErrorApiResponse(http.StatusUnauthorized, "unauthorized"))
			}
			controllers.UserID = int64(subject)
		}
	} else {
		controllers.UserID = 0
	}
}
