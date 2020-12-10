package main

import (
	"github.com/gin-gonic/gin"
	"newsfeeder/src/handler"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingGet)
	r.Run()
}
