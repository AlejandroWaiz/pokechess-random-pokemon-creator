package excel

import (
	"log"
	"os"
	"regexp"

	"github.com/AlejandroWaiz/random-wild-pokemon-creator/entitys"
	"github.com/xuri/excelize/v2"
)

var AllPokemons []entitys.Wild_Pokemon
var AllMovements []entitys.Move

func createDatabase(){

	AllPokemons = createAllPokemons()

	AllMovements = createAllMovements()

}

func createAllPokemons()[]entitys.Wild_Pokemon {

	pokemon_stats_excel, err := excelize.OpenFile(os.Getenv("pokemon_stats_excelname"))

	if err != nil {
		log.Printf("Error opening stats excel: %v", err)
	}

	stats_rows, err :=  pokemon_stats_excel.GetRows("Sheet1")

	if err != nil {
		log.Printf("Error getting stats rows: %v", err)
	}

	var emptyPokemon entitys.Wild_Pokemon

	for _, row := range stats_rows {

		for valueIndex, value := range row {

			switch valueIndex {

			case 0: emptyPokemon.ID = value
				
			case 1: emptyPokemon.Name = value

			case 2: emptyPokemon.HP = value

			case 3: emptyPokemon.Attack = value

			case 4: emptyPokemon.Defense = value

			case 5: emptyPokemon.SpAttack = value

			case 6: emptyPokemon.SpDefense = value

			case 7: emptyPokemon.Speed = value

			case 8,9,10,11,12,13,14,15,16,17:

				emptyPokemon.Available_moves = append(emptyPokemon.Available_moves, value)

			}			

		}

		AllPokemons = append(AllPokemons, emptyPokemon)

	}

	return AllPokemons

}

func createAllMovements()[]entitys.Move{

	moves_rows_excel, err := excelize.OpenFile(os.Getenv("moves_stats_excelname"))

	if err != nil {
		log.Printf("Error opening stats excel: %v", err)
	}

	moves_rows, err :=  moves_rows_excel.GetRows("Sheet1")

	if err != nil {
		log.Printf("Error getting stats rows: %v", err)
	}

	var emptyMove entitys.Move

	for _, row := range moves_rows {

		for valueIndex, value := range row {

			switch valueIndex {

			case 0: emptyMove.ID = value
				
			case 1: emptyMove.Name = value

			case 2: emptyMove.MoveType = value

			case 3: emptyMove.Class = value

			case 4: emptyMove.Power = value

			case 5: emptyMove.Accuracy = value

			case 6: emptyMove.EnergyCost = value

			case 7: emptyMove.Effect = cleanSpaces(value)

			case 8: emptyMove.EffectChance = value

			}			

		}

		AllMovements = append(AllMovements, emptyMove)

	}

	return AllMovements

}

func cleanSpaces(input string) string {
    // Expresión regular para encontrar un punto seguido de espacios y una letra mayúscula
    re := regexp.MustCompile(`\.\s+([A-Z])`)
    // Reemplazar los espacios largos por un solo espacio
    result := re.ReplaceAllString(input, ". $1")
    return result
}
