package models

import "gorm.io/gorm"

type UserAd struct {
	gorm.Model
	UserName string
	Address  string
}
