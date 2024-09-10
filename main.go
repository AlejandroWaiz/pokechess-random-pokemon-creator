package main

import (
	"log"

	excel "github.com/AlejandroWaiz/random-wild-pokemon-creator/adapters/in/excel_adapter"
	"github.com/joho/godotenv"
)

func main() {

	log.Println(godotenv.Load())


	//TODO: crear orquestador en core y que sea este quien coordine entre los adaptadores de salida y entrada
	e := excel.New_Excel_Adapter()

	p := e.GetRandomPokemons(1)

	log.Println(p[0].Name,p[0].Moves)

}

