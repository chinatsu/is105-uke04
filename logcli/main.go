package main

import "fmt"
import "os"
import "strconv"
import "./log"

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: logcli <number>")
    os.Exit(0)
  }
  number, err := strconv.ParseFloat(os.Args[1], 64)
  if err != nil {
      fmt.Println(err)
  } else {
    fmt.Println(log.Log2(number))
  }
}
