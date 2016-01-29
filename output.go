package output

import (
	"github.com/fatih/color"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
) 

/*
	Structure of Out.selector:
	"<Function name>.<Function type>.<flags>.<Color>.<Efect>"
	<Function name>:
	fPr = fmt.Print
	lFa = log.Fatal
	lPa = log.Panic
	lPr = log.Print
	
	<Fuction type>:
	<> Or S = Plain : ftm.Print, log.Print
	Ln = function variation with ln: fmt.Println, log.Println
	F = function variation with f : fmt.Printf, log.Printf
	
	<flags>: <xx>. 
	First place means debug status, second means verbosity.
	0 means „not matter“, or „always“.
	1 means „turned off“
	2 turned „turned on“
	00 means „print always not matter what debug and verbosity status“
	20 means „print when debugging switched on, not matter what status of verbosity is“
*/


const (
	BREAKPOINT = "breakpoint"
	States     = 3
)

type PreFormat struct{
	first string
	last string
}

type PreFormatList map[string]PreFormat

type Environment struct {
	Initialized bool
	LogFile *os.File
	flagDebug bool
	flagVerbose bool
	states int
	canBePrinted map[string]bool
	allowedFunctionList map[string]bool
	preFormatList PreFormatList
	preColorList map[string]string
	colorAttributeList map[string]color.Attribute
}

var Env Environment

func ToInt(str string, base int) int {
	if ui, err := strconv.ParseUint(str, base, 32); err != nil {
		panic(err)
	} else {
		return int(ui)
	}
}

func (env *Environment) Init(logFileName string, flagDebug bool, flagVerbose bool, states int) error {
	if env.Initialized {
		panic(fmt.Errorf("Environment is already initialized!"))
	}
	logDirName := filepath.ToSlash(filepath.Dir(filepath.ToSlash(logFileName)))
	err := os.MkdirAll(logDirName, os.ModeDir)
	logFile, err := os.Create(logFileName)
	if err != nil {
		panic(err)
	}

	env.LogFile     = logFile
	env.flagDebug   = flagDebug
	env.flagVerbose = flagVerbose
	env.states      = states
	
	env.allowedFunctionList = map[string]bool{
		"fPr"  : true,
		"lFt"  : true,
		"lPn"  : true,
		"lPr"  : true,
	}
	env.SetCanBePrinted()
	env.SetPreFormatList()
	env.SetColorAttributeList()
	return nil
}

func (env *Environment) SetCanBePrinted() {
	for flags := 0; flags < 9; flags++ {
	
		//flags := ToInt(strFlags, env.states)
		flagDebug := int(flags / env.states)
		flagVerbose := int(math.Mod(float64(flags), float64(env.states)))
		outDebug := false
		outVerbose := false

		switch flagDebug {
		case 0:
			outDebug = true
		case 1:
			outDebug = !env.flagDebug
		case 2:
			outDebug = env.flagDebug
		}
		
		switch flagVerbose {
		case 0:
			outVerbose = true
		case 1:
			outVerbose = !env.flagVerbose
		case 2:
			outVerbose = env.flagVerbose
		}

		strFlag  := strconv.Itoa(flagDebug) + strconv.Itoa(flagVerbose)
		canPrint := outDebug && outVerbose
		if env.canBePrinted == nil {
			env.canBePrinted = map[string]bool{
				strFlag : canPrint,
			}
		} else {
			env.canBePrinted[strconv.Itoa(flagDebug) + strconv.Itoa(flagVerbose)] = outDebug && outVerbose
		}
	}
}

func (env *Environment) SetPreFormatList() {
	env.preFormatList = PreFormatList{
		"Sep"       : PreFormat{ "----------", "----------" },
		"BigSep"    : PreFormat{ "--------------------", "--------------------" },
		"Out"       : PreFormat{ " ", ""      },
		"Begin"     : PreFormat{ " ", "--->>" },
		"End"       : PreFormat{ " <<---", "" },
		"BeginLoop" : PreFormat{ " ", "--->>" },
		"EndLoop"   : PreFormat{ " <<---", "" },
		"Start"     : PreFormat{ " ", "--->>" },
		"Finish"    : PreFormat{ " <<---", "" },
		"BeginFunc" : PreFormat{ " ", "--->>" },
		"EndFunc"   : PreFormat{ " <<---", "" },
		"BeginProg" : PreFormat{ " ", "--->>" },
		"EndProg"   : PreFormat{ " <<---", "" },
		"_default"  : PreFormat{ "", ""       },
	}
	env.preColorList = map[string]string {
		"Sep"       : ".FgWhite.Bold", 
		"BigSep"    : ".FgWhite.Bold", 
		"Out"       : ".FgYellow.Bold",
		"Begin"     : ".FgRed.Bold", 
		"End"       : ".FgRed.Bold", 
		"BeginLoop" : ".FgGreen.Bold", 
		"EndLoop"   : ".FgGreen.Bold", 
		"Start"     : ".FgWhite.Bold", 
		"Finish"    : ".FgWhite.Bold", 
		"BeginFunc" : ".FgMagenta.Bold",
		"EndFunc"   : ".FgMagenta.Bold",
		"BeginProg" : ".FgCyan.Bold", 
		"EndProg"   : ".FgCyan.Bold", 
		"_default"  : ".Reset.Bold",
	}
}

func (env *Environment) GetPreColor(selector string) string {
	return env.preColorList[selector]
}

func (env *Environment) SetColorAttributeList() {
	env.colorAttributeList = map[string]color.Attribute{
	    "Reset"        : color.Reset,
		"Bold"         : color.Bold,
		"Faint"        : color.Faint,
		"Italic"       : color.Italic,
		"Underline"    : color.Underline,
		"BlinkSlow"    : color.BlinkSlow,
		"BlinkRapid"   : color.BlinkRapid,
		"ReverseVideo" : color.ReverseVideo,
		"Concealed"    : color.Concealed,
		"CrossedOut"   : color.CrossedOut,
		"FgBlack"      : color.FgBlack,  
		"FgRed"        : color.FgRed, 
		"FgGreen"      : color.FgGreen, 
		"FgYellow"     : color.FgYellow,
		"FgBlue"       : color.FgBlue,  
		"FgMagenta"    : color.FgMagenta,
		"FgCyan"       : color.FgCyan, 
		"FgWhite"      : color.FgWhite, 
		"FgHiBlack"    : color.FgHiBlack,  
		"FgHiRed"      : color.FgHiRed, 
		"FgHiGreen"    : color.FgHiGreen, 
		"FgHiYellow"   : color.FgHiYellow,
		"FgHiBlue"     : color.FgHiBlue,  
		"FgHiMagenta"  : color.FgHiMagenta,
		"FgHiCyan"     : color.FgHiCyan, 
		"FgHiWhite"    : color.FgHiWhite, 
		"BgBlack"	   : color.BgBlack,		
		"BgRed"	       : color.BgRed,			
		"BgGreen"      : color.BgGreen,        
		"BgYellow"     : color.BgYellow,       
		"BgBlue"       : color.BgBlue,        
		"BgMagenta"    : color.BgMagenta,      
		"BgCyan"       : color.BgCyan,         
		"BgWhite"      : color.BgWhite,        
		"BgHiBlack"    : color.BgHiBlack,      
		"BgHiRed"      : color.BgHiRed,        
		"BgHiGreen"    : color.BgHiGreen,      
		"BgHiYellow"   : color.BgHiYellow,     
		"BgHiBlue"     : color.BgHiBlue,       
		"BgHiMagenta"  : color.BgHiMagenta,    
		"BgHiCyan"     : color.BgHiCyan,       
		"BgHiWhite"    : color.BgHiWhite,      
	}
}


func (env *Environment) CanBePrinted(strFlags string) bool {
	strFlags = strings.Trim(strFlags, " ")
	var (result bool; succ bool)
	if result, succ = env.canBePrinted[strFlags]; !succ {
		panic(fmt.Errorf("Flags is not found in env.canPrint."))
	}
	return result
}

func (env *Environment) AllowedFunction(strName string) bool {
	strName = strings.Trim(strName, " ")
	var (result bool; succ bool)
	if result, succ = env.allowedFunctionList[strName]; !succ {
		panic(fmt.Errorf("Function name is not found in env.allowedFunctionList."))
	}
	return result
}

type Selector struct {
	groupName string
	funcType string
	strFlags string
	preFormat string
	format string
	attributeList []color.Attribute
}

func (env *Environment) AnalyzeSelector(selector string) (result Selector, err error) {
	selectArr := strings.SplitN(selector, ".?F:", 2)
	selectArrLen := len(selectArr)

	if selectArrLen > 0 {
		selector = strings.Trim(selectArr[0], " ")
	}
	if selector == "" {
		panic(fmt.Errorf("Selector is empty!"))
	}

	if selectArrLen > 1 {
		result.format = selectArr[1]
	}
	
	selectArr = strings.Split(selector, ".")
	for _, value := range selectArr {
 		if allowedFunction, succ := env.allowedFunctionList[value]; succ && allowedFunction {
			result.groupName = value
		}
 		if value == "S" || value == "L" || value == "R" || value == "B" {
			result.funcType = value
		}
		if _, succ := env.canBePrinted[value]; succ {
			result.strFlags = value
		}
		if _, succ := env.preFormatList[value]; succ {
			result.preFormat = value
		}
		if attribute, succ:= env.colorAttributeList[value]; succ {
			result.attributeList = append(result.attributeList, attribute)
		}
	}
	
	if result.groupName == "" || result.funcType == "" || result.strFlags == "" {
		err = fmt.Errorf("Error! Structure of selector is wrong!")
	}
	return
}

func Out(selector string, args ...interface{}){
	selectArr, err := Env.AnalyzeSelector(selector)
	if err != nil {
		panic(err)
	}
	groupName := selectArr.groupName
	funcType  := selectArr.funcType
	strFlags  := selectArr.strFlags
	format    := selectArr.format
	if (format == "" || format == "%*v") && len(args) > 0 {
		format = strings.Trim(strings.Repeat("%v ", len(args)), " ")
	}
	
	if selectArr.preFormat != "" {
		if preFormat, succ := Env.preFormatList[selectArr.preFormat]; succ {
			first := preFormat.first
			last  := preFormat.last
			format = first + format + last
/*   			if selectArr.funcType == "R" {
				args = append( args, "\n" )
				last = last + "\n"
			}
 */			
  		}
	}
//_ = BREAKPOINT
	
	canBePrinted, succ := Env.canBePrinted[strFlags]
	if !succ || !canBePrinted {
		return
	}
//_ = BREAKPOINT
	if len(selectArr.attributeList) > 0 {
		color.Set(selectArr.attributeList...)
	}
	ExecSelectedFunction(groupName, funcType, format, args...)
	if len(selectArr.attributeList) > 0 {
		color.Unset()
	}

}

func ExecSelectedFunction(groupName string, funcType string, format string, args ...interface{}){
	switch funcType {
	case "L":
		format = "\n" + format
	case "R":
		format = format + "\n"
	case "B":
		format = "\n" + format + "\n"
	}
	switch groupName {
	case "fPr" : 
		fmt.Printf(format, args...)
	case "lFt" :
		log.Fatalf(format, args...)
	case "lPn" :
		log.Panicf(format, args...)
	case "lPr" :
		log.Printf(format, args...)
	}
}