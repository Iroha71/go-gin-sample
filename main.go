package main

import (
	"gomodtest/db"
	"gomodtest/server"
	// "gomodtest/db"
)

func main() {
	// engine := gin.Default()
	// ua := ""
	// engine.Use(func(c *gin.Context) {
	// 	ua = c.GetHeader("User-Agent")
	// 	c.Next()
	// })
	// engine.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":    "Hello, go",
	// 		"User-Agent": ua,
	// 	})
	// })
	// engine.Run(":3000")
	db.Init()
	server.Init()
	db.Close()
}
