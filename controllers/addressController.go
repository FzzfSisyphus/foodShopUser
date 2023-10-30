package controllers

import (
	"github.com/FzzfSisyphus/foodShopUser/initializers"
	"github.com/FzzfSisyphus/foodShopUser/models"
	"github.com/gin-gonic/gin"
)

func AddNewAddr(c *gin.Context) {
	// get data off req body
	var httpBody struct {
		UserName string
		Address  string
	}
	c.Bind(&httpBody)
	// create an address
	addr := models.UserAd{UserName: httpBody.UserName, Address: httpBody.Address}
	result := initializers.DB.Create(&addr)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"userAddress": addr,
	})
}

func GetByUserNameAddress(c *gin.Context) {
	//get the post
	id := c.Param("userName")
	var addrReturn models.UserAd
	res := initializers.DB.Where("user_name = ?", id).Find(&addrReturn)
	if res.Error != nil {
		c.Status(400)
		return
	}
	//respond with them
	c.JSON(200, gin.H{
		"UserName": addrReturn.UserName,
		"Address":  addrReturn.Address,
	})
}

// func AddressShow(c *gin.Context) {
// 	//Get id off url
// 	id := c.Param("id")
// 	//get the post
// 	var addrReturn models.Address
// 	initializers.DB.First(&addrReturn, id)
// 	//respond with them
// 	c.JSON(200, gin.H{
// 		"address": addrReturn,
// 	})
// }

// product/buy POST  Userid , address
func AddressUpdate(c *gin.Context) {

	//Get detail from req body
	var httpBody struct {
		UserName string
		Address  string
	}
	c.Bind(&httpBody)
	//Find the post were updating
	var addrUpdate models.UserAd
	initializers.DB.Where("user_name = ?", httpBody.UserName).Find(&addrUpdate)
	//update it
	initializers.DB.Model(&addrUpdate).Update("address", httpBody.Address)

	//respond with them

	//get the post
	var addrReturn models.UserAd
	initializers.DB.Where("user_name = ?", httpBody.UserName).Find(&addrReturn)
	//respond with them
	c.JSON(200, gin.H{
		"address": addrReturn,
	})
}
