package tendencia

import "errors"

type Tendencia struct {
    Valores []int
}

func (t *Tendencia) Promedio() (float32, error) {
    if len(t.Valores) == 0 {
        return -1, errors.New("El Array no puede estar vacio")
    }

    acum := 0
    for _, v := range t.Valores {
        acum += v
    }
    return float32(acum) / float32(len(t.Valores)), nil
}