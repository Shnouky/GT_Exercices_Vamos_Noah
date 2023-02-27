package main

import (
	"errors"
	"fmt"
)

func main() {
	// Demander le nombre d'allumettes
	var n int
	fmt.Print("Nombre d'allumettes (minimum 4) : ")
	fmt.Scanln(&n)

	// Vérifier que le nombre d'allumettes est valide
	err := verifierNombreAllumettes(n)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	// Boucle de jeu
	for n > 0 {
		// Afficher le nombre d'allumettes restantes
		fmt.Println("Il reste", n, "allumettes.")

		// Demander au joueur de prendre des allumettes
		var prises int
		for {
			fmt.Print("Nombre d'allumettes à prendre (1 à 3) : ")
			fmt.Scanln(&prises)
			if prises >= 1 && prises <= 3 && prises <= n {
				break
			}
			fmt.Println("Nombre d'allumettes invalide.")
		}

		// Retirer les allumettes prises
		n -= prises

		// Vérifier si le joueur a perdu
		if n == 0 {
			fmt.Println("Vous avez perdu !")
			return
		}

		// Laisser l'ordinateur prendre des allumettes
		var prisesOrdinateur int
		if n <= 4 {
			prisesOrdinateur = n - 1
		} else {
			prisesOrdinateur = 4 - prises
		}
		fmt.Println("L'ordinateur prend", prisesOrdinateur, "allumettes.")
		n -= prisesOrdinateur

		// Vérifier si l'ordinateur a perdu
		if n == 0 {
			fmt.Println("L'ordinateur a perdu !")
			return
		}
	}
}

// Fonction pour vérifier que le nombre d'allumettes est valide
func verifierNombreAllumettes(n int) error {
	if n < 4 {
		return errors.New("le nombre d'allumettes doit être au moins 4")
	}
	return nil
}
