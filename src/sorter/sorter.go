package sorter
import (
    "flag"
    "os"
    "fmt"
    "bufio"
    "io"
    "strconv"
    "time"
    "sorter/algo")

var infile *string = flag.String("s", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "file to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

func Sorter() {
    flag.Parse()

    values, err := readValues(*infile)

    if err ==nil {
        t1 := time.Now()
        switch *algorithm{
            case "qsort":
                algo.QuickSort(values)
            case "bubble":
                algo.BubboleSort(values)
            default:
                fmt.Println("Unkown sorting algorithm", *algorithm)
        }
        t2 := time.Now()
        fmt.Println("costs ", t2.Sub(t1))
        writeValues(values, *outfile)
    }else {
        fmt.Println(err)
    }


}

func readValues(name string) (values []int, err error) {
    file, err := os.Open(name)
    if err!=nil {
        fmt.Println("Failed to open the input file ", name)
    }

    defer file.Close()

    br := bufio.NewReader(file)
    values=make([]int, 0)
    for {
        line, isPrefix, err1 := br.ReadLine()
        if err1!=nil {
            if err1!=io.EOF {
                err=err1
            }
            break
        }
        if isPrefix {
            fmt.Println("A too long line")
            return
        }

        str := string(line)

        value, err1 := strconv.Atoi(str)

        if err1==nil {
            err=err1
            return
        }

        values=append(values, value)
    }
    return
}

func writeValues(values []int, outfile string) error {
    file, err := os.Create(outfile)
    if err!=nil {
        fmt.Println("Failed to ceate the output file ", outfile)
        return err
    }
    defer file.Close()
    for _, value := range values {
        str := strconv.Itoa(value) //把一个整数转成字符串
        file.WriteString(str+"\n")
    }
    return nil
}