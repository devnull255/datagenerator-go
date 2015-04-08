package datagenerator

import "testing"

func TestFirstName(t *testing.T) {
   first_name := FirstName()
   if first_name == "" {
      t.Errorf("FirstName should have a string!")
   }
}

func TestLastName(t *testing.T) {
  last_name := LastName()
  if len(last_name) == 0 {
     t.Error("LastName should have string!")
  }
}


