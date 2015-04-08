package datagenerator

import "testing"

func TestFirstName(t *testing.T) {
   first_name := FirstName()
   if first_name == "" {
      t.Errorf("FirstName should have a string!")
   }
}

