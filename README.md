# foodShopUser

## Run 
go run main.go 

## test method 
Postman: 
- POST: localhost:3000/address/add  
  Body-raw-json:
  {
      "UserId":2,
      "Address":"Climenti MRT" 
  }

- GET: localhost:3000/product/address/3  (currently 1,2,3 avialiable)
  
- POST : localhost:3000/product/buy
  Body-raw-json:
  {
    "UserId":2,
    "Address":"Houber Front MRT" 
  }
