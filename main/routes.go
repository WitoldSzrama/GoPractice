package main

import (
	"net/http"
	"practiceTwo/handlers"

	"github.com/gin-gonic/gin"
)

func GetRoutes() {
	apiRoutes := GinApp.server.Group("/api")

	apiRoutes.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, handlers.GetSongs())
	})

	apiRoutes.GET("/:id", func(c *gin.Context) {
		song, err := handlers.GetSong(c)
		if err != nil {
			GinApp.ErrorLog.Println(err, "ID: " + c.Param("id"))
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, song)
		}
	})

	apiRoutes.PUT("/:id", func(c *gin.Context) {
		song, err := handlers.UpdateSong(c)
		if err != nil {
			GinApp.ErrorLog.Println(err, "ID: " + c.Param("id"))
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusCreated, song)
		}
	})

	apiRoutes.DELETE("/:id", func(c *gin.Context) {
		err := handlers.DeleteSong(c)
		if err != nil {
			GinApp.ErrorLog.Println(err, "ID: " + c.Param("id"))
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	})

	apiRoutes.POST("/", func(c *gin.Context) {
		song, err := handlers.AddSong(c)
		if err != nil {
			GinApp.ErrorLog.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusCreated, song)
		}
	})
}