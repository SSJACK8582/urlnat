package router

import (
	"github.com/gin-gonic/gin"
	"urlnat/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/add", controller.AddUrlNat)
	router.POST("/delete", controller.DeleteUrlNat)
	router.POST("/update", controller.UpdateUrlNat)
	router.GET("/code/*path", controller.GetUrlNat)
	return router
}
