package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlnat/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/add", controller.AddUrlNat)
	router.POST("/delete", controller.DeleteUrlNat)
	router.POST("/update", controller.UpdateUrlNat)
	router.GET("/code/*path", controller.GetUrlNat)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "start success"})
	})
	return router
}
