package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	pong = "pong"
)

func Pong(c *gin.Context) {
	c.String(http.StatusOK, pong)
}
