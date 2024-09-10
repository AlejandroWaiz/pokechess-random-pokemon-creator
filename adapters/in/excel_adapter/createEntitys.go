package excel

import (
	"math/rand"
	"strings"
	"time"

	"github.com/AlejandroWaiz/random-wild-pokemon-creator/entitys"
)

func createRandomPokemon() entitys.Wild_Pokemon{

		// Crear un nuevo generador de números aleatorios
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Elegir un índice aleatorio
		randomIndex := rng.Intn(len(AllPokemons))
	
		// Obtener el Pokémon aleatorio
		randomPokemon := AllPokemons[randomIndex]

		allMoves := strings.Join(randomPokemon.Available_moves, ", ")

    	// Separar el string por comas y crear un nuevo array de strings
    	movesArray := strings.Split(allMoves, ", ")

    	// Formatear los movimientos: eliminar guiones y capitalizar correctamente
    	for i, move := range movesArray {
    	    move = strings.ReplaceAll(move, "-", " ")
    	    movesArray[i] = strings.Title(strings.ToLower(move))
    	}

    	// Elegir hasta 6 movimientos aleatorios
    	var chosenMoves []string
    	for i := 0; i < 6 && i < len(movesArray); i++ {
    	    chosenIndex := rng.Intn(len(movesArray))
    	    chosenMoves = append(chosenMoves, movesArray[chosenIndex])
    	    // Eliminar el movimiento elegido para evitar duplicados
    	    movesArray = append(movesArray[:chosenIndex], movesArray[chosenIndex+1:]...)
    	}

		for _, chosenMove := range chosenMoves {

			for _, move := range AllMovements {

				if move.Name == chosenMove {
					randomPokemon.Moves = append(randomPokemon.Moves, move)
				}
	
			}

		}

	return randomPokemon	

}