package main

import (
	"recaptcha/example/src/controller"
	"recaptcha/example/src/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/verify-recaptcha", controller.VerifyReCaptcha)
	r.Run()
}
