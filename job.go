package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var jobMap = make(map[uint]map[uint]*Job, 50)
var jobIndex uint = 0

func addJob() gin.HandlerFunc {
	return func(c *gin.Context) {
		deployId := getUintId(c, "deployId")
		envId := getUintId(c, "envId")

		jobIndex ++
		deployList := make(map[uint]*Job, 10)
		newJob := new(Job)
		newJob.Id = jobIndex
		newJob.DeployId = deployId
		newJob.EnvId = envId
		newJob.Status = StatusWaiting
		newJob.Time = dateTime()
		newJob.Log = append(newJob.Log, dateTime()+"created")
		if jobMap[deployId] != nil {
			deployList = jobMap[deployId]
		}
		deployList[envId] = newJob
		jobMap[deployId] = deployList
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "job has created",
		})
	}
}

func getJobList() gin.HandlerFunc {
	return func(c *gin.Context) {
		deployId := getUintId(c, "deployId")
		jobList := jobMap[deployId]
		var data []map[string]interface{}
		for _, val := range jobList {
			envName := envMap[val.EnvId].Name
			status := statusMap[val.Status]
			tmpData := map[string]interface{}{
				"deployId": val.DeployId,
				"envId":    val.EnvId,
				"envName":  envName,
				"time":     val.Time,
				"statusId": val.Status,
				"status":   status}
			data = append(data, tmpData)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "get job list",
			"data":    data,
		})
	}
}

func getJobDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		deployId := getUintId(c, "deployId")
		envId := getUintId(c, "envId")
		job := jobMap[deployId][envId]
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "get job list",
			"data":    job,
		})
	}
}
