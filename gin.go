package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-electronic-component/api"
)

var (
	router = gin.Default()
)

// Run will instagram the server
func Run() {
	err := router.SetTrustedProxies([]string{"192.168.1.2"})
	if err != nil {
		return
	}
	router.Use(CORSMiddleware())
	getRoutes()
	router.Run(":8098")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {

	v1 := router.Group("/v1")

	routes.AddProductsRoutes(v1)
	routes.AddProjectRoutes(v1)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
