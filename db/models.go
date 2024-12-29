package db

import (
	"time"

	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	Hostname          string `gorm:"unique"`
	User              *string
	ExternalID        *string
	ExternalProvider  *string
	SyncedAt          *time.Time
	Tags             []NodeTag  `gorm:"many2many:nodes_tags"`
	Tasks            []Task     `gorm:"many2many:tasks_nodes"`
}

type NodeTag struct {
	gorm.Model
	Name string
}

type Task struct {
	gorm.Model
	Name string
	Nodes   []Node `gorm:"many2many:tasks_nodes"`
}