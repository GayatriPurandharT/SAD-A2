package main

import (
	"./vendingMachine"
	"net/http"
	"github.com/gin-gonic/gin"
)

// var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	r := gin.Default()

	// Ping test
	r.GET("/test", func(c *gin.Context) {
		e := vendingMachine.VendingMachine {
			Name: "Sam",
		}
		e.display()
		// c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
