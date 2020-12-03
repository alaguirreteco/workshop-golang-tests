package calc

import (
	"errors"
	"math"
)

func raiz(num float64) (error, float64) {
	if num < 0 {
		errorMessage := "se necesita un numero positivo"
		return errors.New(errorMessage), 0
	}
	return nil, math.Sqrt(num)
}
