/*
 *********************************************************************************************************
 *  dg.go
 *  Author: Michael Mabin
 *  Description:  dg.go implements the command-line interface for datagenerator.go
 *  Usage:  dg.go -fieldspecs <comma-separated fieldspec list> [-count n] [-recid] [-outfile filename]
 *      
 *  Fieldspec Information
 *  
 *     numeric:n - returns a random numeric string of n digits
 *     alpha:n - returns a random alphabetic string of n characters
 *     lastname - returns a random lastname from the list of lastNames in datagenerator
 *     firstname - returns a random firstname from the list of firstNames in datagenerator
 *     state - returns a random state code from the list of states
 *     streetname - returns a random streetname from the list of streetNames in datagenerator
 *     city - returns a random city from the list of cities in datagenerator
 *     streettype - returns a list of streettype from the list of streettypes
 *     encryptedtext -returns a randomly encrypted string
 *
 *     any fieldspec included in the list that is not recognized is treated as a literal and is returned
 *     as a field with that value.
 *
 *********************************************************************************************************
 * Maintenance Log
 * 
 * Date          Developer                Description
 * ------------  -----------------------  ----------------------------------------------------------------
 * 2015-05-02    Michael Mabin            Added these comments. Made help more descriptive
 *
 *********************************************************************************************************
 */

package main

import (
       "fmt"
       g "datagenerator"
       "flag"
       "log"
       "bytes"
       "strings"
       s "strconv"
       "os"
       "time"
)

var buf bytes.Buffer
func init() {
  log.Println("Starting dg...") 
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

func generate(fieldspecs []string,keyval_format bool,titles []string) string {
  //fmt.Printf("%q\n",fieldspecs)
  result := ""
  rec := make([]string,len(fieldspecs))

  in_fieldspec_names := map[string]bool{
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
    
  for i,spec := range fieldspecs {
     fieldspec := strings.Split(spec,":")
     var field_title string
     switch fieldspec[0] {
       case "numeric":
         num, _:= s.Atoi(fieldspec[1]) 
         rec[i] = g.Numeric(num)
       case "alpha":
         num, _:= s.Atoi(fieldspec[1]) 
         rec[i] = g.Alpha(num)
       case "lastname":
         rec[i] = g.LastName()
       case "firstname":
         rec[i] = g.FirstName()
       case "state":
         rec[i] = g.State()
       case "streetname":
         rec[i] = g.StreetName()
       case "city":
         rec[i] = g.City()
       case "streettype":
         rec[i] = g.StreetType()
       case "encryptedtext":
         rec[i] = g.EncryptedText()
       case "sha1text":
         rec[i] = g.SHA1HashText()
       case "today":
         rec[i] = current_dt()
       case "now":
         rec[i] = current_ts()
       default:
         rec[i] = fieldspec[0]
       }
       if in_fieldspec_names[fieldspec[0]] { 
         field_title = fieldspec[0]
       } else {
         field_title = fmt.Sprintf("Field%d",i + 1)
       }
       if keyval_format {
          rec[i] = fmt.Sprintf("%s=%s",field_title,rec[i])
       }
   }
   result = strings.Join(rec,"|")
   return result
}

func main() {
  //
  var fieldtype_str = flag.String("fieldspecs","none","comma-separated fieldspec list")
  var test_run = flag.Bool("test_run",false,"Run simple test of all datagenerator functions")
  var count int
  var use_rec_id = false
  var test_arg = false
  var use_fieldtitle = false
  var outfile = ""
  var title_list = ""

  flag.StringVar(&outfile,"outfile","","Output file to write generated records to")
  flag.IntVar(&count,"count",1,"Number of records to generate.")
  flag.BoolVar(&use_rec_id,"recid",false,"Include a record ID for first field.")
  flag.BoolVar(&test_arg,"with-test-arg",false,"Passed a testarg.")
  flag.BoolVar(&use_fieldtitle,"with-field-title",false,"Each field is titled in key=value format")
  flag.StringVar(&title_list,"title_list","","Comma-separated list of titles for each field")
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr,`Usage:  dg.go -fieldspecs <comma-separated fieldspec list> [-count n] [-recid] [-outfile filename]
       
   Fieldspec Information
   
   Each fieldspec argument returns a random generated string for each record.
   Some fieldspecs, like numeric and alpha, take an additonal integer n argument preceded by a colon(:) to specify length of the field.

      numeric:n - returns a random numeric string of n digits
      alpha:n - returns a random alphabetic string of n characters
      lastname - returns a random lastname from the list of lastNames in datagenerator
      firstname - returns a random firstname from the list of firstNames in datagenerator
      state - returns a random state code from the list of states
      streetname - returns a random streetname from the list of streetNames in datagenerator
      city - returns a random city from the list of cities in datagenerator
      streettype - returns a list of streettype from the list of streettypes
      encryptedtext -returns a randomly encrypted string
      current_ts - returns a current timestamp string in the format YYYY-mm-dd HH:MM:SS
      current_dt - returns a current date string in the format YYYY-mm-dd

`)
     flag.PrintDefaults()
  }
  flag.Parse()

  if test_arg {
      fmt.Println("You passed the test_arg")
      os.Exit(0)
  }

  if *test_run == true {
     test()
  } else {
     
     titles := strings.Split(title_list,",")
     if *fieldtype_str != "none" {
        fieldspecs := strings.Split(*fieldtype_str,",")
        rec_id_str := ""
        for i := 1; i <= count; i++ {
          if use_rec_id {
             rec_id_str = fmt.Sprintf("%d|",i)
          }
          rec := fmt.Sprintf("%s%s",rec_id_str,generate(fieldspecs,use_fieldtitle,titles))
          fmt.Println(rec)
        }
     }
  }


}
