package main

import (
       "fmt"
       g "datagenerator"
       "flag"
       "log"
)

func init() {
  
}

func test() {
  //run a simple test of all datagenerator functions
  fmt.Println(g.LastName() + "," + g.FirstName() )
  fmt.Println(g.Numeric(5))
  fmt.Println(g.Alpha(11))
  fmt.Println(g.StreetName())
  fmt.Println(g.StreetType())
  fmt.Println(g.City())
  fmt.Println(g.State())
}

func main() {
}
