package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("day1/input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    count := uint64(0)
    prevLine := uint64(0)
    for scanner.Scan() {
      currLine, _ := strconv.ParseUint(scanner.Text(), 10, 64)
      if (prevLine != 0 && prevLine < currLine) {
        count++
      }
      prevLine = currLine
    }

    log.Print(count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
