package main

import (
	"go-otp/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default Gin router with default middleware (logger and recovery middleware)
	router := gin.Default()

	app := api.Config{Router: router}

	app.Routes()
	router.Run(":8000")
}
