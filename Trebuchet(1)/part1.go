package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
)

func main() {
    var calibrationValues []int
    var total int
    /*
    1abc2
    pqr3stu8vwx
    a1b2c3d4e5f
    treb7uchet <--- this will be a little tricky

    In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

    given an input text file loop over each line and add the first and last
    int together and print the result
    */

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
    var fInt string
    var lInt string
    var sum string
    var sumInt int

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
