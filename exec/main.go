package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	duck := Duck{bin: "duckdb"}

	router := gin.Default()

	// configure cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} //app.AllowOrigins
	config.AllowHeaders = []string{"Origin"}
	//config.AllowCredentials = app.AllowCredentials
	router.Use(cors.New(config))

	// Define routes
	router.POST("/exec", duck.ExecCtlr)
	//default route handler
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from datatask.io api proxy"})
	})
	router.POST("/python", ExecPythonCtlr)

	router.Run(":8080")
}
