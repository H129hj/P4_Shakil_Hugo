package main

func InitP4() {
	var grille [6][7]string
	var joueur1, joueur2 string
	joueur1 = "X"
	joueur2 = "O"
	var tour int
	tour = 0
	var colonne int
	var ligne int
}

func AffichageGrille(grille [6][7]string) {
	for i := 0; i < 6; i++ {
        for j := 0; j < 7; j++ {
            grille[i][j] = "."
        }
    }
}

func Victoire(grille [6][7]string, joueur string) {
	var gagnant bool
	gagnant = false
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			if grille[i][j] == joueur && grille[i][j+1] == joueur && grille[i][j+2] == joueur && grille[i][j+3] == joueur {
				gagnant = true
			}
		}
	}
	for j := 0; j < 7; j++ {
		for i := 0; i < 3; i++ {
			if grille[i][j] == joueur && grille[i+1][j] == joueur && grille[i+2][j] == joueur && grille[i+3][j] == joueur {
				gagnant = true
			}
		}
	}
}