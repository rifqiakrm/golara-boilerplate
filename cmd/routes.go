package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqiakrm/golara-boilerplate/controllers"
	"github.com/rifqiakrm/golara-boilerplate/middleware"
	"github.com/rifqiakrm/golara-boilerplate/utils/responses"
	"net/http"
)

func Routes() {
	router.GET("/hello", controllers.Index)
	router.GET("/auth/hello", middleware.Auth,controllers.IndexWithAuth)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, responses.ErrorApiResponse(http.StatusNotFound, "invalid route"))
	})
}
