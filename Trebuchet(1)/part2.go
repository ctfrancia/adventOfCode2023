package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
    "strings"
)

type hasNumber struct {
    firstDigitIndex struct {
        idx int
        val int
    }
    lastDigitIndex struct {
        idx int
        val int
    }
}

var numMap = map[string]int {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

func main() {
    var (
        calibrationValues []int
        total int
    )

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

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("calibrationValues: %d\n", calibrationValues)

    for _, v := range calibrationValues {
        total += v
    }

    fmt.Printf("total: %d\n", total)
}

func getFirstAndLastInt(s string) int {
    var (
        fInt string
        lInt string
        sum string
        sumInt int
    )
    // first we need to check if there are any numbers written in the string
    hasWrittenNumbers, idxs, nmbrs := stringHasNumbersWritten(s)
    // fmt.Printf("string: %s\n", s)
    // fmt.Printf("hasWrittenNumbers: %t\n", hasWrittenNumbers)
    // fmt.Printf("idxs: %d\n", idxs)
    // fmt.Printf("nmbrs: %d\n", nmbrs[0])
    // if it does not have written numbers then we can proceed as normal
    // like in part1.go
    fmt.Printf("string: %s\n", s)
    for i := 0; i < len(s); i++ {
        //check if char is int
        if s[i] >= '0' && s[i] <= '9' {
            if hasWrittenNumbers {
                // check to see if the real or written number comes first
                if idxs[0] < i {
                    // if the written number comes first, then we need to
                    // get the first digit of the written number
                    // convert byte to string
                    fInt = strconv.Itoa(nmbrs[0])
                } else {
                    fInt = string(s[i])
                }
            } else {
                // convert byte to string
                fInt = string(s[i])
                break
            }
        }
    }

    // reverse string loop over each char
    for i := len(s)-1; i >= 0; i-- {
        //check if char is int
        if s[i] >= '0' && s[i] <= '9' {
            if hasWrittenNumbers {
                // check to see if the real or written number comes first
                if idxs[1] > i {
                    // if the written number comes first, then we need to
                    // get the last digit of the written number
                    // convert byte to string
                    lInt = strconv.Itoa(nmbrs[1])
                } else {
                    // else we can proceed as normal
                    // convert byte to string
                    lInt = string(s[i])
                }
            } else {
                // convert byte to string
                lInt = string(s[i])
                break
            }
        }
    }

    sum = fInt + lInt
    // fmt.Printf("sum: %s\n", sum)

    if i, err := strconv.Atoi(sum); err == nil {
        sumInt = i
    }
    // fmt.Printf("sumInt: %d\n", sumInt)

    return sumInt
}

// stringHasNumbersWritten takes a string and returns true if the string contains
// any numbers written out in letters (one, two, three, four, five, six, seven,
// eight, nine) and returns the index of the first and last digit in the string
// if there are any numbers written out in letters. If there are no numbers written
// out in letters, it returns false and [0, 0].
func stringHasNumbersWritten(s string) (bool, [2]int, [2]int) {
    // find the first word in the string that is a number written out in letters
    // and get the index of the first and last digit in the string
    // var firstDigitIndex int
    var lowestIndex int
    var lastIndex int
    var hasWrittenNumber bool = false
    var nmbr [2]int
    var letterNumbers []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    for i, v := range letterNumbers {
        if strings.Contains(s, v) {
            hasWrittenNumber = true
            foundIndex := strings.Index(s, v)

            if foundIndex == 0 {
                lowestIndex = 0
                nmbr[0] = i+1
                break
            } else {
                lowestIndex = foundIndex
                nmbr[0] = i+1
            }
        }

        foundLastIndex := strings.LastIndex(s, v)
        if foundLastIndex > lastIndex {
            lastIndex = foundLastIndex
            nmbr[1] = i+1
        }
    }
    return hasWrittenNumber, [2]int{lowestIndex, lastIndex}, nmbr
}
