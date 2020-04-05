package services

import (
	"github.com/guntutur/go-microservices/mvc/domain"
	"github.com/guntutur/go-microservices/mvc/utils"
	"net/http"
)

type itemService struct {}

var (
	ItemService itemService
)

func (i *itemService)GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
