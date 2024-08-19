package cmd

import (
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

    r.Run()
	},
}

func init() {
	rootCmd.AddCommand(httpStartCmd)
}