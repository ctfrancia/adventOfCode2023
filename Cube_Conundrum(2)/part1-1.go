package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
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
        var matchesDirty[]string = strings.Split(scanner.Text(), ":")
        var matches string = strings.TrimSpace(matchesDirty[1])
        var sets []string = strings.Split(matches, ";")
        for _, set := range sets {
            fmt.Printf("set: %s\n", set)
        }
        // fmt.Printf("set: %s\n", matches)
        // fmt.Printf("line: %s\n", line)
        // cv := getFirstAndLastInt(scanner.Text())
        // calibrationValues = append(calibrationValues, cv)
    }

    fmt.Printf("total: %d\n", total)
}
