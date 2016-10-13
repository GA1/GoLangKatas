package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "flag"
    "strconv"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
    var port = flag.Int("port", 1234, "The port of the service")
    flag.Parse()
    fmt.Println(*port)

    r := gin.Default()
    r.GET("/xml", func(c *gin.Context) {
        if rand.Intn(100) < 10 {
            c.JSON(500, gin.H {
                "message": generateProducts(),
            })
        } else {
            randomLoad()
            c.JSON(200, gin.H {
                "message": generateProducts(),
            })
        }
    })

    portStr := strconv.Itoa(*port)
    fmt.Println("The port chosen is: " + portStr)
    r.Run("0.0.0.0:" + portStr)
}

func randomString(length int) string {
    const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
    result := make([]byte, length)
    for i := 0; i < length; i++ {
        result[i] = chars[rand.Intn(len(chars))]
    }
    return string(result)
}

func generateProducts() string {
    var xml = "\n\t" + ""
    var N = rand.Intn(10)
    for i := 0; i < N; i++ {
        xml = xml + "\n\t<Product>\n\t\t<sku>" + randomString(40) + "</sku>\n\t\t<quantity>42</quantity>\n\t</Product>"
    }
    return xml
}

func randomLoad() {
    var i = rand.Intn(100)
    var d = 0
    if i < 20 {
        d = rand.Intn(10)
    } else if i < 70 {
        d = 50 + rand.Intn(50)
    } else if i < 95 {
        d = 200 + rand.Intn(500)
    }
    time.Sleep(time.Duration(d) * time.Millisecond)
}
