package main

import (
       "fmt"
       g "datagenerator"
)


func main() {
  fmt.Println(g.LastName() + "," + g.FirstName() )
  fmt.Println(g.Numeric(5))
  fmt.Println(g.Alpha(11))
  fmt.Println(g.StreetName())
  fmt.Println(g.StreetType())
  fmt.Println(g.City())
  fmt.Println(g.State())
}
