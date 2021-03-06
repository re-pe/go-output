package output

import (
  "fmt"
  "github.com/fatih/color"
  "log"
  "os"
  "path/filepath"
  "regexp"
  "strings"
)

type Flags struct {
  Debug bool
  Verbose bool
}

type Formated struct{
  Text string
  Attributes []color.Attribute
}

var flags *Flags
var attributeList map[string]color.Attribute

func OuputInit(flagsImp *Flags){
  flags = flagsImp
  attributeList = map[string]color.Attribute{
    "R"    : color.Reset,
    "B"    : color.Bold,
    "F"    : color.Faint,
    "I"    : color.Italic,
    "U"    : color.Underline,
    "BS"   : color.BlinkSlow,
    "BP"   : color.BlinkRapid,
    "RV"   : color.ReverseVideo,
    "C"    : color.Concealed,
    "CO"   : color.CrossedOut,
    "FK"   : color.FgBlack,  
    "FR"   : color.FgRed, 
    "FG"   : color.FgGreen, 
    "FY"   : color.FgYellow,
    "FB"   : color.FgBlue,  
    "FM"   : color.FgMagenta,
    "FC"   : color.FgCyan, 
    "FW"   : color.FgWhite, 
    "FHK"  : color.FgHiBlack,  
    "FHR"  : color.FgHiRed, 
    "FHG"  : color.FgHiGreen, 
    "FHY"  : color.FgHiYellow,
    "FHB"  : color.FgHiBlue,  
    "FHM"  : color.FgHiMagenta,
    "FHC"  : color.FgHiCyan, 
    "FHW"  : color.FgHiWhite, 
    "BK"   : color.BgBlack,    
    "BR"   : color.BgRed,      
    "BG"   : color.BgGreen,        
    "BY"   : color.BgYellow,       
    "BB"   : color.BgBlue,        
    "BM"   : color.BgMagenta,      
    "BC"   : color.BgCyan,         
    "BW"   : color.BgWhite,        
    "BHK"  : color.BgHiBlack,      
    "BHR"  : color.BgHiRed,        
    "BHG"  : color.BgHiGreen,      
    "BHY"  : color.BgHiYellow,     
    "BHB"  : color.BgHiBlue,       
    "BHM"  : color.BgHiMagenta,    
    "BHC"  : color.BgHiCyan,       
    "BHW"  : color.BgHiWhite,      
  }
}

func NewLogFile(fDir, fRunFile, fExt string) (logFile *os.File, err error) {
  err = os.MkdirAll(fDir, os.ModeDir | 0777)
  logFileName := filepath.Base(fRunFile)
  logFileName  = strings.Replace(logFileName, filepath.Ext(logFileName), "", -1)
  logFileName  = fDir + logFileName + fExt
  logFile, err = os.Create(logFileName)
  if err != nil { panic(err) }
  return
}

func checkFormat(args []interface{}) (format string, hasFormat bool, attributes []color.Attribute) {
  if len(args) < 1 { return }
  switch arg := args[0].(type){
  case string: 
    format = arg
    hasFormat = true
  }
  if !hasFormat { return }
  
  if len(format) < 1 || strings.Index(format, "?:") < 0 {
    hasFormat = false
    return
  }

  re := regexp.MustCompile("^(?:([^.?:]+)[.]?)*[?]:")
  attributesStr := re.FindString(format)
  format = re.ReplaceAllString(format, "")

  if len(attributesStr) < 2 { return }
  attributesArr := strings.Split(attributesStr[:len(attributesStr)-2], ".")

  if len(attributesArr) < 1 { return }
  for _, attr := range attributesArr {
    attributes = append(attributes, attributeList[attr])
  }
  return
}

func Format(args ...interface{}) (result Formated) {
  var (format string; hasFormat bool)
  format, hasFormat, result.Attributes = checkFormat(args)
  if hasFormat {
    result.Text = fmt.Sprintf(format, args[1:]...)
  } else {
    result.Text = fmt.Sprint(args...)
  }
  return
}

func Print(args ...interface{}) {
  formated := Format(args...)
  if len(formated.Attributes) > 0 {
    color.Set(formated.Attributes...)
  }
  fmt.Print(formated.Text)
  if len(formated.Attributes) > 0 {
    color.Unset()
  }
}

func Log(args ...interface{}) {
  formated := Format(args...)
  log.Print(formated.Text)
}

func Out(args ...interface{}) {
  formated := Format(args...)
  log.Print(formated.Text)
  if len(formated.Attributes) > 0 {
    color.Set(formated.Attributes...)
  }
  fmt.Print(formated.Text)
  if len(formated.Attributes) > 0 {
    color.Unset()
  }
}

func Debug(args ...interface{}) {
  if flags.Debug {
    Out(args...)
  }
}