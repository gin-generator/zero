package bootstrap

import (
	demo "github.com/gin-generator/zero/app/http/demo/route"
	middlewares "github.com/gin-generator/zero/middleware"
	http "github.com/gin-generator/zero/package/respond"
	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddleware(r *gin.Engine) {
	r.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.Cors(),
		middlewares.ForceUA(),
	)
}

func RegisterDemoApiRoute(router *gin.Engine) {
	// route not found.
	http.Alert404Route(router)
	// global middleware.
	RegisterGlobalMiddleware(router)
	// Initialize route.
	demo.RegisterDemoAPI(router)
}
