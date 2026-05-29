package main

import (
	"fmt"
)

var poids float64
var taille float64
var lastname string
var firstname string

const (
	IMCMaigreur = 18.5
	IMCNormal   = 25.0
	IMCSurpoids = 30.0
)

func main() {
	fmt.Print("Entrez votre nom de famille : ")
	fmt.Scan(&lastname)
	fmt.Print("Entrez votre prénom : ")
	fmt.Scan(&firstname)
	fmt.Print("Entrez votre poid : ")
	fmt.Scan(&poids)
	fmt.Print("Entrez votre taille : ")
	fmt.Scan(&taille)
	IMC := poids / (taille * taille)
	fmt.Printf("Votre IMC est de : %2f \n", IMC)
	if IMC < IMCMaigreur {
		fmt.Println("Maigreur")
	} else if IMC >= IMCMaigreur && IMC < IMCNormal {
		fmt.Println("Normal")
	} else if IMC >= IMCNormal && IMC < IMCSurpoids {
		fmt.Println("Surpoids")
	} else {
		fmt.Println("Obésité")
	}
}
