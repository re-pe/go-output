package main

import (
  "fmt"
  "github.com/fatih/color"
  . "github.com/re-pe/go/output"
  "log"
  "os"
)

const (
  fRunFile   = "test.go"
  fDir       = "_tmp/log/"
  fExt       = ".log"
  BREAKPOINT = "breakpoint"
)

var (
  flags       Flags
  logFileName string
)

type CAttrs []color.Attribute

type Tests []struct {
  function string        // tested function name
  input    []interface{} // input
  expected string        // expected result
  attributes CAttrs      // attributes
}

var tests Tests

func OutputTests(){
  for _, test := range tests {
    fmt.Printf("%s: expected ", test.function)
    
    if len(test.attributes) > 0 {
        color.Set(test.attributes...)
    }
    fmt.Print(test.expected)
    if len(test.attributes) > 0 {
      color.Unset()
    }
    fmt.Print(", actual ") 
    switch (test.function){
    case "Debug":
      Debug(test.input...)
    case "Log":
      Log(test.input...)
      fmt.Print("\nFor output, look at file ", logFileName)
    case "Print":
      Print(test.input...)
    case "Out":
      Out(test.input...)
      fmt.Print("\nFor other output, look at file ", logFileName)
    }
    fmt.Println()
  }
}

func main() {
  args := os.Args
  for _, arg := range args {
    switch arg {
    case "--debug" :
            flags.Debug = true
    case "--verbose" :
      flags.Verbose = true
    }
  }
_ = BREAKPOINT  
  OuputInit(&flags)
  
  fmt.Println("\nFile", fRunFile, "is running.\n")
  
  fmt.Println("\nFlags: --debug - debugging mode.\n")
// creation of logfile
  logFile, err := NewLogFile(fDir, fRunFile, fExt)
  if err != nil { return }
  logFileInfo, err := logFile.Stat()
  if err != nil { return }
  logFileName = logFileInfo.Name()
  defer logFile.Close()
  log.SetOutput(logFile)

  tests = Tests{
    {"Debug", []interface{}{"FY.B?:%s %s", "Labas", "Rytas"}, "Labas rytas", CAttrs{color.FgYellow, color.Bold}},
    {"Print", []interface{}{"FC.B?:%s %s", "Labas", "Rytas"}, "Labas rytas", CAttrs{color.FgCyan, color.Bold}},
    {"Log",   []interface{}{"Laba", " ", "Naktis"}, "Laba Naktis", CAttrs{}},
    {"Out",   []interface{}{"FG.B?:%s %s", "Laba", "Naktis"}, "Laba Naktis", CAttrs{color.FgGreen, color.Bold}},
  }

  OutputTests()
  
  fmt.Println("\nBye bye!")
  
}