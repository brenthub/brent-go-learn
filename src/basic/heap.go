package basic
import "fmt"

type HeapTree struct {
    Values []int //存放所有的节点
}

func (tree *HeapTree) Delete() {
    //删除根节点
    datas := tree.Values[1:]
    tree.Build(datas)
}
func (tree *HeapTree) Sort() []int {
    size := len(tree.Values)
    for i := size; i >= 1; i-- {
        tree.Values[i-1],tree.Values[0] = tree.Values[0],tree.Values[i-1]
        tree.Values = tree.maxHeapfify(tree.Values, 1, i-1)
    }
    return tree.Values

}
func (tree *HeapTree) Build(datas []int) *HeapTree {
    var a []int
    for i := len(datas) / 2; i > 0; i-- {
        a = tree.maxHeapfify(datas, i, len(datas))
    }
    tree.Values = a
    return tree
}

func (tree *HeapTree) Add(data int) *HeapTree {
    datas := tree.Values
    datas = append(datas, data)
    size := len(datas)
    tree.Values = tree.heapup(datas, size)
    return tree
}
func (tree *HeapTree) heapup(datas []int, index int) []int {
    if index > 1 {
        parent := index / 2

        parentValue := datas[parent-1]
        indexValue := datas[index-1]
        if parentValue < indexValue {
            tmp := datas[parent-1]
            datas[parent-1] = datas[index-1]
            datas[index-1] = tmp
            tree.heapup(datas, parent)
        } else {
            //没有发生交换，说明新增的数据已经找到它的位置了
            return datas
        }
    }

    return datas
}
func (tree *HeapTree) maxHeapfify(datas []int, i int, heapsize int) []int {
    left := i * 2
    right := i*2 + 1
    maxIndex := i
    if left <= heapsize && datas[left-1] > datas[maxIndex-1] {
        maxIndex = left
    }
    if right <= heapsize && datas[right-1] > datas[maxIndex-1] {
        maxIndex = right
    }
    if maxIndex != i {
        tmp := datas[i-1]
        datas[i-1] = datas[maxIndex-1]
        datas[maxIndex-1] = tmp

        //根与左子节点或者右子节点的值交换了，此时子树可能不符合堆性质
        //因此，要对子树递归同样处理
        tree.maxHeapfify(datas, maxIndex, heapsize)
    }
    return datas
}

func HeapTest(){
    a:=HeapTree{make([]int,0)}
    a.Add(4)
    a.Add(1)
    a.Add(3)
    a.Add(2)
    a.Add(16)
    a.Add(9)
    a.Add(10)
    a.Add(14)
    a.Add(8)
    a.Add(7)

    fmt.Println(a)

    a.Sort()
    fmt.Println(a)
}