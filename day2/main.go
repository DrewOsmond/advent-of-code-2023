package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("./data.txt")

	if err != nil {
		return
	}

	var idCount int
	var powerOfFewest int
	scanner := bufio.NewScanner(data)	
	for scanner.Scan() {
		line := scanner.Text()
		// game 1
		maxValues := gameData{id: 0, red: 12, green: 13, blue: 14}
		maxFoundCubes := parseGame(line)

		if compareValues(maxFoundCubes, maxValues) {
			idCount += maxFoundCubes.id
		}
		powerOfFewest += powerOfColors(maxFoundCubes)
	}
	
	fmt.Println(idCount)
	fmt.Println(powerOfFewest)
}

type gameData struct {
	id int
	blue int
	green int
	red int	
}

func parseGame(s string) gameData {
	intPattern := regexp.MustCompile(`\d+`)
	slicedStr := strings.Split(s, ":")
	gameInfo := slicedStr[0]
	game := gameData{}
	gameId := intPattern.FindString(gameInfo)
	id, err := strconv.Atoi(gameId)

	if err != nil {
		panic("bad data")
	}

	game.id = id
	colorNumbers := slicedStr[1]
	for _, round := range strings.Split(colorNumbers, ";") {
		for _, val := range strings.Split(round, ",") {
			data := strings.Split(val, " ")			
			found, err := strconv.Atoi(data[1])
		
			if len(data) < 2 || err != nil {
				panic(err)
			}

			color := data[2]
			if color == "blue" && game.blue < found {
				game.blue = found
			} else if color == "red" && game.red < found  {
				game.red = found
			} else if color == "green" && game.green < found {
				game.green = found
			}

			
		}

	}

	return game
}

func compareValues(data gameData, max gameData) bool {
	if data.blue > max.blue || data.green > max.green || data.red > max.red {
		return false
	}

	return true
}

func powerOfColors(game gameData) int {
	return  game.blue * game.green * game.red
}