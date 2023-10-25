package main

import (
	"github.com/FzzfSisyphus/foodShopUser/controllers"
	"github.com/FzzfSisyphus/foodShopUser/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/address/add", controllers.AddNewAddr)
	//r.GET("/get", controllers.AddressGetById)
	r.GET("/product/address/:id", controllers.GetByIdAddress)
	r.POST("/product/buy ", controllers.AddressUpdate)
	//r.GET("/product/address ", controllers.AddressGetById)
	r.Run() // listen and serve on 0.0.0.0:8080
}
