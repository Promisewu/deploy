package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"time"
)

func getUintId(c *gin.Context, key string) (id uint) {
	str := c.Param(key)
	tmpId, _ := strconv.ParseUint(str, 10, 0)
	id = uint(tmpId)
	return id
}

func dateTime() (str string) {
	str = "【" + time.Now().Format("2006-01-02 15:04:05") + "】"
	return str
}
