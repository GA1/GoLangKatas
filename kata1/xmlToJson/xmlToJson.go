package main

import (
    "encoding/xml"
    "fmt"
    "encoding/json"
)

func main() {

    type Stock struct {
        ProductList  []struct {
            Sku string `xml:"sku" json:"sku"`
            Quantity  int `xml:"quantity" json:"quantity"`
        } `xml:"Product" json:"products"`
    }

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

    var stock Stock
    xml.Unmarshal(xmlData, &stock)
    stockJson, err := json.Marshal(stock)

    if err != nil {
        fmt.Println("Houston, we got a problem.")
    }

    fmt.Println("The xml is: " + string(xmlData))
    fmt.Println(stock)
    fmt.Println("The json is " + string(stockJson))
}