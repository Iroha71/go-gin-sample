// ルーティングの設定を行う
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
	// ユーザ関連のURLとコントローラの紐づけ
	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
	}

	return r
}
