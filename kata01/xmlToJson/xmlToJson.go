package main

import (
    "encoding/xml"
    "fmt"
    "encoding/json"
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

func main() {

    resp, _ := http.Get("http://0.0.0.0:5555/xml")
    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    xmlData := []byte(body)

    done := make(chan string)

    var parser Parser

    for i := 0; i < 10; i++ {
        parser = Parser{done}
        go parser.xmlToJson(xmlData)
    }

    supu := <-done
    fmt.Println(supu)
    fmt.Println("The end")
}

