package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strings"
)

func tagList() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId := getUintId(c, "deployId")
		project := projectMap[projectId]
		repository := project.Repository
		name := project.Name

		cmd := exec.Command("/bin/bash", "-c", "mkdir "+name)
		cmd.Run()

		cmd2 := exec.Command("/bin/bash", "-c", "git init")
		cmd2.Dir = "./" + name
		cmd2.Run()

		cmd3 := exec.Command("/bin/bash", "-c", "git remote add origin "+repository)
		cmd3.Dir = "./" + name
		cmd3.Run()

		cmd4 := exec.Command("/bin/bash", "-c", "git fetch origin --tags")
		cmd4.Dir = "./" + name
		cmd4.Run()

		cmd5 := exec.Command("/bin/bash", "-c", "git tag")
		cmd5.Dir = "./" + name
		out, _ := cmd5.Output()

		outstring := string(out)
		arr := strings.Split(outstring, "\n")
		newTag := arr[0 : len(arr)-1]

		c.JSON(http.StatusOK, gin.H{
			"message": "tag list",
			"data":    newTag,
		})
	}
}
