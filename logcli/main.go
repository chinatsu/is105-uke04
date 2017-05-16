package main

import (
  "fmt"
  "os"
  "strconv"
  log "../liblog"
)

func main() {
  if len(os.Args) != 2 {
    // Hvis vi ikke har riktig antall argumenter,
    // skriver vi hvordan programmet brukes.
    fmt.Println("Usage: logcli <number>")
    os.Exit(0)
  }
  // Tolk flagg 2 som tall med desimaler (float)
  number, err := strconv.ParseFloat(os.Args[1], 64)
  if err != nil {
    // Hvis erroren fins, printer vi den
    fmt.Println(err)
  } else {
    // Hvis ikke, printer vi resultatet fra log.Log2 
    fmt.Println(log.Log2(number))
  }
}
