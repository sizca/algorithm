package sort

import "fmt"

// InsertionSort : [] = []
//                 [x] = [x]
//                 [x:xs] = insert(InsertionSort(xs), x)
// insert : [], i = [i]
//          [x:xs], i = if x > i then [i:x:xs]
//                               else [x:insert(xs,i)]
func InsertionSort(items []int) []int {
	if len(items) == 0 {
		return items[:]
	}

	x, xs := items[0], items[1:]

	return insert(InsertionSort(xs), x)
}

func insert(items []int, i int) (sorted []int) {
	if len(items) == 0 {
		return []int{i}
	}

	x, xs := items[0], items[1:]
	if x > i {
		return append([]int{i}, items...)
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
