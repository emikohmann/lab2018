package main

import (
    "fmt"
    "github.com/emikohmann/lab2018/ejemplos/errores-estadistica/tendencia"
)

func main() {

    t := tendencia.Tendencia{
        Valores : []int{8, 5, 7, 5, 4, 3, 4, 4, 6, 7, 9},
    }

    prom, err := t.Promedio()
    if err != nil {
        fmt.Println("Error: ", err.Error())
        return
    }
    fmt.Println("Promedio: ", prom)
}