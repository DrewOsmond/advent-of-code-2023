package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.Open("./data.txt")

	checkError(err)

	scanner := bufio.NewScanner(data)

	var firstLastAddedv1 int
	var firstLastAddedv2 int
	
	

	for scanner.Scan() {
		line := scanner.Text()
		// first solution
		firstLast := grabFirstLastInt(line)
		firstLastAddedv1 = firstLastAddedv1 + calculateNumber(firstLast)

		// second solution	
		firstLastAddedV2 := grabFirstLastIntV2(line)
		firstLastAddedv2 = firstLastAddedv2 + calculateNumber(firstLastAddedV2) 				
	}

	fmt.Println(firstLastAddedv1)
	fmt.Println(firstLastAddedv2)
}

func grabFirstLastInt(s string) [2]int {
	var firstLast [2]int

	for _, char := range s {
		str := string(char)
		isInt, _ := regexp.MatchString(`\d`, str)

		if isInt && firstLast[0] == 0 {
			firstLast[0] = int(char - '0')
		}

		if isInt {
			firstLast[1] = int(char - '0')
		}
		
	}

	return firstLast
}

func grabFirstLastIntV2(s string) [2]int {
	var firstLast [2]int

	wordPattern := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)
	numberPattern := regexp.MustCompile(`\d`)
	textDigits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i, val := range s {
		curr := []byte(string(val))
		if numberPattern.Match(curr) {
			// this had me tripped up on converting a rune to an int, "subtracting the value of rune '0' from any rune '0' through '9' will give you an integer 0 through 9"
			firstLast[1] = int(val - '0') 
		} else {
			subStr := s[i:]
			currentMatch := wordPattern.FindString(subStr)

			if val, has := textDigits[currentMatch]; has {
				firstLast[1] = val
			}
		}
	}

	wordOrNumber := regexp.MustCompile(`\d|(one|two|three|four|five|six|seven|eight|nine)`)
	firstMatch := wordOrNumber.FindString(s)

	if val, has := textDigits[firstMatch]; has {
		firstLast[0] = val
	} else if len(firstMatch) == 1 {
		firstLast[0] = int(firstMatch[0] - '0')
	}

	return firstLast
}


func calculateNumber(ar [2]int) int {
	digitPattern := regexp.MustCompile(`\d`)
	digits := digitPattern.FindAllString(fmt.Sprintf("%d%d", ar[0], ar[1]), -1)

	num, err := strconv.Atoi(strings.Join(digits, ""))
	checkError(err)

	return num
}
