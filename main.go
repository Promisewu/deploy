package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	projectGroup := r.Group("/project")
	{
		projectGroup.GET("", allProject())
		projectGroup.GET("/:id", getProject())
		projectGroup.POST("/add", addProject())
		projectGroup.DELETE("/:id", deleteProject())
		projectGroup.PUT("/:id", updateProject())
	}

	r.GET("/tag/list", tagList())

	deployGroup := r.Group("/deploy")
	{
		deployGroup.GET("", allDeploy())
		deployGroup.GET("/:id", getDeploy())
		deployGroup.POST("/add", addDeploy())
		deployGroup.DELETE("/:id", deleteDeploy())
		deployGroup.PUT("/:id", updateDeploy())
	}

	r.Run(":10008")
}
