// Set up server for app
package main

import "github.com/gin-gonic/gin"

func GetServer() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.SetTrustedProxies([]string{"localhost"})
	
	return engine
}