package output_test

import (
	"fmt"
	"log"
	"output"
	"runtime"
	"strings"
) 

const BREAKPOINT = "breakpoint"

func CurFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}
	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}
	names := strings.Split(me.Name(), ".")
	return names[len(names)-1]
}

func Ln00(args ...interface{}) {
	curFuncName := CurFuncName()
	args = append([]interface{}{curFuncName+":"}, args...)
	output.Env.Out(curFuncName, "log.Print", "00", "ln", "", args...)
	output.Env.Out(curFuncName, "fmt.Print", "00", "ln", "", args...)
}

func Ln20(args ...interface{}) {
	curFuncName := CurFuncName()
	args = append([]interface{}{curFuncName+":"}, args...)
	output.Env.Out(curFuncName, "log.Print", "20", "ln", "", args...)
	output.Env.Out(curFuncName, "fmt.Print", "20", "ln", "", args...)
}

func Ln02(args ...interface{}) {
	curFuncName := CurFuncName()
	args = append([]interface{}{curFuncName+":"}, args...)
	output.Env.Out(curFuncName, "log.Print", "02", "ln", "", args...)
	output.Env.Out(curFuncName, "fmt.Print", "02", "ln", "", args...)
}

func Ln21(args ...interface{}) {
	curFuncName := CurFuncName()
	args = append([]interface{}{curFuncName+":"}, args...)
	output.Env.Out(curFuncName, "log.Print", "21", "ln", "", args...)
	output.Env.Out(curFuncName, "fmt.Print", "21", "ln", "", args...)
}

func Ln12(args ...interface{}) {
	curFuncName := CurFuncName()
	args = append([]interface{}{curFuncName+":"}, args...)
	output.Env.Out(curFuncName, "log.Print", "12", "ln", "", args...)
	output.Env.Out(curFuncName, "fmt.Print", "12", "ln", "", args...)
}

func Ln22(args ...interface{}) {
	curFuncName := CurFuncName()
	args = append([]interface{}{curFuncName+":"}, args...)
	output.Env.Out(curFuncName, "log.Print", "22", "ln", "", args...)
	output.Env.Out(curFuncName, "fmt.Print", "22", "ln", "", args...)
}

func Mix(args ...interface{}) {
	curFuncName := CurFuncName()
	args = []interface{}{curFuncName+":"}
	output.Env.Out(curFuncName, "fmt.Print", "00", "ln", "", []interface{}{"log.Print.00"}...)
	output.Env.Out(curFuncName, "fmt.Print", "11", "ln", "", []interface{}{"log.Print.11"}...)
	output.Env.Out(curFuncName, "fmt.Print", "12", "ln", "", []interface{}{"log.Print.12"}...)
	output.Env.Out(curFuncName, "fmt.Print", "21", "ln", "", []interface{}{"log.Print.21"}...)
	output.Env.Out(curFuncName, "fmt.Print", "22", "ln", "", []interface{}{"log.Print.22"}...)
}

func ExampleToInt(){
	s := "20"
	fmt.Printf("%s (%d) -> %d\n", s, 3, output.ToInt(s, 3))
	// Output:
	// 20 (3) -> 6
	
}

func Example(){
	fmt.Println(CurFuncName())
	const (
		cDebug    = true
		cVerbose  = true
	)
	output.Env.Init("_tmp/log/_settings.log", cDebug, cVerbose, 3)
	defer output.Env.LogFile.Close()
	log.SetOutput(output.Env.LogFile)
//_ = BREAKPOINT
	Ln00("ĄČĘĖĮŠŲŪŽąčęėįšųūž")
	Ln20("ĄČĘĖĮŠŲŪŽąčęėįšųūž")
	Ln02("ĄČĘĖĮŠŲŪŽąčęėįšųūž")
	Ln21("ĄČĘĖĮŠŲŪŽąčęėįšųūž")
	Ln12("ĄČĘĖĮŠŲŪŽąčęėįšųūž")
	Ln22("ĄČĘĖĮŠŲŪŽąčęėįšųūž")
	// Output: 
	// Example
	// Ln00: ĄČĘĖĮŠŲŪŽąčęėįšųūž
	// Ln20: ĄČĘĖĮŠŲŪŽąčęėįšųūž
	// Ln02: ĄČĘĖĮŠŲŪŽąčęėįšųūž
	// Ln22: ĄČĘĖĮŠŲŪŽąčęėįšųūž
}

func ExampleX(){
	fmt.Println(CurFuncName())
	const (
		cDebug    = false
		cVerbose  = false
	)
	output.Env.Init("_tmp/log/_settings.log", cDebug, cVerbose, 3)
	defer output.Env.LogFile.Close()
	log.SetOutput(output.Env.LogFile)
_ = BREAKPOINT
	Mix()
	// Output: 
	// ExampleX
	// log.Print.21
}