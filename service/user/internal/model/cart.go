package model

import "gorm.io/gorm"

type Cart struct {
    gorm.Model
    Uid        string `gorm:"type:varchar(256);not null;uniqueIndex"`
    PidArray   string `gorm:"type:varchar(256);"`
    CountArray string `gorm:"type:varchar(128);"`
}
