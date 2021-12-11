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
    window := make([]uint64, 0, 3) 
    for scanner.Scan() {
      currLine, _ := strconv.ParseUint(scanner.Text(), 10, 64)
      
      if (len(window) == 3) {
        prevWindowSum := window[0]+window[1]+window[2]
        currWindowSum := window[1]+window[2]+currLine
        if (currWindowSum > prevWindowSum) {
          count++
        }

        window = window[1:]
      }

      window = append(window, currLine)
    }

    log.Print(count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
