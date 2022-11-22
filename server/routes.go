package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) RegisterRoutes() {
	s.Engine = gin.Default()
	s.Engine.GET("", func(c *gin.Context) {
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		c.JSON(http.StatusOK, "The server is running...")
	})
	s.Engine.GET("/healthcheck", func(c *gin.Context) {
		if ok, err := s.Service.Check(); !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

}
