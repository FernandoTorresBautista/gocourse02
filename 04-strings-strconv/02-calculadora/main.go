package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) (int, error) {
	arrInts := strings.Split(expresion, "+")
	var resultado int
	for _, v := range arrInts {
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, errors.New("Bad expresion, only +")
		}
		resultado += num
	}
	return resultado, nil
}

func main() {
	var expresion string
	var resultado int
	var err error

	fmt.Println("=>")
	fmt.Scanln(&expresion)

	resultado, err = sumar(expresion)
	if err != nil {
		fmt.Println("=> \n", err)
	} else {
		fmt.Println("=> \n", resultado)
	}
}
