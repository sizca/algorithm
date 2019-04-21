package sort

import "fmt"

// InsertionSort : [] = []
//                 [x] = [x]
//                 [x:xs] = insert(InsertionSort(xs), x)
// insert : [], i = [i]
//          [x:xs], i = if x > i then [i:x:xs]
//                               else [x:insert(xs,i)]
func InsertionSort(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}

	x, xs := a[0], a[1:]

	return insert(InsertionSort(xs), x)
}

func insert(a []int, i int) (sorted []int) {
	if len(a) == 0 {
		return []int{i}
	}

	x, xs := a[0], a[1:]
	if x > i {
		return append([]int{i}, a...)
	}

	return append([]int{x}, insert(xs, i)...)
}

// InsertionSortInPlace sort in place
func InsertionSortInPlace(a []int) {
	size := len(a)
	if size <= 1 {
		return
	}

	for i := 1; i < size; i++ {
		for j := i - 1; j >= 0; j-- {
			if a[j] <= a[j+1] {
				break
			}
			a[j], a[j+1] = a[j+1], a[j]
			fmt.Println(a)
		}
	}
}
