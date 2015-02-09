package sorter
import "flag"

var infile *string =flag.String("s","infile","File contains values for sorting")
var outfile *string = flag.String("o","outfile","file to receive sorted values")
var algo *string =flag.String("a","qsort","sort algorithm")

func Sorter(){
    flag.Parse()

}