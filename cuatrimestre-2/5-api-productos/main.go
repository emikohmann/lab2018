package main

import (
    "github.com/gin-gonic/gin"
    productosController "github.com/emikohmann/lab2018/cuatrimestre-2/5-api-productos/controllers/productos"
)

func main() {
    router := gin.Default()

    router.GET("/productos", productosController.GetProductos)
    router.POST("/productos", productosController.PostProductos)

    router.Run(":8080")
}

