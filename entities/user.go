package entities

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"name"`
	BirthDay  string `gorm:"birthday"`
	Gender    string `gorm:"gender"`
	PhotoURL  string `gorm:"photo_url"`
	Time      int64  `gorm:"current_time"`
	Active    bool   `gorm:"active"`
}