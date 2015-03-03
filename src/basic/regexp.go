package basic
import (
    "fmt"
    "regexp"
    "time"
)

func RegTest(){
    matched, err := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+$", "dddd163.com")
    fmt.Println(matched, err)

    time := time.Date(2015,time.May,2,10,34,23,200,time.UTC)

    fmt.Println(time)
}