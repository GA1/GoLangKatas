package main

import (
    "encoding/xml"
    "fmt"
)

func main() {

    type Stock struct {
        ProductList  []struct {
            Sku string `xml:"sku" json:"quantity"`
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

    var p Stock
    xml.Unmarshal(xmlData, &p)


    fmt.Println(p)
}