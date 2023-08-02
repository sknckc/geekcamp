package main

import (
	"fmt"
	"geekcamp/week1/slice"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := []string{"i", "am", "xiaoming", "and", "i", "love", "go"}
	fmt.Println(slice.SliceDel(arr, 3))
	fmt.Println(slice.SliceDel2(arr, 3))
	fmt.Println(slice.SliceDel3(arr2, 3))
}
