package main

import (
	"go-src/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helloWorld",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Get1(),
		})
	})
	r.GET("/api2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Get2(),
		})
	})
	r.GET("/api3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Get3(),
		})
	})
	r.GET("/api4", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Get4(),
		})
	})
	r.GET("/api5", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Get5(),
		})
	})
	r.POST("/apipost1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Post1(),
		})
	})
	r.POST("/apipost2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Post2(),
		})
	})
	r.POST("/apipost3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Post3(),
		})
	})
	r.POST("/apipost4", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Post4(),
		})
	})
	r.POST("/apipost5", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": api.Post5(),
		})
	})
	r.Run(":3000")
}
