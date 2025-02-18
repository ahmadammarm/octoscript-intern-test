package model

import (
	"time"
)

// @Description Post model info
type Post struct {
	ID        uint      `gorm:"primarykey"`
	Title     string    `json:"title" binding:"required" example:"Sepatu Running"`
	Brand     string    `json:"brand" binding:"required" example:"Nike Indonesia"`
	Platform  string    `json:"platform" binding:"required" example:"Instagram"`
	DueDate   time.Time `json:"due_date" binding:"required" example:"2025-02-21T23:59:59Z"`
	Price     float64   `json:"price" example:"400000"`
	Status    string    `json:"status" binding:"required,oneof=pending in_progress completed" example:"completed"`
}
