package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "flag"
)

func main() {
    var port = flag.Int("port", 1234, "help message for flagname")
    flag.Parse()
    fmt.Println(*port)

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run("0.0.0.0:8081") // listen and server on 0.0.0.0:8080
}
