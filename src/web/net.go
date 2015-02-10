package web
import (
    "fmt"
    "net/url")

func ExampleURL(){
    u, err := url.Parse("http://baidu.com")
    if err != nil {
        fmt.Println(err)
    }
    v := url.Values{}
    v.Set("name", "Ava")
    v.Add("friend", "Jess")
    v.Add("friend", "中国")
    v.Add("friend", "Zoe")
    u.RawQuery=v.Encode()
    fmt.Println(u)
}