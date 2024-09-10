package excel

import "github.com/AlejandroWaiz/random-wild-pokemon-creator/entitys"

type Excel_In_Adapter struct {
}

//TODO: mover la logica lo suficiente para que se cree la instancia de este adaptador en el archivo port

type Excel_In_Adapter_Implementation interface {
	GetRandomPokemons(amount int)[]entitys.Wild_Pokemon
}

func New_Excel_Adapter() Excel_In_Adapter_Implementation {

	e := Excel_In_Adapter{}

	createDatabase()

	return &e

}