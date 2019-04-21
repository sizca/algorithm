package sort

import (
	"fmt"
	"testing"
)

var testItems = []int{29, 10, 14, 37, 14}
var testItems2 = []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 38}

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

func TestMergeSortInPlace(t *testing.T) {
	fmt.Println(len(testItems2))
	fmt.Println(testItems2)
	MergeSortInPlace(testItems2)
}

func TestQuickSort(t *testing.T) {
	fmt.Println(testItems2)
	fmt.Println(QuickSort(testItems2))
}

func TestQuickSortInPlace(t *testing.T) {
	fmt.Println(testItems2)
	QuickSortInPlace(testItems2)
}
