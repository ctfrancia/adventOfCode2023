package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
)

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

    for i := 0; i < len(s); i++ {
        //check if char is int
        if s[i] >= '0' && s[i] <= '9' {
            // convert byte to string
            fInt = string(s[i])
            break
        }
    }

    // reverse string loop over each char
    for i := len(s)-1; i >= 0; i-- {
        //check if char is int
        if s[i] >= '0' && s[i] <= '9' {
            // convert byte to string
            lInt = string(s[i])
            break
        }
    }
    sum = fInt + lInt

    if i, err := strconv.Atoi(sum); err == nil {
        sumInt = i
    }
    return sumInt
}
