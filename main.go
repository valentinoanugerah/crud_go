package main

import (
	"github.com/gin-gonic/gin"
	
	"github.com/valentinoanugerah/crud_go/database"
)


func main(){
	database.Connect()
	r := gin.Default()
	r.Run(":8080")
}
