package user

import "gorm.io/gorm"

type Account struct {
    gorm.Model
    Uid       string `gorm:"type:varchar(128);not null;uniqueIndex"`
    Name      string `gorm:"type:varchar(128);not null"`
    Email     string `gorm:"type:varchar(128);not null;uniqueIndex"`
    Password  string `gorm:"type:varchar(128);not null"`
    AvatarUrl string `gorm:"type:varchar(255);not null"`
}
