package main

import (
  "fmt"
  "os"
  "github.com/flacito/bcbsup/create_machine"
)

func main() {
  argsWithoutProg := os.Args[1:]

  if (len(argsWithoutProg) == 0) {
    PrintHelp()
  }

  switch argsWithoutProg[0] {
  case "machine":
    create_machine.Process(argsWithoutProg[1:])
  default:
    PrintHelp()
  }
}

func PrintHelp() {
  fmt.Printf("ERROR: You need to pass a sub-command (e.g., bcbsup SUB-COMMAND):\n\n")
  fmt.Printf("Available subcommands: (for details, bcbsup SUB-COMMAND --help)\n\n")
  create_machine.PrintCommands()
  os.Exit(1)
}
