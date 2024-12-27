package db

import (
	"time"

	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Hostname  string
	User      string
	ExternalID string
	ExternalProvider string
	LastSync time.Time
}

type OAuthClient struct {
	gorm.Model
	Name        string
	Description string
	IsConnected bool
	ClientID    string
	ClientSecret string
}
