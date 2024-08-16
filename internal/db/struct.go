package db

import (
	"time"
)

type ResourceField struct {
	Id            int
	FieldName     string
	ResourceRefId int
	Type          string
	IsNullable    bool
	IsArray       bool
	Default       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Resource struct {
	Id        int
	Name      string
	Slug      string
	Field     []ResourceField
	CreatedAt time.Time
	UpdatedAt time.Time
}
