package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open("day2/input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    depth := int64(0)
    horizontalPosition := int64(0)
    for scanner.Scan() {
      currentCommandLine := strings.Fields(scanner.Text())
      command := currentCommandLine[0]
      commandValue, _ := strconv.ParseInt(currentCommandLine[1], 10, 64)

      switch command {
        case "forward":
          horizontalPosition = horizontalPosition + commandValue
        case "up":
          depth -= commandValue
        case "down":
          depth += commandValue
      }
    }

    log.Printf("Depth: %d | Horizontal Position: %d | Answer: %d", depth, horizontalPosition, depth*horizontalPosition)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
