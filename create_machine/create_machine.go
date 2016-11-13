package create_machine

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

func Process(argsWithoutProg []string) {
  // 1st arg is JSON, like '{"HostName":"hallo-tapco","MemorySizeGB":2,"CpuCount":2}'
  b := []byte(argsWithoutProg[0])

  var m CreateMachineMessage
  e := json.Unmarshal(b, &m)
  if e != nil {
      fmt.Printf("File error: %v\n", e)
      os.Exit(1)
  }

  fmt.Printf("\n")
  fmt.Printf("Processing create machine request:\n")
  fmt.Printf("  you want a machine named \033[0;32m%v\033[0m that has \033[0;32m%vGB\033[0m in memory and \033[0;32m%v CPUs\033[0m.\n", m.HostName, m.MemorySizeGB, m.CpuCount)
  fmt.Printf("\n")
  fmt.Printf("\n")
}
