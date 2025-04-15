package main 

import (
	"gin-mongo-api/configs"
	"github.com/gin-gonic/gin"
	"gin-mongo-api/routes" 
)

func main(){
	 router:=gin.Default()
	 configs.ConnectDB()

	//  router.GET("/",func(c *gin.Context){
	// 	c.JSON(200,gin.H{
	// 		"data":"hello from Gin-Gonic & mongoDB",
	// 	})
		routes.UserRoute(router)
		router.Run("localhost:6000")
	 //})
}

//user:sivalokesh084
//password:BAXHUpkGVh1pP3P9


//mongodb+srv://sivalokesh084:<db_password>@cluster0.jflbyp2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0

