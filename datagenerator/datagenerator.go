package datagenerator
//datagenerator package
//This package contains functions for generating datasets for testing and 
//development.  This is the go implementation
//

import (
     r "crypto/rand"
     "crypto/rsa"
     "crypto/md5"
     "crypto/sha1"
     "math/rand"
     "time"
     "fmt"
)

/********************************************************************************
 Module Variables/Constants
 I should probably make those constants
 ********************************************************************************/
var firstNames = [12]string  {"Michael","Paul","Amy","Cheryl","Randy","Becky",
      "Vicky","David","Heidi","Richard","Joseph","Patricia"}

var lastNames = [28]string  {"Anderson","Amherst","Baines","Carlson","De Jong","Everson","Furman","Garfield","Haynes","Isaacs", "Jackson","Klopper","Lamb","Martin","Nieman","O'Doole","Prince","Smith","Quayle","Rhodes","Stark",
                        "Thomas","Uhura","Vulcan","Williams","Xavier","Yeoman","Zane"}
var streetNames = [4]string {"Pine", "Oak", "Main", "Maple"}
var streetTypes = [4]string {"St.", "Dr.", "Ave.", "Rd."}
var cities = [5]string {"Oakland", "Three Oaks", "Paradise", "Hell", "Concepcion"}
var states = [50]string {"AL","AS","AR","AK","CA","CO","CN","DE","FL","GA","HI","ID","IL","IN",
         "IA","KS","KY","LA","ME","MD","MA","MI","MN","MS","MO","MT","NE","NV","NJ","NH",
         "NM","NY","NC","ND","OH","OK","OR","PA","RI","SC","SD","TN","TX","UT","VM",
         "VA","WA","WV","WI","WY"}

func init() {
   // perform module initialization
   rand.Seed(time.Now().UnixNano())
}

func LowerAlpha() string {
   // returns the string of alphabetic characters
   p := make([]byte, 26)
   for i := range p {
       p[i] = 'a' + byte(i)
   }
   return string(p)
}

func FirstName() string {
  //returns a first name from the firstNames array
  return firstNames[rand.Intn(len(firstNames))]
}

func LastName() string {
  // returns a lastname from the lastNames array
  return lastNames[rand.Intn(len(lastNames))]
}

func Numeric(num int) string {
  //returns a random numeric strength of num length
  if num <= 0 {
    panic("Error:Numeric num parameter must be greater than 0!")
  }
  result := ""
  for i := 1; i <= num; i++ {
    result = fmt.Sprintf("%s%d",result,rand.Intn(10))
  }
  return result
}

func Alpha(num int) string {
  // returns a random alphabetic string of num length
  if num <= 0 {
    panic("Error: Alpha parameter must be greater than 0!")
  }
  letters := LowerAlpha()
  result := ""
  for i := 1; i <= num; i++ {
     result = fmt.Sprintf("%s%s",result,string(letters[rand.Intn(26)]))
  }
  return result
}

func StreetName() string {
  //returns a random streetname from streetNames array
  return streetNames[rand.Intn(len(streetNames))]
}

func StreetType() string {
  //returns a random streettype from streetTypes array
  return streetTypes[rand.Intn(len(streetTypes))]
}

func EncryptedText() string {
  //returns a string of encrypted text for randomly generated alpha 20
  privkey,err := rsa.GenerateKey(r.Reader,512)
  if err != nil {
     panic(err)
  }
  md5hash := md5.New()
  label := []byte("datagenerator")
  var publickey *rsa.PublicKey
  str := []byte(Alpha(20))
  publickey = &privkey.PublicKey

  encrypted_data,err := rsa.EncryptOAEP(md5hash,r.Reader,publickey,str,label)
  if err != nil {
     panic(err)
  }
  return string(encrypted_data)
}

func SHA1HashText() string {
  //returns an SHA1 hash string for randomly generated alpha 20
  data := []byte(Alpha(20))
  hash_text := fmt.Sprintf("%x",sha1.Sum(data))
  return hash_text 
}

func City() string {
   //returns a random city from cities array
   return cities[rand.Intn(len(cities))]
}

func State() string {
  //returns a random state from states arrau
  return states[rand.Intn(len(states))]
}

