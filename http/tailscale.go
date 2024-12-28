package http

import (
	"rune/tailscale"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterTailscaleRoutes() {
	s.router.GET("/tailscale/devices", func(c *gin.Context) {
		nodes, err := tailscale.NewTailscaleClient().FetchDevices()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, nodes)
	})

	s.router.GET("/tailscale/devices/:id", func(c *gin.Context) {
		device, err := tailscale.NewTailscaleClient().FetchDevice(c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, device)
	})

	s.router.GET("/tailscale/devices/check", func(c *gin.Context) {
		devices, err := tailscale.NewTailscaleClient().CheckDevices()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, devices)
	})

	s.router.POST("/tailscale/devices/sync", func(c *gin.Context) {
		err := tailscale.NewTailscaleClient().SyncDevices()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Header("X-Message", "Devices synced")
		c.Status(200)
	})

	s.router.POST("/tailscale/devices/:id/sync", func(c *gin.Context) {
		err := tailscale.NewTailscaleClient().SyncDevice(c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Header("X-Message", "Device synced")
		c.Status(200)
	})
}