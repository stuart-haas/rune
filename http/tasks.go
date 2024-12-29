package http

import (
	"rune/db"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterTasksRoutes() {
	s.router.GET("/tasks", func(c *gin.Context) {
		var tasks []db.Task
		result := db.Client.Find(&tasks)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, tasks)
	})

	s.router.POST("/tasks", func(c *gin.Context) {
		var task db.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Client.Create(&task)
		c.Header("X-Message", "Task created")
		c.JSON(200, task)
	})

	s.router.PUT("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		var task db.Task
		if err := db.Client.First(&task, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Client.Save(&task)
		c.Header("X-Message", "Task updated")
		c.JSON(200, task)
	})

	s.router.DELETE("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Client.Delete(&db.Task{}, id)
		c.Header("X-Message", "Task deleted")
		c.Status(200)
	})
}