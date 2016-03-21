package main

import (
	"fmt"
	. "github.com/re-pe/output"
	"log"
	"os"
)

const (
	fRunFile   = "output_test.go"
	fDir       = "_tmp/log/"
	fExt       = ".log"
	BREAKPOINT = "breakpoint"
)

var (
	flags       Flags
	logFileName string
)
	
type Tests []struct {
	function string // tested function name
	input    []interface{} // input
	expected string // expected result
}

var tests Tests

func OutputTests(){
	for _, test := range tests {
		fmt.Printf("%s: expected %s, actual ", test.function, test.expected)
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
	
// logfailo sukūrimas
	logFile, err := NewLogFile(fDir, fRunFile, fExt)
	if err != nil { return }
	logFileInfo, err := logFile.Stat()
	if err != nil { return }
	logFileName = logFileInfo.Name()
	defer logFile.Close()
	log.SetOutput(logFile)

	tests = Tests{
	  {"Debug", []interface{}{"FY.B?:%s %s", "Labas", "Rytas"}, "Labas rytas"},
	  {"Print", []interface{}{"FC.B?:%s %s", "Labas", "Rytas"}, "Labas rytas"},
	  {"Log",   []interface{}{"Laba", " ", "Naktis"}, "Laba Naktis"},
	  {"Out",   []interface{}{"FG.B?:%s %s", "Laba", "Naktis"}, "Laba Naktis"},
	}

	OutputTests()
	
	fmt.Println("\nBye bye!")
	
}