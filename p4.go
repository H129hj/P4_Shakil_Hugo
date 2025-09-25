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
	var gagnant bool
	gagnant = false
}

func AffichageGrille(grille [6][7]string) {
	for i := 0; i < 6; i++ {
        for j := 0; j < 7; j++ {
            grille[i][j] = "."
        }
    }
}
