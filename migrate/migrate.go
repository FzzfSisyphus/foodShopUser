package main

import (
	"github.com/FzzfSisyphus/foodShopUser/initializers"
	"github.com/FzzfSisyphus/foodShopUser/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.UserAd{})

}
