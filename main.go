package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	test111()
	return
	r := gin.Default()

	projectGroup := r.Group("/project")
	{
		projectGroup.GET("", allProject())
		projectGroup.GET("/:id", getProject())
		projectGroup.POST("", addProject())
		projectGroup.DELETE("/:id", deleteProject())
		projectGroup.PUT("/:id", updateProject())
	}

	r.GET("/tag/:id", tagList())

	deployGroup := r.Group("/deploy")
	{
		deployGroup.GET("", allDeploy())
		deployGroup.GET("/:id", getDeploy())
		deployGroup.POST("", addDeploy())
		deployGroup.DELETE("/:id", deleteDeploy())
		deployGroup.PUT("/:id", updateDeploy())
	}

	envGroup := r.Group("/env")
	{
		envGroup.GET("", allEnv())
		envGroup.GET("/:id", getEnv())
		envGroup.POST("", addEnv())
		envGroup.DELETE("/:id", deleteEnv())
		envGroup.PUT("/:id", updateEnv())
	}

	r.Run(":10008")
}
