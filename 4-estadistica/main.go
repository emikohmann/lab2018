package main

import (
    "fmt"
    "errors"
)

type Valores []int

func (valores *Valores) Promedio() (float32, error) {
    if len(*valores) == 0 {
        return -1, errors.New("el array no puede estar vacio")
    }
    acum := 0
    for _, v := range *valores {
        acum += v
    }
    return float32(acum) / float32(len(*valores)), nil
}

func main() {
    valores := Valores{1,2,3,3,3,4,5,6}
    
    promedio, err := valores.Promedio()
    if err != nil {
        fmt.Println("No se pudo calcular el promedio")
    } else {
        fmt.Println("Promedio de los valores: ", promedio)
    }
}