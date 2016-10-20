package main

import (
    "github.com/gin-gonic/gin"
    "encoding/xml"
    "fmt"
    "flag"
    "encoding/json"
    "strconv"
    "time"
    "math/rand"
    "net/http"
    "io/ioutil"
)

type Stock struct {
    ProductList  []struct {
        Sku string `xml:"sku" json:"sku"`
        Quantity  int `xml:"quantity" json:"quantity"`
    } `xml:"Product" json:"products"`
}

type Parser struct {
    close chan string
}

func (p Parser) xmlToJson(xmlData []byte) {

    time.Sleep(time.Duration(rand.Int31n(1999)) * time.Millisecond)
    var stock Stock
    errXml := xml.Unmarshal(xmlData, &stock)
    if errXml != nil {
        return
    }

    stockJson, errJson := json.Marshal(stock)

    if errJson != nil {
        return
    }

    p.close <- string(stockJson)
}

func callBackend() (string, error) {
    resp, err1 := http.Get("http://0.0.0.0:5555/xml")
    if err1 != nil {
        return "", err1
    }
    body, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        return "", err2
    }


    defer resp.Body.Close()

    xmlData := []byte(body)

    done := make(chan string)

    var parser Parser

    for i := 0; i < 10; i++ {
        parser = Parser{done}
        go parser.xmlToJson(xmlData)
    }

    supu := <-done
    return supu, nil
}

func main() {

    var port = flag.Int("port", 1234, "The port of the service")
    flag.Parse()
    fmt.Println(*port)

    r := gin.Default()
    r.GET("/json", func(c *gin.Context) {
        if rand.Intn(100) < 0 {
            c.JSON(500, gin.H {
                "message":"Houston, we got a problem!",
            })
        } else {
            str, err := callBackend()
            if err == nil {
                c.String(200, str)
            } else {
                c.String(500, "err")
            }

        }
    })

    portStr := strconv.Itoa(*port)
    fmt.Println("The port chosen is: " + portStr)
    r.Run("0.0.0.0:" + portStr)

    fmt.Println("!!!")
    fmt.Println(callBackend())

}

