package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/guntutur/go-microservices/mvc/services"
	"github.com/guntutur/go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	// know the difference
	// path variable = c.Param("<defined in mapUrls entry>")
	// query parameter = c.Query("<defined in mapUrls entry>")
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		// Just return the Bad Request to the client.
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.ResponseError(c, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		// Handle the error and return to client
		utils.ResponseError(c, apiErr)
		return
	}

	// return user to client
	utils.Response(c, http.StatusOK, user)
}
