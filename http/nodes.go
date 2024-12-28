package http

import (
	"rune/db"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterNodesRoutes() {
	s.router.GET("/nodes", func(c *gin.Context) {
		var nodes []db.Node
		result := db.Client.Preload("Tags").Find(&nodes)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, nodes)
	})

	s.router.GET("/nodes/tags", func(c *gin.Context) {
		var nodeTags []db.NodeTag
		result := db.Client.Find(&nodeTags)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, nodeTags)
	})

	s.router.POST("/nodes", func(c *gin.Context) {
		var node db.Node
		if err := c.ShouldBindJSON(&node); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Create the node first
		if err := db.Client.Create(&node).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Then handle the associations
		if err := db.Client.Model(&node).Association("Tags").Replace(node.Tags); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.Header("X-Message", "Node created")
		c.JSON(200, node)
	})

	s.router.PUT("/nodes/:id", func(c *gin.Context) {
		id := c.Param("id")
		
		// First find the existing node
		var existingNode db.Node
		if err := db.Client.First(&existingNode, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Node not found"})
			return
		}
		
		// Bind the new data
		var updatedNode db.Node
		if err := c.ShouldBindJSON(&updatedNode); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		// Update the node and its associations
		db.Client.Model(&existingNode).Updates(updatedNode)
		if err := db.Client.Model(&existingNode).Association("Tags").Replace(updatedNode.Tags); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		
		c.Header("X-Message", "Node updated")
		c.JSON(200, existingNode)
	})

	s.router.DELETE("/nodes/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Client.Delete(&db.Node{}, id)
		c.Header("X-Message", "Node deleted")
		c.Status(200)
	})
}