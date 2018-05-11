package tendencia

import "errors"

type Tendencia struct {
    Valores []float32
}

func Promedio(vs []float32) (float32, error) {
    if len(vs) == 0 {
        // La función promedio debería fallar si
        // el array está vacío
        return -1, errors.New("El Array no puede estar vacio")
    }

    var acum float32 = 0
    for _, v := range vs {
        acum += v
    }
    return acum / float32(len(vs)), nil
}

func Moda(el []int) int {
    return 0
}