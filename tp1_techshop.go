package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Produit struct {
	Id       int16   `json:"id"`
	Nom      string  `json:"nom"`
	Marque   string  `json:"marque"`
	Prix     float64 `json:"prix"`
	Stock    int16   `json:"stock"`
	Category string  `json:"category"`
	active   bool
}

type Catalogue struct {
	Produits []Produit `json:"produits"`
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	c.Produits = append(c.Produits, p)
	return nil
}

func (c *Catalogue) TrouverParId(id int16) (Produit, error) {
	for _, produit := range c.Produits {
		if produit.Id == id {
			return produit, nil
		}
	}
	return Produit{}, errors.New("Aucun produit ne correspond à cette id")
}

func (c *Catalogue) TrouverParCategorie(cat string) []Produit {
	produits := make([]Produit, 0)
	for _, produit := range c.Produits {
		if produit.Category == cat {
			produits = append(produits, produit)
		}
	}
	return produits
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nbProductUpdated := 0
	for i, produit := range c.TrouverParCategorie(categorie) {
		c.Produits[i].Prix -= produit.Prix * (pct / 100)
		nbProductUpdated++
	}
	return nbProductUpdated
}

func (c *Catalogue) Vendre(id int16, qte int16) error {
	for i, produit := range c.Produits {
		if produit.Id == id {
			if produit.Stock < qte {
				return errors.New("Il n'y a pas assez de stock")
			}
			c.Produits[i].Stock -= qte
			fmt.Printf("Vente de %d unités effectuée. Nouveau stock : %d\n", qte, c.Produits[i].Stock)
			return nil
		}
	}
	return errors.New("Aucun produit ne correspond à cette id")
}
func (c *Catalogue) Rapport() string {
	totalCost := 0.0
	totalStock := 0
	for _, produit := range c.Produits {
		totalCost += produit.Prix * float64(produit.Stock)
		totalStock += int(produit.Stock)
	}
	compteRendu := "nombre total de produit en stock : " + strconv.Itoa(totalStock) + ". Le stock a une valeur total de : " + strconv.FormatFloat(totalCost, 'f', 2, 64)
	return compteRendu
}

func formatProduct(p Produit) string {
	return "Nom du produit " + p.Nom + " " + p.Marque + ", prix : " + strconv.FormatFloat(p.Prix, 'f', 2, 64) + ", stock : " + strconv.Itoa(int(p.Stock))
}

var choixMenu string
var catalogue Catalogue

func main() {
	err := catalogue.AjouterProduit(Produit{
		Id:       1,
		Nom:      "MacBook Pro",
		Marque:   "Apple",
		Prix:     1299.99,
		Stock:    5,
		Category: "Ordinateurs",
		active:   true,
	})
	if err != nil {
		return
	}

	err = catalogue.AjouterProduit(Produit{
		Id:       2,
		Nom:      "ThinkPad X1",
		Marque:   "Lenovo",
		Prix:     899.99,
		Stock:    8,
		Category: "Ordinateurs",
		active:   true,
	})
	if err != nil {
		return
	}

	err = catalogue.AjouterProduit(Produit{
		Id:       3,
		Nom:      "iPhone 15",
		Marque:   "Apple",
		Prix:     799.99,
		Stock:    15,
		Category: "Smartphones",
		active:   true,
	})
	if err != nil {
		return
	}

	err = catalogue.AjouterProduit(Produit{
		Id:       4,
		Nom:      "Galaxy S24",
		Marque:   "Samsung",
		Prix:     899.99,
		Stock:    12,
		Category: "Smartphones",
		active:   true,
	})
	if err != nil {
		return
	}

	err = catalogue.AjouterProduit(Produit{
		Id:       5,
		Nom:      "AirPods Pro",
		Marque:   "Apple",
		Prix:     249.99,
		Stock:    20,
		Category: "Accessoires",
		active:   true,
	})
	if err != nil {
		return
	}
	for {
		fmt.Println("[1] Ajouter un produit, [2] Chercher, [3] Soldes, [4] Vendre, [5] Rapport, [0] Quitter")
		fmt.Scanln(&choixMenu) // Utiliser Scanln au lieu de Scan

		if choixMenu == "0" {
			break
		} else if choixMenu == "1" {
			fmt.Println("Merci de rentrer les information d'un produit")

			productName := lireString("Nom du produit : ")

			fmt.Print("Marque du produit (exemples: ")
			for _, produit := range catalogue.Produits {
				fmt.Print(produit.Marque, " ")
			}
			fmt.Println(")")
			productMarque := lireString("Marque : ")

			productPrix := lireFloat("Prix du produit : ")
			productStock := lireInt("Stock du produit : ")

			fmt.Print("Categorie du produit (exemples: ")
			for _, produit := range catalogue.Produits {
				fmt.Print(produit.Category, " ")
			}
			fmt.Println(")")
			productCategory := lireString("Categorie : ")

			newProduct := Produit{
				Id:       int16(len(catalogue.Produits)) + 1,
				Nom:      productName,
				Marque:   productMarque,
				Prix:     productPrix,
				Stock:    productStock,
				Category: productCategory,
				active:   true,
			}
			err := catalogue.AjouterProduit(newProduct)
			if err != nil {
				return
			}
			fmt.Println("Le produit a été rajouté dans le catalogue")
		} else if choixMenu == "2" {
			fmt.Println("Merci de rentrer l'identifiant d'un produit")
			productId := lireInt("Identifiant du produit : ")
			produit, err := catalogue.TrouverParId(productId)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Produit trouvé")
				fmt.Println(formatProduct(produit))
			}
		} else if choixMenu == "3" {
			fmt.Println("Merci de rentrer l'identifiant d'un produit")
			fmt.Print("Categorie du produit (exemples: ")
			for _, produit := range catalogue.Produits {
				fmt.Print(produit.Category, " ")
			}
			fmt.Println(")")
			category := lireString("Categorie : ")
			fmt.Println("Entré la reduction à réaliser")
			reduction := lireFloat(" Réduction en %: ")
			nbProductUpdated := catalogue.AppliquerReduction(category, reduction)
			if nbProductUpdated == 0 {
				fmt.Println("Aucun produit n'a été trouvé dans le catalogue avec la categorie : " + category)
			} else {
				fmt.Println(strconv.Itoa(nbProductUpdated) + "produits ont été modifié")
			}
		} else if choixMenu == "4" {
			fmt.Println("Merci de rentrer l'identifiant d'un produit")
			productId := lireInt("Identifiant du produit : ")
			produit, err := catalogue.TrouverParId(productId)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Produit trouvé")
				nbProduct := lireInt("Stock vendu : ")
				err := catalogue.Vendre(produit.Id, nbProduct)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		} else if choixMenu == "5" {
			fmt.Println(catalogue.Rapport())
		}
	}
}

func lireFloat(prompt string) float64 {
	var valeur float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&valeur)
		if err != nil {
			fmt.Println("Erreur : veuillez entrer un nombre décimal valide")
			continue
		}
		return valeur
	}
}

func lireInt(prompt string) int16 {
	var valeur int16
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&valeur)
		if err != nil {
			fmt.Println("Erreur : veuillez entrer un nombre entier valide")
			continue
		}
		return valeur
	}
}

func lireString(prompt string) string {
	var valeur string
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&valeur)
		if err != nil || valeur == "" {
			fmt.Println("Erreur : veuillez entrer une valeur valide")
			continue
		}
		return valeur
	}
}
