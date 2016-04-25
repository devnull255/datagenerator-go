package datagenerator

import (
   "testing"
   "strconv"
   "fmt"
   "strings"
)

func NonZeroLenCheck(check_name string,s string,t *testing.T) {
  if len(s) == 0 {
    msg := fmt.Sprintf("%s should not be zero length!",check_name)
    t.Error(msg)
  }
} 

func TestFirstName(t *testing.T) {
   first_name := FirstName()
   if first_name == "" {
      t.Errorf("FirstName should have a string!")
   }
   fmt.Printf("Firstname Test: %s",first_name)
}

func TestLastName(t *testing.T) {
  last_name := LastName()
  if len(last_name) == 0 {
     t.Error("LastName should have string!")
  }
  fmt.Printf("LastName Test: %s\n",last_name)
}

func TestNumeric(t *testing.T)  {
  tst_str := Numeric(8)
  if len(tst_str) != 8 {
    t.Error("Numeric did not return string of valid length!")
  }

  //conversion below will throw an error if string is not numeric
  tst_num,_ := strconv.Atoi(tst_str)
  fmt.Printf("Numeric string converted: %d\n",tst_num)
}

func TestAlpha(t *testing.T)  {
  tst_str := Alpha(8)
  letters := LowerAlpha()
  if len(tst_str) != 8 {
    t.Error("Alpha did not return string of valid length!")
  }
  for i := 0; i < len(tst_str);i++ {
     if ! strings.ContainsAny(letters,tst_str[i:i+1]) {
       t.Error("Alpha did not return a string with valid alphabetic characters")
     }
     //fmt.Println(tst_str[i:i+1])
  } 
  fmt.Printf("Alpha Test: %s\n",tst_str)
  fmt.Printf("Alpha Test: %s\n",letters)

}


func TestStreetName(t *testing.T) {
  streetname := StreetName()
  if len(streetname) == 0 {
     t.Error("StreetName should have string!")
  }
  fmt.Printf("StreetName Test: %s\n",streetname)

}

func TestStreetType(t *testing.T) {
  NonZeroLenCheck("StreetType",StreetType(),t)
}

func TestCity(t *testing.T)  {
  NonZeroLenCheck("City",City(),t)
}

func TestState(t *testing.T)  {
  NonZeroLenCheck("State",State(),t)
}

