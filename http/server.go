package http

import (
	"rune/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/**/*")
	return &Server{router: router}
}

func (s *Server) Start() {
	s.router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Home Page",
		})
	})

	s.router.GET("/nodes", func(c *gin.Context) {
		var nodes []db.Node
		result := db.Client.Find(&nodes)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, gin.H{"nodes": nodes})
	})

	s.router.POST("/nodes", func(c *gin.Context) {
		var node db.Node
		if err := c.ShouldBindJSON(&node); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Client.Create(&node)
		c.JSON(200, gin.H{"message": "Node created"})
	})

	s.router.PUT("/nodes/:id", func(c *gin.Context) {
		id := c.Param("id")
		var node db.Node
		if err := c.ShouldBindJSON(&node); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Client.Model(&db.Node{}).Where("id = ?", id).Updates(&node)
		c.JSON(200, gin.H{"message": "Node updated"})
	})

	s.router.DELETE("/nodes/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Client.Delete(&db.Node{}, id)
		c.JSON(200, gin.H{"message": "Node deleted"})
	})

	s.router.Run()
}