// @title			Gin Chat API
// @version		1.0
// @description	Real-time chat application.
// @host			21.6.70.96:8080
// @BasePath		/
package router

import (
	_ "ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", service.GETIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.POST("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("FindUserByNameAndPwd", service.FindUserByNameAndPwd)

	return r
}
