package main

import (
	"fmt"
	"strconv"
)

type Personne struct {
	prenom string
	nom    string
	age    int
	email  string
}

func (p *Personne) nomComplet() string {
	return "Prénom : " + p.prenom + ", nom : " + p.nom
}

func (p *Personne) presentation() string {
	return p.nomComplet() + ", âge : " + strconv.Itoa(p.age) + ", email : " + p.email
}

type Adresse struct {
	rue        string
	ville      string
	codePostal string
}

func (a *Adresse) format() string {
	return " Habite : " + a.rue + " " + a.ville + " " + a.codePostal
}

type Employee struct {
	Personne
	Adresse
	salaire float64
}

func (e *Employee) ficheEmploye() string {
	return e.presentation() + e.format() + " A un salaire de : " + strconv.FormatFloat(e.salaire, 'f', 2, 64)
}

func (e *Employee) augmenterSalaire(pct float64) {
	e.salaire = e.salaire + pct
}

type Etudiant struct {
	Personne
	promo   string
	moyenne float64
}

func (e Etudiant) mentionObtenue() string {
	switch {
	case e.moyenne >= 10.0 && e.moyenne < 12.0:
		return "Passable"
	case e.moyenne >= 12.0 && e.moyenne < 14.0:
		return "Assez Bien"
	case e.moyenne >= 14.0 && e.moyenne < 16.0:
		return "Bien"
	case e.moyenne >= 16.0 && e.moyenne < 18.0:
		return "Très Bien"
	case e.moyenne >= 18.0 && e.moyenne <= 20.0:
		return "Excellent"
	default:
		return "Aucune mention"
	}
}

func main() {
	etudiants := []Etudiant{}
	employes := []Employee{}

	etudiant1 := Etudiant{
		Personne: Personne{
			prenom: "Alice",
			nom:    "Dupont",
			age:    20,
			email:  "alice@example.com",
		},
		promo:   "M2",
		moyenne: 15.5,
	}
	etudiants = append(etudiants, etudiant1)

	etudiant2 := Etudiant{
		Personne: Personne{
			prenom: "Bob",
			nom:    "Martin",
			age:    21,
			email:  "bob@example.com",
		},
		promo:   "M2",
		moyenne: 12.0,
	}
	etudiants = append(etudiants, etudiant2)

	employe1 := Employee{
		Personne: Personne{
			prenom: "Caroline",
			nom:    "Bernard",
			age:    30,
			email:  "caroline@company.com",
		},
		Adresse: Adresse{
			rue:        "123 Rue de Paris",
			ville:      "Lyon",
			codePostal: "69001",
		},
		salaire: 2500.50,
	}
	employes = append(employes, employe1)

	employe2 := Employee{
		Personne: Personne{
			prenom: "David",
			nom:    "Lefebvre",
			age:    35,
			email:  "david@company.com",
		},
		Adresse: Adresse{
			rue:        "456 Avenue Lyon",
			ville:      "Paris",
			codePostal: "75001",
		},
		salaire: 3000.00,
	}
	employes = append(employes, employe2)

	fmt.Println("Etudiants:")
	for _, etudiant := range etudiants {
		fmt.Println(etudiant.presentation(), "- Mention:", etudiant.mentionObtenue())
	}

	fmt.Println("\nEmployés:")
	for _, employe := range employes {
		fmt.Println(employe.ficheEmploye())
	}

	simulationAugmentationSalaire := make([]Employee, len(employes))

	copy(simulationAugmentationSalaire, employes)
	fmt.Println("Augmentation de salaire de 500 :")
	for _, employe := range simulationAugmentationSalaire {
		employe.augmenterSalaire(500)
		fmt.Println(employe.ficheEmploye())
	}

	fmt.Println("Mais on va rester avec cette grille de salaire sinon on va perdre de l'argent :")
	for _, employe := range employes {
		fmt.Println(employe.ficheEmploye())
	}
}
