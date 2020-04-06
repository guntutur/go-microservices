package domain

import (
	"fmt"
	"github.com/guntutur/go-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "zer0",
			LastName:  "Maverick",
			Email:     "nodata@mail.com",
		},
	}

	UserDao userDaoInterface
)

type userDao struct{}

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {

	log.Println("we're accessing the database")

	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
