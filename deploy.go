package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var deployMap = make(map[uint]*deploy, 50)
var deployIndex uint = 0

func addDeploy() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form DeployForm
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": err.Error(),
			})
			return
		}

		for _, val := range deployMap {
			if val.Name == form.Name {
				c.JSON(http.StatusOK, gin.H{
					"code":    1,
					"message": "部署名称已存在",
				})
				return
			}
		}

		for i := 0; i < len(form.Relations); i++ {
			projectId := form.Relations[i].ProjectId
			_, ok := projectMap[projectId]
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    1,
					"message": "项目不存在",
				})
				return
			}
		}

		deployIndex ++

		var relations []DepProRelation
		for i := 0; i < len(form.Relations); i++ {
			relation := DepProRelation{
				ProjectId: form.Relations[i].ProjectId,
				TagName:   form.Relations[i].TagName,
				Ordering:  form.Relations[i].Ordering,
			}
			relations = append(relations, relation)
		}

		deployMap[deployIndex] = &deploy{
			Id:        deployIndex,
			Name:      form.Name,
			Relations: relations,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "add deploy",
			"data":    deployMap[deployIndex],
		})
		return
	}
}

func deleteDeploy() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := getUintId(c, "deployId")
		delete(deployMap, id)

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "delete deploy",
			"data":    "",
		})
	}
}

func updateDeploy() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form DeployForm
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": err.Error(),
			})
			return
		}

		id := getUintId(c, "deployId")
		_, ok := deployMap[id]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "没有该部署",
			})
			return
		}

		for _, val := range deployMap {
			if val.Name == form.Name && id != val.Id {
				c.JSON(http.StatusOK, gin.H{
					"code":    1,
					"message": "部署名称已存在",
				})
				return
			}
		}

		var relations []DepProRelation
		for i := 0; i < len(form.Relations); i++ {
			relation := DepProRelation{
				ProjectId: form.Relations[i].ProjectId,
				TagName:   form.Relations[i].TagName,
				Ordering:  form.Relations[i].Ordering,
			}
			relations = append(relations, relation)
		}

		deployMap[id] = &deploy{
			Id:        id,
			Name:      form.Name,
			Relations: relations,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "update deploy",
			"data":    deployMap[id],
		})
	}
}

func allDeploy() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data []deploy
		for _, val := range deployMap {
			data = append(data, *val)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "get deploy list",
			"data":    data,
		})
	}
}

func getDeploy() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := getUintId(c, "deployId")
		data := deployMap[id]

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "get deploy",
			"data":    data,
		})
	}
}
