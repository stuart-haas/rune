package db

import (
	"time"

	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Hostname          string `gorm:"unique"`
	User              string
	ExternalID        string
	ExternalProvider  string
	LastSync          time.Time
	Tags             []NodeTag  `gorm:"many2many:nodes_tags"`
}

type NodeTag struct {
	gorm.Model
	Name string
}

type OAuthClient struct {
	gorm.Model
	Name          string
	Description   string
	IsConnected   bool
	ClientID      string
	ClientSecret  string
}
