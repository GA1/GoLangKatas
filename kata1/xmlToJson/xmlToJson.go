package main

import (
    "encoding/xml"
    "fmt"
    "encoding/json"
    "time"
    "math/rand"
)

type Stock struct {
    ProductList  []struct {
        Sku string `xml:"sku" json:"sku"`
        Quantity  int `xml:"quantity" json:"quantity"`
    } `xml:"Product" json:"products"`
}

func main() {

    xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
    <ProductList>
        <Product>
            <sku>ABC123</sku>
            <quantity>2</quantity>
        </Product>
        <Product>
            <sku>ABC124</sku>
            <quantity>20</quantity>
        </Product>
    </ProductList>`)


    fmt.Println(xmlToJson(xmlData))
}

func xmlToJson(xmlData []byte) (string, error) {

    time.Sleep(time.Duration(rand.Int31n(1999)) * time.Millisecond)
    var stock Stock
    errXml := xml.Unmarshal(xmlData, &stock)
    if errXml != nil {
        return "", errXml
    }

    stockJson, errJson := json.Marshal(stock)

    if errJson != nil {
        return "", errJson
    }

    return string(stockJson), nil
}


