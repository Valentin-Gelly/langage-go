package main

import "fmt"

const (
	_         = iota
	directeur = (100 * iota)
	salarie
	alternant
	stagiaire
)

type Employe struct {
	nom     string
	salaire int
}

func main() {
	employeSalaire := []Employe{
		{"Jean", 100},
		{"Marie", 200},
		{"Pierre", 300},
		{"Sophie", 400},
		{"Luc", 500},
		{"Émilie", 600},
	}

	slice := make([]Employe, 3, 5)

	fmt.Printf("Nombre d'employés : %d\n", len(employeSalaire))
	fmt.Printf("Nombre maximum d'employé : %d\n", cap(employeSalaire))

	for _, employe := range employeSalaire {
		fmt.Print(employe.nom)
		slice = append(slice, employe)
		switch employe.salaire {
		case directeur:
			fmt.Print("à le droit à une voiture de fonction, ")
			fallthrough
		case salarie:
			fmt.Print("à le droit à 4 jour de RTT par ans, ")
			fallthrough
		case alternant:
			fmt.Print("à le droit à une mutuelle d'entreprise, ")
			fallthrough
		case stagiaire:
			fmt.Print("à le droit aux tickets restaurant, ")
			fallthrough
		default:
			fmt.Println(" reçois un salaire.")
		}
	}
	print("Slice : ")
	for _, employe := range slice {
		fmt.Print(employe.nom, " ")
	}

	fmt.Printf("Nombre d'employés : %d\n", len(slice))
	fmt.Printf("Nombre maximum d'employé : %d\n", cap(slice))
}
