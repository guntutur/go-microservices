package app

import (
	"github.com/guntutur/go-microservices/src/api/controllers/ping"
	"github.com/guntutur/go-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.POST("/repositories", repositories.CreateRepo)
	// dummy example
	router.GET("/ping", ping.Pong)
}
