package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"net/url"
	"encoding/json"
)

func tagList() gin.HandlerFunc {
	return func(c *gin.Context) {

		repository := c.DefaultQuery("repository", "")
		name := url.PathEscape(repository)

		address := GITLAB_ADDRESS
		address += "/api/v3/projects/" + name + "/repository/tags"

		req, _ := http.NewRequest("GET", address, nil)
		req.Header.Add(PRIVATE_TOKEN_NAME, PRIVATE_TOKEN_VALUE)
		client := &http.Client{Timeout: 1e11}
		resp, _ := client.Do(req)

		s, _ := ioutil.ReadAll(resp.Body)
		str := string(s)
		var mapResult []map[string]interface{}
		json.Unmarshal([]byte(str), &mapResult)

		c.JSON(http.StatusOK, gin.H{
			"message": "delete project",
			"data":    mapResult,
		})
	}
}
