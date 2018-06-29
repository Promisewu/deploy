package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var projectMap = make(map[uint]*Project, 50)
var index uint = 0

func addProject() gin.HandlerFunc {
	return func(c *gin.Context) {

		var form projectForm
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1,
				"error": err.Error(),
			})
			return
		}

		index++
		newProject := new(Project)
		newProject.Id = index
		newProject.Name = form.Name
		newProject.Repository = form.Repository
		projectMap[index] = newProject

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "add project",
			"data":    *newProject,
		})
	}
}

func deleteProject() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := getId(c)
		delete(projectMap, id)

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "delete project",
			"data":    "",
		})
	}
}

func updateProject() gin.HandlerFunc {
	return func(c *gin.Context) {

		var form projectForm
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": err.Error(),
			})
			return
		}

		id := getId(c)
		oldProject := projectMap[id]
		if oldProject == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "没有该项目",
			})
			return
		}

		oldProject.Name = form.Name
		oldProject.Repository = form.Repository

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "update project",
			"data":    projectMap[id],
		})
	}
}

func allProject() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data []Project
		for _, val := range projectMap {
			data = append(data, *val)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "get project list",
			"data":    data,
		})
	}
}

func getProject() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := getId(c)
		data := projectMap[id]

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "delete project",
			"data":    data,
		})
	}
}
