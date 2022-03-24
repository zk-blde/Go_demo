package main

import (
	"fmt"
	"strconv"
)

type Price int

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

type Price2 string
func (i Price2) String() string {
	return string(i)
}

type Stringer interface {
	String() string
	~int | ~string   // 这个必须挑明Price和Price2底层类型必须是int或者string
}

func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}


// 自带的comparable约束类型, 可以比较的类型
func findFunc[T comparable](a []T, v T) int {
	for i, e := range a{
		if e == v {
			return i
		}
	}
	return -1
}
// comparable 的约束类型支持整数和字符, 自定义结构体(值类型才可以, 内嵌切片或者inteface的不可以)也可以嵌套在自定义约束中(待验证), 例:
// type ShowPrice interface {
// 	int | string | comparable
// }
func main() {
	a := []Price{1, 2, 3, 4, 5}
	b := []Price2{"a", "b", "c", "d", "e"}
	// Stringify(a)
	fmt.Println(Stringify(a))
	fmt.Println(Stringify(b))

	fmt.Println(findFunc(a, 5))
	fmt.Println(findFunc(b, "e"))
}