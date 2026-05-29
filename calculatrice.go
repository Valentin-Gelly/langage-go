package main

import (
	"fmt"
)

var a float64
var b float64
var op string

func main() {
	for {
		fmt.Print("Entrez une valeur a : ")
		fmt.Scan(&a)
		fmt.Print("Entrez une valeur b : ")
		fmt.Scan(&b)
		fmt.Print("Entrez un opérateur (+, -, *, /, quit) : ")
		fmt.Scan(&op)
		if op == "quit" {
			break
		}
		res, err := creerOperation(op)(a, b)
		if err != nil {
			fmt.Println("Erreur : ", err)
		} else {
			fmt.Printf("Résultat de %.2f %s %.2f = %.2f\n", a, op, b, res)
		}
	}
}

func creerOperation(op string) func(float64, float64) (float64, error) {
	return func(a, b float64) (float64, error) {
		return operer(a, b, op)
	}
}

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division par zero impossible")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("operateur inconnu")
	}
}
