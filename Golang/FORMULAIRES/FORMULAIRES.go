package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Vérifie si le formulaire a été soumis
		if r.Method == "POST" {
			// Récupère les données du formulaire
			name := r.FormValue("name")
			email := r.FormValue("email")

			// Affiche les données dans la console
			fmt.Println("Nom:", name)
			fmt.Println("Email:", email)
		}

		// Affiche le formulaire HTML
		fmt.Fprintf(w, `
			<form method="post" action="/">
				<label for="name">Nom:</label>
				<input type="text" id="name" name="name"><br><br>
				<label for="email">Email:</label>
				<input type="email" id="email" name="email"><br><br>
				<input type="submit" value="Envoyer">
			</form>
		`)
	})

	// Démarre le serveur web sur le port 8080
	fmt.Println("Serveur web démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
