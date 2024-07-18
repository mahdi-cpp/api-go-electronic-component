package model

import (
	"github.com/lib/pq"
	"time"
)

type Topic struct {
	ID          uint           `gorm:"primarykey"`
	Name        string         `gorm:"default:null"`
	Category    string         `gorm:"default:null"`
	Photos      pq.StringArray `gorm:"type:null"`
	Title       string         `gorm:"default:null"`
	Description string         `gorm:"default:null"`
	CreatedAt   time.Time
}
