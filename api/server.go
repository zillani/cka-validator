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
			resultNs := workload.GetNamespaces()
			c.JSON(http.StatusOK, gin.H{
				"Namespaces": resultNs,
			})
		})
		api.GET("/namespaces/:id", func(c *gin.Context) {
			ns := c.Param("id")
			resultNs := workload.GetNamespace(ns)
			c.JSON(http.StatusOK, gin.H{
				"Namespace": resultNs,
			})
		})
		api.GET("/deployments", func(c *gin.Context) {
			ns := c.Query("ns")
			if ns == "" {
				ns = "default"
			}
			dp := workload.GetDeployments(ns)
			c.JSON(http.StatusOK, gin.H{
				"deployments": dp,
			})
		})
		api.GET("/deployments/:id", func(c *gin.Context) {
			depName := c.Param("id")
			ns := c.Query("ns")
			if ns == "" {
				ns = "default"
			}
			dp, replicas, _ := workload.GetDeployment(depName, ns)
			c.JSON(http.StatusOK, gin.H{
				"deployment": gin.H{
					"name":      dp,
					"namespace": ns,
					"replicas":  replicas,
				},
			})
		})
		api.GET("/exam/:id", func(c *gin.Context) {
			qid := c.Param("id")
			// Create deploy webapp from image nginx in namespace web
			if qid == "1" {
				log.Info("First question!")
				depName := "webapp"
				ns := "default"
				dp, activeReplicas, depObj := workload.GetDeployment(depName, ns)
				image := depObj.Spec.Template.Spec.Containers[0].Image
				containerName := depObj.Spec.Template.Spec.Containers[0].Name
				c.JSON(http.StatusOK, gin.H{
					"deployments": gin.H{
						"name":      dp,
						"namespace": ns,
						"container": gin.H{
							"name":  containerName,
							"image": image,
						},
						"activeReplicas": activeReplicas,
					},
				})
			} else if qid == "2" {
				log.Info("Second question!")
				dp, active, _ := workload.GetDeployment(deployment, namespace)
				c.JSON(http.StatusOK, gin.H{
					"Deploy": gin.H{
						"name":      dp,
						"namespace": namespace,
						"active":    active,
					},
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"Message": "Error!",
				})
			}

		})
	}

	// Start and run the terminal-server
	router.Run(":3000")
}
