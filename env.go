package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var envMap = make(map[uint]*Env, 50)
var envIndex uint = 0

func addEnv() gin.HandlerFunc {
	return func(c *gin.Context) {

		var form envForm
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1,
				"error": err.Error(),
			})
			return
		}

		envIndex++
		newEnv := new(Env)
		newEnv.Id = envIndex
		newEnv.Name = form.Name

		envMap[envIndex] = newEnv

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "add env",
			"data":    *newEnv,
		})
	}
}

func deleteEnv() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := getId(c)
		delete(envMap, id)

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "delete env",
			"data":    "",
		})
	}
}

func updateEnv() gin.HandlerFunc {
	return func(c *gin.Context) {

		var form envForm
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": err.Error(),
			})
			return
		}

		id := getId(c)
		oldEnv := envMap[id]
		if oldEnv == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "没有该环境",
			})
			return
		}

		oldEnv.Name = form.Name

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "update env",
			"data":    envMap[id],
		})
	}
}

func allEnv() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data []Env
		for _, val := range envMap {
			data = append(data, *val)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "get project list",
			"data":    data,
		})
	}
}

func getEnv() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := getId(c)
		data := envMap[id]

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "delete project",
			"data":    data,
		})
	}
}
