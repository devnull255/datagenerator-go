package datagenerator
// datagenerator package
// This package contains functions for generating datasets for testing and
// development.  This is the go implementation
//

import (
     r "crypto/rand"
     "crypto/rsa"
     "crypto/md5"
     "crypto/sha1"
     "math/rand"
     s "strings"
     c "strconv"
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

var in_fieldspec_names = map[string]bool{
       "numeric": true,
       "alpha": true,
       "lastname": true,
       "firstname": true,
       "state": true,
       "streetname": true,
       "city": true,
       "streettype": true,
       "encryptedtext": true,
       "sha1text": true,
       "current_dt": true,
       "current_ts": true,
}


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

func current_dt() string {
   // return current date in YYYY-MM-DD format
   const layout = "2006-01-02"
   return fmt.Sprintf(time.Now().Format(layout))
}

func current_ts() string {
  // returns a current timestamp in YYY-mm-dd HH:MM:SS
  const layout = "2006-01-02 15:04:05"
  return fmt.Sprintf(time.Now().Format(layout))
}

func getData(fieldspec []string) string {
     result := ""
     switch fieldspec[0] {
       case "numeric":
         num, _:= c.Atoi(fieldspec[1])
         result = Numeric(num)
       case "alpha":
         num, _:= c.Atoi(fieldspec[1])
         result = Alpha(num)
       case "lastname":
         result = LastName()
       case "firstname":
         result = FirstName()
       case "state":
         result = State()
       case "streetname":
         result = StreetName()
       case "city":
         result = City()
       case "streettype":
         result = StreetType()
       case "encryptedtext":
         result = EncryptedText()
       case "sha1text":
         result = SHA1HashText()
       case "today":
         result = current_dt()
       case "now":
         result = current_ts()
       default:
         result = fieldspec[0]
      }
      return result
}

func Map(key_value_pairs string) string {
  //returns string representation of a map of a keys and values based on comma-separated list of key=value pairs
  //if value is a recognized generator-name, a randomized value for that
  // generator will be output.
  // Example:  home=address,first_name=first_name,last_name=last_name will generate output like
  // {'home': {'street': '123 Baker Street', 'city': 'Grand Forks', 'state': 'SD'},'first_name': 'Bill', 'last_name': 'Dobbs'}
  sMap := ""
  // kv_tokens := s.SplitN(key_value_pairs, ":",2)
  kv_list := s.Split(key_value_pairs, ",")
  m := make(map[string]string)
  for _,t := range(kv_list) {
     pair := s.Split(t, "=")
     fieldspec := s.Split(pair[1], ":")
     if in_fieldspec_names[fieldspec[0]] {
         m[pair[0]] = getData(fieldspec)
     } else {
         m[pair[0]] = pair[1]
     }
  } 
  sMap = fmt.Sprintf("%#v",m)
  start_brace := s.Index(sMap, "{")
  if start_brace >= 0 {
      sMap = sMap[start_brace:]
  }   
  sMap = s.ReplaceAll(sMap, "\"","'")
  return sMap
}

func List(gen_type string, count int) string {
   // returns a string representation of a list of gen_type values
   sList := ""
   return sList 
}

func Set(gen_type string, count int) string {
   // returns a string representation of a set of unique values of
   // gen_type
   sSet := ""
   return sSet
}
