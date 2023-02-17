package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

	var hearts int
	hearts = 10

	var letterGuessedCorrectly bool
	letterGuessedCorrectly = true

	var guessingStarted bool
	guessingStarted = false

	for {
		AddSpace()
		if guessingStarted {
			if letterGuessedCorrectly {
				fmt.Println("You guessed!")
			} else {
				fmt.Println("Wrong answer!")
			}
		}
		guessingStarted = true
		fmt.Println("HEARTS [" + strconv.Itoa(hearts) + "]")
		fmt.Println("\n" + outputString + "\n")

		var answer string
		fmt.Println("Guess letter: ")
		fmt.Print("Type letter >>>")
		fmt.Scan(&answer)

		// TASK: handle some errors related to wrong letter or wrong char

		// CHECK IF LETTER IN WORD TO GUESS
		if strings.Index(pokemonToGuess, answer) > -1 {
			letterGuessedCorrectly = true
			// FUNCTION TO REPLACE UNDERSCORES WITH LETTERS IF GUESSED CORRECTLY
			// TASK: replace it with function
			for i := 0; i < len(pokemonToGuess); i++ {
				pokemonToGuessLetter := string(pokemonToGuess[i])
				if answer == pokemonToGuessLetter {
					outputString = outputString[:i*2] + answer + outputString[i*2+1:]
				}
			}
		} else {
			letterGuessedCorrectly = false
			hearts -= 1

			fmt.Println("Wrong letter!")

			if hearts == 0 {
				fmt.Print("You lost")
				time.Sleep(1)
				fmt.Print(".")
				time.Sleep(1)
				fmt.Print(".")
				time.Sleep(1)
				fmt.Print(".")
				fmt.Println("Maybe try another attempt!")
				os.Exit(0)
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
