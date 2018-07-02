package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
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

		for _, val := range envMap {
			if val.Name == form.Name {
				c.JSON(http.StatusOK, gin.H{
					"code":    1,
					"message": "环境已存在",
				})
				return
			}
		}

		envIndex++
		ioutil.WriteFile("config/"+form.Name+"_config", []byte(form.Config), 0644)
		newEnv := &Env{
			Id:        envIndex,
			Name:      form.Name,
			Config:    form.Name + "_config",
			Namespace: form.Namespace,
		}

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
		id := getUintId(c, "envId")
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
		id := getUintId(c, "envId")

		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": err.Error(),
			})
			return
		}

		oldEnv, ok := envMap[id]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "没有该环境",
			})
			return
		}

		for _, val := range envMap {
			if val.Name == form.Name && val.Id != id {
				c.JSON(http.StatusOK, gin.H{
					"code":    1,
					"message": "环境已存在",
				})
				return
			}
		}

		ioutil.WriteFile("config/"+oldEnv.Name+"_config", []byte(form.Config), 0644)
		envMap[id] = &Env{
			Id:        id,
			Name:      form.Name,
			Config:    form.Name + "_config",
			Namespace: form.Namespace,
		}
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
			"message": "get env list",
			"data":    data,
		})
	}
}

func getEnv() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := getUintId(c, "envId")
		env, ok := envMap[id]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "没有该环境",
			})
			return
		}

		tmpConfig, _ := ioutil.ReadFile("config/" + env.Config)
		data := map[string]interface{}{
			"Name":      env.Name,
			"Id":        env.Id,
			"Namespace": env.Namespace,
			"Config":    string(tmpConfig),
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "get env",
			"data":    data,
		})
	}
}
