package basic
import (
    "fmt"
    "regexp")

func RegTest(){
    matched, err := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+$", "dddd163.com")
    fmt.Println(matched, err)
}