package server

import (
	"github.com/gin-gonic/gin"
	"github.com/katsuomi/NotifyOfficeIlluminance/controllers"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	i := r.Group("/illuminances")
	{
		ctrl := controllers.IlluminanceController{}
		i.GET("", ctrl.Index)
		i.POST("", ctrl.Create)
	}

	return r
}
