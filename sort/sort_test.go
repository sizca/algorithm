package sort

import (
	"fmt"
	"testing"
)

var testItems = []int{29, 10, 14, 37, 14}
var testItems2 = []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 38}

func TestInsertionSort(t *testing.T) {
	fmt.Println(InsertionSort([]int{}))
	fmt.Println(InsertionSort([]int{1}))
	fmt.Println(testItems)
	fmt.Println(InsertionSort(testItems))
}

func TestInsertionSortInPlace(t *testing.T) {
	fmt.Println(testItems)
	InsertionSortInPlace(testItems)
}

func TestMergeSort(t *testing.T) {
	fmt.Println(testItems2)
	fmt.Println(MergeSort(testItems2))
}
