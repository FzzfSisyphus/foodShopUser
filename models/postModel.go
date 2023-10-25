package models

import "gorm.io/gorm"

type UserAd struct {
	gorm.Model
	UserId  int
	Address string
}
