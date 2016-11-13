package main

import (
  "fmt"
  "os"
  "github.com/flacito/bcbsup/create_machine"
)

func main() {
  argsWithoutProg := os.Args[1:]

  create_machine.Process(argsWithoutProg)

  fmt.Printf("Processing complete\n\n")
}
