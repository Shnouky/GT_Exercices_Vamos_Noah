package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    for {
        fmt.Println("Que voulez-vous faire ?")
        fmt.Println("1. Récupérer le contenu du fichier")
        fmt.Println("2. Ajouter du texte dans le fichier")
        fmt.Println("3. Supprimer le contenu du fichier")
        fmt.Println("4. Remplacer le contenu du fichier")
        fmt.Println("5. Créer un nouveau fichier")
        fmt.Println("6. Quitter")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            content, err := ioutil.ReadFile("file.txt")
            if err != nil {
                fmt.Println("Erreur lors de la lecture du fichier :", err)
            } else {
                fmt.Println(string(content))
            }
        case 2:
            fmt.Println("Veuillez entrer le texte à ajouter :")
            scanner := bufio.NewScanner(os.Stdin)
            scanner.Scan()
            text := scanner.Text()

            file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)
            if err != nil {
                fmt.Println("Erreur lors de l'ouverture du fichier :", err)
            } else {
                _, err = file.WriteString(text)
                if err != nil {
                    fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
                } else {
                    fmt.Println("Le texte a été ajouté avec succès !")
                }
                file.Close()
            }
        case 3:
            err := ioutil.WriteFile("file.txt", []byte(""), 0644)
            if err != nil {
                fmt.Println("Erreur lors de la suppression du contenu du fichier :", err)
            } else {
                fmt.Println("Le contenu du fichier a été supprimé avec succès !")
            }
        case 4:
            fmt.Println("Veuillez entrer le nouveau contenu du fichier :")
            scanner := bufio.NewScanner(os.Stdin)
            scanner.Scan()
            text := scanner.Text()

            err := ioutil.WriteFile("file.txt", []byte(text), 0644)
            if err != nil {
                fmt.Println("Erreur lors du remplacement du contenu du fichier :", err)
            } else {
                fmt.Println("Le contenu du fichier a été remplacé avec succès !")
            }
        case 5:
            fmt.Println("Veuillez entrer le nom du nouveau fichier à créer :")
            scanner := bufio.NewScanner(os.Stdin)
            scanner.Scan()
            filename := scanner.Text()

            _, err := os.Create(filename)
            if err != nil {
                fmt.Println("Erreur lors de la création du fichier :", err)
            } else {
                fmt.Println("Le fichier", filename, "a été créé avec succès !")
            }
        case 6:
            fmt.Println("Au revoir !")
            return
        default:
            fmt.Println("Choix invalide")
        }
    }
}
