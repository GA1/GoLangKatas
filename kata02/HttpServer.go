package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "flag"
    "strconv"
    "math/rand"
    "time"
)

func main() {
    var port = flag.Int("port", 1234, "The port of the service")
    flag.Parse()
    fmt.Println(*port)

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": randomString(5),
        })
    })

    portStr := strconv.Itoa(*port)
    fmt.Println("The port chosen is: " + portStr)
    r.Run("0.0.0.0:" + portStr)
}

func randomString(length int) string {
    rand.Seed(time.Now().UTC().UnixNano())
    const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
    result := make([]byte, length)
    for i := 0; i < length; i++ {
        result[i] = chars[rand.Intn(len(chars))]
    }
    return string(result)
}
