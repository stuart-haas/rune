package http

import (
	"rune/db"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterNodesRoutes() {
	s.router.GET("/nodes", func(c *gin.Context) {
		var nodes []db.Node
		result := db.Client.Find(&nodes)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, nodes)
	})

	s.router.POST("/nodes", func(c *gin.Context) {
		var node db.Node
		if err := c.ShouldBindJSON(&node); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Client.Create(&node)
		c.Header("X-Message", "Node created")
		c.JSON(200, node)
	})

	s.router.PUT("/nodes/:id", func(c *gin.Context) {
		id := c.Param("id")
		var node db.Node
		if err := c.ShouldBindJSON(&node); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Client.Model(&db.Node{}).Where("id = ?", id).Updates(&node)
		c.Header("X-Message", "Node updated")
		c.JSON(200, node)
	})

	s.router.DELETE("/nodes/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Client.Delete(&db.Node{}, id)
		c.Header("X-Message", "Node deleted")
		c.Status(200)
	})
}