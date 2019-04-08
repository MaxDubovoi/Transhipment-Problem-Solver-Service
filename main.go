package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"./app"
)
func SetupRouter() *gin.Engine {
	router := gin.Default()
  
	v1 := router.Group("api/v1") 
	{
	  v1.GET("/test", app.GetTest)
	  v1.POST("/compute", app.PostComputeDistribution)
	}
  
	return router
  }
 
func main() {
	router := SetupRouter()
  	router.Run(":8080")
}
