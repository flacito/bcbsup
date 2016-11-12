package main

import (
  "fmt"
  "os"
  "encoding/json"
)

type CreateMachineMessage struct {
    HostName string
    MemorySizeGB int
    CpuCount int
}

func main() {
  // 1st arg is JSON, like '{"HostName":"hallo-tapco","MemorySizeGB":2,"CpuCount":2}'
  argsWithoutProg := os.Args[1:]
  b := []byte(argsWithoutProg[0])

  var m CreateMachineMessage
  e := json.Unmarshal(b, &m)
  if e != nil {
      fmt.Printf("File error: %v\n", e)
      os.Exit(1)
  }

	fmt.Printf("Processing create machine request:\n  %v\n",m.HostName)
}
