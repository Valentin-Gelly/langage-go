package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Printf("64 puissance 8 fait : %g \n", math.Pow(64, 8))
	// date with format yyyy-mm-dd hh:mm:ss.ms
	date := time.Now()
	fmt.Printf("Nous sommes le %d %d %d et il est %d:%d\n", date.Day(), date.Month(), date.Year(), date.Hour(), date.Minute())
	// birthdate with format yyyy-mm-dd hh:mm:ss.ms
	birthDate := time.Date(2003, 2, 28, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Je suis né le %d %d %d \n", birthDate.Day(), birthDate.Month(), birthDate.Year())

	// Différence entre les deux dates
	diff := date.Sub(birthDate)
	// jours
	totalDays := int(diff.Hours() / 24)
	// Années
	years := totalDays / 365
	// Jours restants
	days := totalDays % 365

	fmt.Printf("Et j'ai %d ans et %d jours\n", years, days)
}
