package main

import (
  "fmt"
  "github.com/flacito/gosup/messages"
)

func main() {
  m := messages.NewMachineMessage{
    HostName: "tapdock",
    MemorySizeGB: 4,
    CpuCount: 2,
  }

	fmt.Printf("Processing create machine request:\n  %v\n",m.HostName)
}
