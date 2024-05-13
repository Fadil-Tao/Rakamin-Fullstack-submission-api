package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakamins-pbi/final-task-pbi-rakamin-fullstack-HadadFadilah/Internals/controllers"
)
func Router() *gin.Engine{
	router := gin.Default()

	// Versioning
	v1 := router.Group("/v1/pbiapi")
	v1.GET("/healthcheck", func (ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Available")
	})
	
	// User routers
	users := v1.Group("/users")
	users.GET("/alluser",CheckAuth, controllers.GetAllUser )
	users.GET("/details/:userid", CheckAuth,controllers.GetUserDetails )
	users.POST("/register", controllers.HandleRegister)
	users.POST("/login" , controllers.HandleLogin)
	users.PUT("/update/:userid",CheckAuth, controllers.UpdateUser)
	users.DELETE("/delete/:userid", CheckAuth , controllers.DeleteUser)
	users.POST("/logout",controllers.ClearCookie)
	// Photos routers
	photos := v1.Group("/photos")
	photos.GET("/allphotos")
	photos.POST("/post",CheckAuth, controllers.AddPhoto)
	photos.PUT("/update/:photoid",CheckAuth, controllers.UpdatePhoto)
	photos.DELETE("/delete/:photoid" , CheckAuth, controllers.DeletePhoto)

	return router
}