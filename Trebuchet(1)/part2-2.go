package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	var calibrationValues []int = []int{}
	total := 0

	f, err := os.Open("./Trebuchet(1)/calibrationInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cv := getFirstAndLastInt(scanner.Text())
		calibrationValues = append(calibrationValues, cv)
	}

	// fmt.Printf("calibrationValues: %#v\n", calibrationValues)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range calibrationValues {
		total += v
	}

	fmt.Printf("total: %d\n", total)
}

func getFirstAndLastInt(s string) int {
	var (
		fInt   string
		lInt   string
		sum    string
	)

	for i, v := range numMap {
		s = strings.ReplaceAll(s, i, v)
	}
	// fmt.Printf("s: %s\n", s)
	for i := 0; i < len(s); i++ {
		//check if char is int
		if s[i] >= '0' && s[i] <= '9' {
			// convert byte to string
			fInt = string(s[i])
			break
		}
	}
	// reverse string loop over each char

	for i := len(s) - 1; i >= 0; i-- {
		//check if char is int
		if s[i] >= '0' && s[i] <= '9' {
			// convert byte to string
            fmt.Println("asdfasdfasdfa", string(s[i]))
			lInt = string(s[i])
			break
		}
	}
	sum = fInt + lInt
    // fmt.Printf("\n")
    // fmt.Printf("sum: %#v ", sum)

    sumInt, err := strconv.Atoi(sum)
    if err != nil {
        fmt.Errorf("error", err)
    }
    fmt.Printf("sumInt: %#v ", sumInt)

    return sumInt
    /*
	if i, err := strconv.Atoi(sum); err == nil {
		sumInt = i
	}
    */
	// return sumInt
}
