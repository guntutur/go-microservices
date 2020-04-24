package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/guntutur/go-microservices/src/api/domain/repositories"
	"github.com/guntutur/go-microservices/src/api/services"
	"github.com/guntutur/go-microservices/src/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
