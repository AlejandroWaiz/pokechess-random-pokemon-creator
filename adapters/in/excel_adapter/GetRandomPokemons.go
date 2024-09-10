package excel

import "github.com/AlejandroWaiz/random-wild-pokemon-creator/entitys"

func (e *Excel_In_Adapter) GetRandomPokemons(amount int) []entitys.Wild_Pokemon {

	var CreatedPokemons []entitys.Wild_Pokemon

	for i := 0; i < amount; i++ {

		CreatedPokemons = append(CreatedPokemons,  createRandomPokemon())

	}

	return CreatedPokemons

}