package http

import (
	"rune/db"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterOAuthRoutes() {
	s.router.GET("/oauth/clients", func(c *gin.Context) {
		var clients []db.OAuthClient
		result := db.Client.Find(&clients)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, clients)
	})
	
	s.router.POST("/oauth/clients", func(c *gin.Context) {
		var client db.OAuthClient
		if err := c.ShouldBindJSON(&client); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Client.Create(&client)
		c.JSON(200, client)
	})
}