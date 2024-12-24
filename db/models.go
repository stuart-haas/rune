package db

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	Hostname  string
	User      string
	PublicKey string
}