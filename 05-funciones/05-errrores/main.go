package main

import (
	"errors"
	"fmt"
)

func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	r, err := dividir(4, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
