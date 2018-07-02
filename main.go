package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	projectGroup := r.Group("/project")
	{
		projectGroup.GET("", allProject())
		projectGroup.GET("/:deployId", getProject())
		projectGroup.POST("", addProject())
		projectGroup.DELETE("/:deployId", deleteProject())
		projectGroup.PUT("/:deployId", updateProject())
		projectGroup.GET("/:deployId/tag", tagList())
	}

	envGroup := r.Group("/env")
	{
		envGroup.GET("", allEnv())
		envGroup.GET("/:envId", getEnv())
		envGroup.POST("", addEnv())
		envGroup.DELETE("/:envId", deleteEnv())
		envGroup.PUT("/:envId", updateEnv())
	}

	deployGroup := r.Group("/deploy")
	{
		deployGroup.GET("", allDeploy())
		deployGroup.GET("/:deployId", getDeploy())
		deployGroup.POST("", addDeploy())
		deployGroup.DELETE("/:deployId", deleteDeploy())
		deployGroup.PUT("/:deployId", updateDeploy())
		deployGroup.GET("/:deployId/env", getJobList())
		deployGroup.GET("/:deployId/env/:envId", getJobDetail())
		deployGroup.POST("/:deployId/env/:envId", addJob())
		deployGroup.PUT("/:deployId/env/:envId", doDeploy())
	}

	r.Run(":10008")
}
