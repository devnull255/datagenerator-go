# datagenerator-go

## Overview

Implemented a basic datagenerator package in go.
This is really just another way for me to learn a new language while creating a useful piece of reusable code in that language.  Coding the datagenerator in any language requires me to learn core common features I need for most common programming tasks, like seeding and generating random numbers, manipulating array and list structures, string formatting and manipulation, mathematical operations, iteration, stream and file I/O, processing command-line arguments, etc.

## Files

| File | Description |
| --- | --- |
| README.md | This file |
| src/dg/dg.go | Run at the command line to generate text data|
| src/datagenerator/datagenerator.go | Contains datagenerator functions|
| src/datagenerator/datagenerator_test.go |Testing package |

## datagenerator functions

| Function | Description |
| --- | --- |
| func LowerAlpha() string | Returns lowercase alphabet |
| func FirstName() string | Returns a random firstname string from firstNames array |
| func LastName() string | Returns a random lastname string from lastNames array |
| func Numeric(num int) string | Returns a string of numbers with num length |
| func Alpha(num int) string | Returns a string of lowercase numbers with num length |
| func StreetName() string | Returns a random streetname from streetNames array |
| func StreetType() string | Returns a random streetType (like ST., AVE., BLVD ) |
| func City() string | Returns a random city from cities array |
| func State() string | Returns a random state code from array of states |

