package services

import (
	"github.com/guntutur/go-microservices/mvc/domain"
	"github.com/guntutur/go-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var(
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct {}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (u *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 does not exists",
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{Id:123}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}

// the conclusion here is that we try to mock a dao layer in our test case
// so that we did not have to use a call to database which is actually (well, not that actually)
// happening inside our dao layer application, outside this test case
// the idea is that we did not want to involve a database call on every test case we create
// that would exhausting the database, and outside the scope of this exercise
// only the functionality and the core algorithm should be test against our test cases
