package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ReadIDParam(c *gin.Context, key string) (int, bool) {
	keyString := c.Param(key)
	n, err := strconv.Atoi(keyString)
	if err != nil || n <= 0 {
		return 0, false
	}
	return n, true
}
