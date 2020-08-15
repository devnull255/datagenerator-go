package datagenerator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func contains(s []string, value string) bool {
	for _, elem := range s {
		if value == elem {
			return true
		}
	}
	return false
}

func NonZeroLenCheck(check_name string, s string, t *testing.T) {
	if len(s) == 0 {
		msg := fmt.Sprintf("%s should not be zero length!", check_name)
		t.Error(msg)
	}
}

func validStringValuesCheck(check_name string, inputValues []string, validValues []string, t *testing.T) {
	for _, elem := range inputValues {
		if !contains(validValues, elem) {
			t.Errorf("unexpected value %s in inputvalues for %s", elem, check_name)
		}
	}
}

func TestFirstName(t *testing.T) {
	first_name := FirstName()
	if first_name == "" {
		t.Errorf("FirstName should have a string!")
	}
	fmt.Printf("Firstname Test: %s", first_name)
}

func TestLastName(t *testing.T) {
	last_name := LastName()
	if len(last_name) == 0 {
		t.Error("LastName should have string!")
	}
	fmt.Printf("LastName Test: %s\n", last_name)
}

func TestNumeric(t *testing.T) {
	tst_str := Numeric(8)
	if len(tst_str) != 8 {
		t.Error("Numeric did not return string of valid length!")
	}

	//conversion below will throw an error if string is not numeric
	tst_num, _ := strconv.Atoi(tst_str)
	fmt.Printf("Numeric string converted: %d\n", tst_num)
}

func TestAlpha(t *testing.T) {
	tst_str := Alpha(8)
	letters := LowerAlpha()
	if len(tst_str) != 8 {
		t.Error("Alpha did not return string of valid length!")
	}
	for i := 0; i < len(tst_str); i++ {
		if !strings.ContainsAny(letters, tst_str[i:i+1]) {
			t.Error("Alpha did not return a string with valid alphabetic characters")
		}
		//fmt.Println(tst_str[i:i+1])
	}
	fmt.Printf("Alpha Test: %s\n", tst_str)
	fmt.Printf("Alpha Test: %s\n", letters)

}

func TestStreetName(t *testing.T) {
	streetname := StreetName()
	if len(streetname) == 0 {
		t.Error("StreetName should have string!")
	}
	fmt.Printf("StreetName Test: %s\n", streetname)

}

func TestStreetType(t *testing.T) {
	NonZeroLenCheck("StreetType", StreetType(), t)
}

func TestCity(t *testing.T) {
	NonZeroLenCheck("City", City(), t)
}

func TestState(t *testing.T) {
	NonZeroLenCheck("State", State(), t)
}

func TestMap(t *testing.T) {
	simple_input := "first_name=firstname,last_name=lastname"
	output := Map(simple_input)
	with_length_fields := "hotel_id=alpha:5,hotel_name=Marriott,poi_name=Mayo Clinic"
	fmt.Printf("Map Test: %s\n", output)
	output = Map(with_length_fields)
	fmt.Printf("Map Test: %s\n", output)
	with_address := "hotel_id=alpha:5,hotel_name=Bates Motel,address=address"
	output = Map(with_address)
	fmt.Printf("Map Test: %s\n", output)
}

func TestList(t *testing.T) {
	fname_list := List("firstname", 5)
	fmt.Println("List Test: ", fname_list)
}

func TestGetAddress(t *testing.T) {
	address := GetAddress()
	fmt.Println("GetAddress Test: ", address)
}

func TestSet(t *testing.T) {
	s1 := Set("firstname", 5)
	s2 := Set("section", 5)
	fmt.Println("Set Test: ", s1)
	fmt.Println("Set Test: ", s2)
}

func TestStates(t *testing.T) {
	expected_states := states[:]
	actual_states := States()

	if !reflect.DeepEqual(expected_states, actual_states) {
		t.Fatalf("expected_states and actual_states are not equal: %s, %s",
			expected_states, actual_states)
	}
}

func TestCities(t *testing.T) {
	expected_cities := cities[:]
	actual_cities := Cities()

	if !reflect.DeepEqual(expected_cities, actual_cities) {
		t.Fatalf("expected_cities and actual_cities are not equal: %s, %s",
			expected_cities, actual_cities)
	}
}

func TestFirstNames(t *testing.T) {
	expected_firstnames := firstNames[:]
	actual_firstnames := FirstNames()

	if !reflect.DeepEqual(expected_firstnames, actual_firstnames) {
		t.Fatalf("expected_firstnames and actual_firstnames are not equal: %s, %s",
			expected_firstnames, actual_firstnames)
	}
}

func TestLastNames(t *testing.T) {
	expected_lastnames := lastNames[:]
	actual_lastnames := LastNames()

	if !reflect.DeepEqual(expected_lastnames, actual_lastnames) {
		t.Fatalf("expected_lastnames and actual_lastnames are not equal: %s, %s",
			expected_lastnames, actual_lastnames)
	}
}

func TestStreetNames(t *testing.T) {
	expected_streetnames := streetNames[:]
	actual_streetnames := StreetNames()

	if !reflect.DeepEqual(expected_streetnames, actual_streetnames) {
		t.Fatalf("expected_streetnames, actual_streetnames are not equal: %s, %s",
			expected_streetnames, actual_streetnames)
	}
}

func TestStreetTypes(t *testing.T) {
	expected_streettypes := streetTypes[:]
	actual_streettypes := StreetTypes()

	if !reflect.DeepEqual(expected_streettypes, actual_streettypes) {
		t.Fatalf("expected_streettypes and actual_streettypes are not equal: %s, %s",
			expected_streettypes, actual_streettypes)
	}
}
