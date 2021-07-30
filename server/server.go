package server

import (
	"gomodtest/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run(":3000")
}

func router() *gin.Engine {
	r := gin.Default()
	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
	}

	return r
}
