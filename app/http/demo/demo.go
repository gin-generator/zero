package main

import (
	"fmt"
	"github.com/gin-generator/ginctl/package/get"
	"github.com/gin-generator/zero/bootstrap"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load config.
	get.NewViper("env.yaml", "./etc")

	// Start basic server.
	bootstrap.SetupLogger()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()

	// Set gin running mode, support debug, release, test.
	// release masks debugging information and is suitable for production environments.
	gin.SetMode(gin.ReleaseMode)
	// New gin.
	router := gin.New()
	// Bound route.
	bootstrap.RegisterDemoApiRoute(router)
	// Running http Services.
	log.Println("Demo api serve start: " + get.String("app.host") +
		":" + get.String("app.port"))
	err := router.Run(fmt.Sprintf("%s:%d",
		get.Get("app.host"), get.Int("app.port")))
	if err != nil {
		panic("Unable to start server, error: " + err.Error())
	}
}
