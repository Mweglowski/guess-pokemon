package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"math/rand"
	"strings"
	"os"
)

func AddSpace() {
	var space string
	for i := 0; i < 5; i++ {
		space += "\n"
	}
	fmt.Println(space)
}

func main() {
	fmt.Println("Guess a Pokemon from Kanto generation!")

	url := "http://pokeapi.co/api/v2/pokedex/kanto/"

	// FUNCTION THAT HANDLES GET HTTP REQUEST FOR RANDOM POKEMON
	response, error := http.Get(url)
	if error != nil {
		fmt.Print(error.Error())
	}

	responseData, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Print(error.Error())
	}

	type PokemonData struct {
		Name string `json:"name"`
	}
	type PokemonsList struct {
		PokemonData PokemonData `json:"pokemon_species"`
	}
	type Response struct {
		PokemonsList []PokemonsList `json:"pokemon_entries"`
	}

	var data Response
	json.Unmarshal(responseData, &data)

	pokemonsListLength := len(data.PokemonsList)
	randomNumber := rand.Intn(pokemonsListLength)

	pokemonToGuess := data.PokemonsList[randomNumber].PokemonData.Name
	fmt.Println(pokemonToGuess)

	// DEFINING OUTPUT STRING - WHAT USER SEES AT THE SCREEN (UNDERSCORES)
	var outputString string

	for i := 0; i < len(pokemonToGuess); i++ {
		letter := string(pokemonToGuess[i])

		if letter != " " {
			outputString += "_ "
		} else {
			outputString += "  "
		}
	}

	for {
		AddSpace()
		fmt.Println("\n" + outputString + "\n")

		var answer string
		fmt.Println("Guess letter: ")
		fmt.Print("Type letter >>>")
		fmt.Scan(&answer)
		
		// TASK: handle some errors related to wrong letter or wrong char
		
		// FUNCTION TO REPLACE UNDERSCORES WITH LETTERS IF GUESSED CORRECTLY
		for i := 0; i < len(pokemonToGuess); i++ {
			pokemonToGuessLetter := string(pokemonToGuess[i])
			if answer == pokemonToGuessLetter {
				outputString = outputString[: i * 2] + answer + outputString[i * 2 + 1 :]
			}
		}

		// CHECK FOR WIN (WHETHER THERE ARE ANY _ LEFT)
		if strings.Index(outputString, "_") < 0 {
			AddSpace()
			fmt.Println("\n" + outputString + "\n")
			fmt.Println("Congratulations! You have just won!")
			os.Exit(0)
		}
	}
}
