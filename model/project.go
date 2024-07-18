package model

import (
	"github.com/lib/pq"
	"time"
)

type Project struct {
	ID          uint           `gorm:"primarykey"`
	Name        string         `gorm:"default:null"`
	Category    string         `gorm:"default:null"`
	Description string         `gorm:"default:null"`
	Photos      pq.StringArray `gorm:"type:text[]"`
	Documents   pq.StringArray `gorm:"type:text[]"`
	Youtubes    pq.StringArray `gorm:"type:text[]"`
	Components  pq.StringArray `gorm:"type:text[]"`
	Price       int32          `gorm:"default:null"`
	Count       int16          `gorm:"default:null"`
	CreatedAt   time.Time
}
