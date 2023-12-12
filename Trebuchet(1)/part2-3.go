package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inArray(val string, array []string) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	bytesRead, _ := os.ReadFile("./Trebuchet(1)/input.txt")
	fileContent := string(bytesRead)

	var total int = 0
	var games []string = strings.Split(fileContent, "\n")

	g := ""
	xdigits := []string{}
	for _, game := range games {
		g = ""
		xdigits = []string{}

		for _, letter := range game {
			if inArray(string(letter), nums) {
				xdigits = append(xdigits, string(letter))
				continue
			}

			g += string(letter)
			for i, word := range words {
				if strings.HasSuffix(g, word) {
					xdigits = append(xdigits, nums[i])
				}
			}
		}

		if len(xdigits) > 0 {
			first := xdigits[0]
			last := xdigits[len(xdigits)-1]

			var together string = first + last
			sum, err := strconv.Atoi(together)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}

			total += sum
		}
	}

	fmt.Printf("total: %d\n", total)
}
