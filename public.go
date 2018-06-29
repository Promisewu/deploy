package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func getId(c *gin.Context) (id uint) {
	str := c.Param("id")
	tmpId, _ := strconv.ParseUint(str, 10, 0)
	id = uint(tmpId)
	return id
}
