package main

import (
	"fmt"
)

func AffichageGrille(grille [6][7]string) {
	fmt.Println("+--------------+")
	for i := 0; i < 6; i++ {
		fmt.Print("|")
		for j := 0; j < 7; j++ {
			cell := grille[i][j]
			if cell == "" {
				cell = "."
			}
			fmt.Print(cell)
			if j < 6 {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+--------------+")
	fmt.Println(" 1 2 3 4 5 6 7")
}

func AjoutPion(grille *[6][7]string, colonne int, joueur string) bool {
	for i := 5; i >= 0; i-- {
		if grille[i][colonne] == "" {
			grille[i][colonne] = joueur
			return true
		}
	}
	return false
}

func Victoire(grille [6][7]string, joueur string) bool {
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			if grille[i][j] == joueur && grille[i][j+1] == joueur && grille[i][j+2] == joueur && grille[i][j+3] == joueur {
				return true
			}
		}
	}
	for j := 0; j < 7; j++ {
		for i := 0; i < 3; i++ {
			if grille[i][j] == joueur && grille[i+1][j] == joueur && grille[i+2][j] == joueur && grille[i+3][j] == joueur {
				return true
			}
		}
	}
	return false
}

func Tourjoueur(tour *int, grille *[6][7]string, colonne int, joueur string) {
	if AjoutPion(grille, colonne, joueur) {
		(*tour)++
		if *tour%2 == 0 {
			fmt.Println("C'est au tour du joueur 2 (O)")
		} else {
			fmt.Println("C'est au tour du joueur 1 (X)")
		}
	} else {
		fmt.Println("Colonne pleine, choisissez une autre colonne")
		return
	}
}

func MatchNul(grille [6][7]string) bool {
	var plein bool
	plein = true
	for j := 0; j < 7; j++ {
		if grille[0][j] == "" {
			plein = false
		}
	}
	return plein
}

func main() {
	var grille [6][7]string
	joueurActuel := "X"
	tour := 0

	fmt.Println("Bienvenue dans Puissance 4 !")

	for {
		AffichageGrille(grille)
		fmt.Printf("Joueur %s, choisissez une colonne (1-7, 0 pour quitter): ", joueurActuel)

		var inputCol int
		fmt.Scan(&inputCol)

		if inputCol == 0 {
			fmt.Println("Vous avez quitté la partie.")
			break
		}

		if inputCol < 1 || inputCol > 7 {
			fmt.Println("Colonne invalide. Choisissez entre 1 et 7.")
			continue
		}

		colIndex := inputCol - 1
		if !AjoutPion(&grille, colIndex, joueurActuel) {
			fmt.Println("Cette colonne est pleine. Choisissez une autre.")
			continue
		}

		tour++

		if Victoire(grille, joueurActuel) {
			AffichageGrille(grille)
			if joueurActuel == "X" {
				fmt.Println("Le joueur 1 (X) a gagné !")
			} else {
				fmt.Println("Le joueur 2 (O) a gagné !")
			}
			break
		}

		if MatchNul(grille) {
			AffichageGrille(grille)
			fmt.Println("Match nul !")
			break
		}

		if joueurActuel == "X" {
			joueurActuel = "O"
		} else {
			joueurActuel = "X"
		}
	}
}
