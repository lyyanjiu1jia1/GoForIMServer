package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	ginEngine *gin.Engine
}

func GetServerInstance() Server {
	var server = Server{gin.Default()}
	return server
}

func (server Server) Run() {
	server.ping()

	_ = server.ginEngine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (server Server) ping() {
	resp := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}

	server.ginEngine.GET("/ping",resp)
}
