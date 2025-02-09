package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialisation du routeur Gin
	router := gin.Default()

	// Définition d'une première route GET
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})

	// Démarrer le serveur sur le port 8080
	router.Run(":8080")
}
