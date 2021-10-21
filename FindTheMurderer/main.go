package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type AssetType int

const (
	Suspect = iota
	Location
	Weapon
)

type Asset struct {
	mode        AssetType
	description string
	id          int
}

func (a Asset) String() string {
	return fmt.Sprintf("%d. %s", a.id, a.description)
}

func main() {

	var userInput string

	rand.Seed(time.Now().UnixNano())
	types := []AssetType{Suspect, Location, Weapon}

	// Data initialization, could be improved in the future
	suspects := []Asset{{Suspect, "Charles B. Abbage", 1}, {Suspect, "Donald Duck Knuth", 2}}
	locations := []Asset{{Location, "Redmond", 1}, {Location, "Palo Alto", 2}, {Location, "San Francisco", 3}}
	weapons := []Asset{{Weapon, "Peixeira", 1}, {Weapon, "DynaTAC 8000X", 2}, {Weapon, "Trezoitão", 3}}

	// Create the game answer to be guessed by players
	crime := generateAnswer(types, suspects, locations, weapons)
	log.Printf("[main] Crime: %v\n", crime)

	done := false
	for !done {

		printAssets(suspects, locations, weapons)
		fmt.Printf("Your guess:")
		_, _ = fmt.Scanf("%s", &userInput)
		guess, err := parseUserGuess(userInput, types)

		if err == nil {
			guessResponse := processGuess(guess, crime)
			if guessResponse == 0 {
				done = true
				fmt.Println("You won!")
			} else {
				fmt.Printf("Your guess is wrong! Return: %d\n", guessResponse)
			}
		} else {
			fmt.Printf("Invalid input! %s\n", err.Error())
		}
	}
}

func generateAnswer(types []AssetType, suspects, locations, weapons []Asset) map[AssetType]Asset {

	crime := make(map[AssetType]Asset)

	for _, assetType := range types {
		switch assetType {
		case Suspect:
			crime[Suspect] = suspects[rand.Intn(len(suspects))]
		case Location:
			crime[Location] = locations[rand.Intn(len(locations))]
		case Weapon:
			crime[Weapon] = weapons[rand.Intn(len(weapons))]
		default:
			panic("Unrecognized type")
		}
	}
	// Example mock result for a test case
	//return map[AssetType]Asset{Suspect: suspects[1], Location: locations[0], Weapon: weapons[0]}
	return crime
}

func parseUserGuess(input string, types []AssetType) (map[AssetType]int, error) {

	var err error

	fields := strings.Split(input, ",")
	if len(fields) != len(types) {
		return nil, errors.New("invalid input, wrong number of fields")
	}

	result := make(map[AssetType]int)
	for i, field := range fields {
		switch i {
		case 0:
			result[Suspect], err = strconv.Atoi(field)
		case 1:
			result[Location], err = strconv.Atoi(field)
		case 2:
			result[Weapon], err = strconv.Atoi(field)
		default:
			return nil, errors.New("invalid number of fields")
		}
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func processGuess(userGuess map[AssetType]int, crime map[AssetType]Asset) int {

	var response []int

	for k, v := range crime {
		if assetId, ok := userGuess[k]; !ok {
			panic("parsed guess does not contain all fields")
		} else if v.id != assetId {
			switch k {
			case Suspect:
				response = append(response, 1)
			case Location:
				response = append(response, 2)
			case Weapon:
				response = append(response, 3)
			default:
				panic("unrecognized asset type")
			}
		}
	}
	if len(response) == 0 {
		return 0
	} else {
		rand.Seed(time.Now().UnixNano())
		return response[rand.Intn(len(response))]
	}
}

func printAssets(suspects, locations, weapons []Asset) {

	fmt.Println("Suspects ---------")
	for _, asset := range suspects {
		fmt.Println(asset)
	}
	fmt.Println("Weapons ---------")
	for _, asset := range weapons {
		fmt.Println(asset)
	}
	fmt.Println("Locations ---------")
	for _, node := range locations {
		fmt.Println(node)
	}
}
