package main

import "os"
import "fmt"
import "strings"

func main() {
  name := os.Args[1:]
	fmt.Printf("Hello, %s!\n", strings.Join(name, " "))
}
