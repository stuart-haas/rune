package db

import (
	"time"

	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Hostname          string    `json:"hostname"`
	User              string    `json:"user"`
	ExternalID        string    `json:"externalId"`
	ExternalProvider  string    `json:"externalProvider"`
	LastSync          time.Time `json:"lastSync"`
	Tags             []Tag     `gorm:"many2many:node_tags;" json:"tags"`
}

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

type OAuthClient struct {
	gorm.Model
	Name          string `json:"name"`
	Description   string `json:"description"`
	IsConnected   bool   `json:"isConnected"`
	ClientID      string `json:"clientId"`
	ClientSecret  string `json:"clientSecret"`
}
