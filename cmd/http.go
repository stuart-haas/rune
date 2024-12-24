package cmd

import (
	"rune/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var httpStartCmd = &cobra.Command{
	Use:   "http",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()

    r.LoadHTMLGlob("templates/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "title": "Home Page",
        })
		})

		r.GET("/nodes", func(c *gin.Context) {
			var nodes []db.Node
			result := db.Client.Find(&nodes)
			if result.Error != nil {
				c.JSON(500, gin.H{"error": result.Error.Error()})
				return
			}
			c.JSON(200, gin.H{"nodes": nodes})
		})

		r.POST("/nodes", func(c *gin.Context) {
			var node db.Node
			if err := c.ShouldBindJSON(&node); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			db.Client.Create(&node)
			c.JSON(200, gin.H{"message": "Node created"})
		})

		r.PUT("/nodes/:id", func(c *gin.Context) {
			id := c.Param("id")
			var node db.Node
			if err := c.ShouldBindJSON(&node); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			db.Client.Model(&db.Node{}).Where("id = ?", id).Updates(&node)
			c.JSON(200, gin.H{"message": "Node updated"})
		})

		r.DELETE("/nodes/:id", func(c *gin.Context) {
			id := c.Param("id")
			db.Client.Delete(&db.Node{}, id)
			c.JSON(200, gin.H{"message": "Node deleted"})
		})

		r.Run()
	},
}

func init() {
	rootCmd.AddCommand(httpStartCmd)
}