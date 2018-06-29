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

		deployIndex ++
		newDeploy := new(deploy)
		newDeploy.Id = deployIndex
		newDeploy.Name = form.Name
		for i := 0; i < len(form.Relations); i++ {
			relation := new(DepProRelation)
			relation.ProjectId = form.Relations[i].ProjectId
			relation.TagName = form.Relations[i].TagName
			relation.Ordering = form.Relations[i].Ordering
			newDeploy.Relations = append(newDeploy.Relations, *relation)
		}
		deployMap[deployIndex] = newDeploy

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "add deploy",
			"data":    newDeploy,
		})
		return
	}
}

func deleteDeploy() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := getId(c)
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

		id := getId(c)
		myDeploy := deployMap[id]
		if myDeploy == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "没有该部署",
			})
			return
		}

		myDeploy.Name = form.Name
		myDeploy.Relations = []DepProRelation{}
		for i := 0; i < len(form.Relations); i++ {
			relation := new(DepProRelation)
			relation.ProjectId = form.Relations[i].ProjectId
			relation.TagName = form.Relations[i].TagName
			relation.Ordering = form.Relations[i].Ordering
			myDeploy.Relations = append(myDeploy.Relations, *relation)
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

		id := getId(c)
		data := deployMap[id]

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "delete deploy",
			"data":    data,
		})
	}
}
