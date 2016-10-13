package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "flag"
    "strconv"
)

func main() {
    var port = flag.Int("port", 1234, "The port of the service")
    flag.Parse()
    fmt.Println(*port)

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    portStr := strconv.Itoa(*port)
    fmt.Println(portStr)
    r.Run("0.0.0.0:" + portStr)
}
