package main

import "fmt"
import "os"
import "strconv"

func main() {
  arg1, err := strconv.ParseInt(os.Args[1], 10, 64)
  arg2, err := strconv.ParseInt(os.Args[2], 10, 64)

  result := SumInt64(arg1, arg2)

  if err == nil {
    fmt.Printf("Sum of %d + %d is %d", arg1, arg2, result)
  }
}
