package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
)

func main() {
    var total int = 0
    f, err := os.Open("./Cube_Conundrum(2)/puzzleInput.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        // cv := getFirstAndLastInt(scanner.Text())
        // calibrationValues = append(calibrationValues, cv)
    }

    fmt.Printf("total: %d\n", total)
}
