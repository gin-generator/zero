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

	r.Group("user")
	{
		userLogic := user.NewLogic()
		r.GET("", userLogic.Index)
		r.GET(":id", userLogic.Show)
		r.POST("", userLogic.Create)
		r.PUT(":id", userLogic.Update)
		r.DELETE(":id", userLogic.Destroy)
	}

}
