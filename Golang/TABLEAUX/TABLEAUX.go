package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Déterminer le nombre de dimensions du tableau (entre 3 et 5)
	nbDimensions := rand.Intn(3) + 3

	// Déterminer la taille de chaque dimension (entre 2 et 5)
	dimensions := make([]int, nbDimensions)
	for i := range dimensions {
		dimensions[i] = rand.Intn(4) + 2
	}

	// Créer le tableau avec les dimensions et remplir avec des valeurs aléatoires
	tableau := makeNDimensionalArray(dimensions)
	fillArrayWithRandomValues(tableau)

	// Afficher le tableau sous forme de tableau HTML dans une page web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<html><body><table>")
		for i := 0; i < dimensions[0]; i++ {
			fmt.Fprint(w, "<tr>")
			if nbDimensions > 1 {
				afficherSousTableau(w, tableau[i], dimensions[1:], nbDimensions-1)
			} else {
				afficherCellule(w, tableau[i])
			}
			fmt.Fprint(w, "</tr>")
		}
		fmt.Fprint(w, "</table></body></html>")
	})
	fmt.Println("Serveur web démarré sur le port 8080.")
	http.ListenAndServe(":8080", nil)
}

// Crée un tableau à n dimensions avec des dimensions données
func makeNDimensionalArray(dimensions []int) interface{} {
	if len(dimensions) == 0 {
		return nil
	}
	if len(dimensions) == 1 {
		return make([]int, dimensions[0])
	}
	arr := make([]interface{}, dimensions[0])
	for i := range arr {
		arr[i] = makeNDimensionalArray(dimensions[1:])
	}
	return arr
}

// Remplit un tableau à n dimensions avec des valeurs aléatoires
func fillArrayWithRandomValues(arr interface{}) {
	switch arr := arr.(type) {
	case []int:
		for i := range arr {
			arr[i] = rand.Intn(100)
		}
	case []interface{}:
		for i := range arr {
			fillArrayWithRandomValues(arr[i])
		}
	}
}

// Affiche un sous-tableau du tableau à n dimensions sous forme de table HTML
func afficherSousTableau(w http.ResponseWriter, sousTableau interface{}, dimensions []int, nbDimensions int) {
	if nbDimensions == 1 {
		afficherCellule(w, sousTableau)
	} else {
		for i := 0; i < dimensions[0]; i++ {
			fmt.Fprint(w, "<td>")
			afficherSousTableau(w, sousTableau.([]interface{})[i], dimensions[1:], nbDimensions-1)
			fmt.Fprint(w, "</td>")
		}
	}
}

// Affiche une cellule du tableau à n dimensions sous forme de cellule HTML
func afficherCellule(w http.ResponseWriter, cellule interface{}) {
	fmt.Fprint(w, "<td>")
	fmt.Fprint(w, strconv.Itoa(cellule))
		fmt.Fprint(w, "</td>")
		}