package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Création d'un routeur HTTP
    router := http.NewServeMux()

    // Définition des différentes pages
    router.HandleFunc("/", homePage)
    router.HandleFunc("/page1", page1)
    router.HandleFunc("/page2", page2)
    router.HandleFunc("/page3", page3)

    // Définition de la page 404
    router.HandleFunc("/404", notFoundPage)

    // Définition du port d'écoute du serveur
    port := ":8080"

    // Lancement du serveur
    fmt.Printf("Serveur démarré sur le port %s\n", port)
    http.ListenAndServe(port, router)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    // Redirection vers la page 1
    http.Redirect(w, r, "/page1", http.StatusSeeOther)
}

func page1(w http.ResponseWriter, r *http.Request) {
    // Redirection vers la page 2
    http.Redirect(w, r, "/page2", http.StatusSeeOther)
}

func page2(w http.ResponseWriter, r *http.Request) {
    // Redirection vers la page 3
    http.Redirect(w, r, "/page3", http.StatusSeeOther)
}

func page3(w http.ResponseWriter, r *http.Request) {
    // Affichage du contenu de la page 3
    fmt.Fprintln(w, "Bienvenue sur la page 3 !")
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
    // Affichage d'une erreur 404
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintln(w, "Page non trouvée.")
}
