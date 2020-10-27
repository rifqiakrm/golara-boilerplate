package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqiakrm/golara-boilerplate/utils/responses"
	"net/http"
)

func Index (c *gin.Context){
	c.JSON(http.StatusOK, responses.SuccessApiResponseList(http.StatusOK, "success", "world"))
	return
}

func IndexWithAuth(c *gin.Context)  {
	c.JSON(http.StatusOK, responses.SuccessApiResponseList(http.StatusOK, "success", "world with auth"))
	return
}
