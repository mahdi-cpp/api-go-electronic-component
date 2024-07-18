package model

import (
	"github.com/lib/pq"
	"time"
)

type Products struct {
	ID           uint           `gorm:"primarykey"`
	PartNumber   string         `gorm:"default:null"`
	Price        int32          `gorm:"default:null"`
	Photos       pq.StringArray `gorm:"type:text[]"`
	Datasheet    string         `gorm:"default:null"`
	Category     string         `gorm:"default:null"`
	Manufacturer string         `gorm:"default:null"`
	Package      string         `gorm:"default:SOP8"`
	Description  string         `gorm:"default:null"`
	Seller       string         `gorm:"default:null"`
	Url          string         `gorm:"default:null"`
	Count        int16          `gorm:"default:null"`
	CreatedAt    time.Time
}

type ProductsEntity struct {
	//Category Category
	Products []Products
	Count    int64
}
