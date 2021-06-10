package main

// Author: yangzq80@gmail.com
// Date: 2020-02-02
//

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	rootPath := flag.String("d", "./", "-d=./")

	port := flag.String("p", "3000", "-p=3000")

	flag.Parse()

	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/health.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("/assets", *rootPath)

	fmt.Println(*port)

	r.Run(":" + *port)
}
