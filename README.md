# foodShopUser

## Run 
go run main.go 

## test method 
Postman: 
- POST: localhost:3000/address/add  
  Body-raw-json:
  {
      "UserName":"tom",
      "Address":"Climenti MRT" 
  }

- GET: localhost:3000/product/address/name  (currently 1,2,3 avialiable)
  e.g: localhost:3000/product/address/jack 
  
- POST : localhost:3000/product/buy
  Body-raw-json:
  {
    "UserName":"tom",
    "Address":"Houber Front MRT" 
  }
  购买后更改成最新的地址
