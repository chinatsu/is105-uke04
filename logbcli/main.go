package main

import "fmt"
import "os"
import "strconv"
import "./log"

func main() {
  if len(os.Args) != 3 {
    fmt.Println("Usage: logbcli <base> <number>")
    os.Exit(0)
  }
  base, base_err := strconv.ParseFloat(os.Args[1], 64)
  if base_err != nil {
    fmt.Println(base_err)
    os.Exit(0)
  }
  number, number_err := strconv.ParseFloat(os.Args[2], 64)
  if number_err != nil {
    fmt.Println(number_err)
    os.Exit(0)
  }
  fmt.Println(log.Logb(number, base))
}
