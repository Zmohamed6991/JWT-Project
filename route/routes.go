package route

import (

	"github.com/gin-gonic/gin"
	"github.com/Zmohamed6991/JWT-Project/controllers"


)

func Router() *gin.Engine {
	route := gin.Default()

	route.POST("/signup", controllers.CreateUser)
	route.POST("/login", controllers.LoginUser)
	
	route.Run(":8080")

	return route
}
