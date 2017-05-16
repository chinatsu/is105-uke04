package main

import (
  "fmt"
  "os"
  "strconv"
  log "../liblog"
)

func main() {
  if len(os.Args) != 3 {
    // Hvis vi ikke har riktig antall argumenter,
    // skriver vi hvordan programmet brukes og lukker.
    fmt.Println("Usage: logbcli <base> <number>")
    os.Exit(0)
  }

  // Tolk flagg nummer 2 og 3 som float (tall med desimaler)
  base, base_err := strconv.ParseFloat(os.Args[1], 64)
  if base_err != nil {
    // Vi kunne brukt log.Fatal her i stedet
    fmt.Println(base_err)
    os.Exit(0)
  }
  number, number_err := strconv.ParseFloat(os.Args[2], 64)
  if number_err != nil {
    fmt.Println(number_err)
    os.Exit(0)
  }
  // Print resultatet fra Logb
  fmt.Println(log.Logb(number, base))
}
