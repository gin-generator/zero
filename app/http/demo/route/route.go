package route

import (
	"github.com/gin-generator/zero/app/http/demo/logic/user"
	"github.com/gin-generator/zero/app/http/demo/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterDemoAPI(r *gin.Engine) {
	// register middlewares
	r.Use(
		middleware.Auth(),
	)

	// API 分组
	v1 := r.Group("v1")

	// 用户
	u := v1.Group("user")
	{
		userLogic := user.NewLogic()
		u.GET("", userLogic.Index)
		u.GET(":id", userLogic.Show)
		u.POST("", userLogic.Create)
		u.PUT(":id", userLogic.Update)
		u.DELETE(":id", userLogic.Destroy)
	}

}
