package model

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Pid        string  `gorm:"type:varchar(255);not null;uniqueIndex"`
    Title      string  `gorm:"type:varchar(255);not null"`
    Desc       string  `gorm:"type:blob"`
    Price      float64 `gorm:"type:decimal(15,5);not null"`
    CoverUrl   string  `gorm:"type:varchar(255);not null"`
    PictureUrl string  `gorm:"type:varchar(255);not null"`
}
