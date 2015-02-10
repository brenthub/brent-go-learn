package basic

import (
	"fmt"
	"os"
)

var V1 int
var V2 string
var V3 [10]int //数组
var V4 []int   //切片
var V5 struct {
	F int
}
var V6 *int           //指针
var V7 map[string]int //Map
var V8 func(a int) int

var (
	V9  int
	V10 string
)

var V11 int = 10
var V12 = 10

const Pi float64 = 3.1415926

const zero = 0.0

const u, v = 0, "foo"

const ( //iota在出现const的时候会被重设为0
	c0 = iota //0
	c1 = iota //1
	c2 = iota //2
)

const ( //枚举
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func Test() {
	v13 := 10
	v14 := 5 % 3
	fmt.Println("running", v13, v14)

	var ini int32
	var inj int64
	ini, inj = 1, 2
	/*if i==j{//error invalid operation: i == j (mismatched types int32 and int64)
	    fmt.Print("i,j")
	}*/
	if ini == 1 || inj == 2 {
		fmt.Println("i,j")
	}
	str := "Hello world!"
	fmt.Printf("The length of %s is %d \n", str, len(str))

	for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Print(ch, " ")
	}
	fmt.Println()

	for _, ch := range str {
		fmt.Print(ch, " ")
	}
	fmt.Println()

	var array [3]int = [3]int{1, 2, 3}
	var slice []int = array[:2]

	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	slice1 := make([]int, 5, 7) //类型,长度，预留空间(当空间不够时，增长为它的倍数)
	for _, v := range slice1 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("cap", cap(slice1))
	fmt.Println("len", len(slice1))

	slice1 = append(slice1, []int{4, 8, 5}...) //三点表示打散数组为元素

	fmt.Println("cap", cap(slice1))
	fmt.Println("len", len(slice1))

	copy(slice, slice1) //从slice1中复制到slice(长度不一样没关系)

	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

type Person struct {
	ID      string
	Name    string
	Address string
}

func MapTest() {
	personDB := make(map[string]Person)
	personDB["111"] = Person{"111", "Tom", "Room 204"}
	personDB["222"] = Person{"222", "Kobe", ""}

	person, ok := personDB["111"] //map查找
	if ok {
		fmt.Println(person)
	}
	delete(personDB, "111")

	person, ok = personDB["111"]

	fmt.Println(person, ok)
}

func Add(args ...int) int { //不定参数，必须是最后一个
	sum := 0
	for _, i := range args {
		sum += i
	}
	return sum
}

type PathError struct { //自定义异常
	Op   string
	Path string
	Err  error
}

/**
* 构造函数
*/
func NewPathError(op,path string,err error) *PathError{
    return &PathError{op,path,err}
}

func (e *PathError) Error() string { //实现error接口的Error方法
	return e.Op + " " + e.Path + ":" + e.Err.Error()
}

func OpenFile(name string) (file *os.File, err error) {
	file, e := os.Open(name)

	defer file.Close()

	if e != nil {
        //创建对象new(PathError)
        //或者&PathError{}
        //或者&PathError{Op:"opfile"}
        //NewPathError
        //&PathError{"opfile", name, e}
		return nil, NewPathError("opfile", name, e)
	} else {
		return file, nil
	}

}
