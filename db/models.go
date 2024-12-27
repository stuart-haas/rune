package db

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	Hostname  string
	User      string
	PublicKey string
}

type OAuthClient struct {
	gorm.Model
	Name        string
	Description string
	IsConnected bool
	ClientID    string
	ClientSecret string
}
