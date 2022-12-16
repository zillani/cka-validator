package api

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/zillani/cka-validator/k8s/workload"
	"net/http"
)

var deployment = "webapp"
var namespace = "exam"

func Server() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/namespaces", func(c *gin.Context) {
			ns := c.Request.URL.Query().Get("name")
			resultNs := workload.GetNamespace(ns)
			c.JSON(http.StatusOK, gin.H{
				"Namespace": resultNs,
			})
		})
		api.GET("/namespaces/:id", func(c *gin.Context) {
			ns := c.Param("id")
			resultNs := workload.GetNamespace(ns)
			c.JSON(http.StatusOK, gin.H{
				"Namespace": resultNs,
			})
		})
		api.GET("/exam/:id", func(c *gin.Context) {
			qid := c.Param("id")
			var ns = "exam"
			if qid == "1" {
				log.Info("First question!")
				resultNs := workload.GetNamespace(ns)
				c.JSON(http.StatusOK, gin.H{
					"Namespace": resultNs,
				})
			}
			if qid == "2" {
				log.Info("Second question!")
				dp := workload.GetDeploy(deployment, namespace)
				c.JSON(http.StatusOK, gin.H{
					"Namespace": namespace,
					"Deploy":    dp,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"Message": "Error!",
			})
		})
	}

	// Start and run the server
	router.Run(":3000")
}
