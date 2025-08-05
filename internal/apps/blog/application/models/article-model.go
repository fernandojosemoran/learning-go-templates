package blog

import (
	"html/template"
	"time"
)

type Article struct {
	Id        uint `gorm:"primaryKey"`
	Title     string
	Content   template.HTML
	CreatedAt time.Time
	UpdatedAt time.Time
}
