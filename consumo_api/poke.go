package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// http://pokeapi.co/api/v2/pokedex/kanto/

type Response struct {
	Nome    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	numero  int            `json:"entry_number"`
	especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Nome string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(responseData))
	var responseObject Response

	if err = json.Unmarshal(responseData, &responseObject); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(responseObject.Nome)
	fmt.Println(responseObject.Pokemon)

	for _, pok := range responseObject.Pokemon {
		fmt.Println(pok.especie.Nome)
	}
}
