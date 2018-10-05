package productos

import (
    productosDomain "github.com/emikohmann/lab2018/cuatrimestre-2/5-api-productos/domain/productos"
    "encoding/json"
)

var (
    productsList = make([]productosDomain.Producto, 0)
)

func GetProductsList() *[]productosDomain.Producto {
    return &productsList
}

func AddProduct(body []byte) {

    var p productosDomain.Producto

    json.Unmarshal(body, &p)

    productsList = append(productsList, p)
}