package sort

import "fmt"

// QuickSort : [] = []
//             [x] = [x]
//             [x:xs] = [QuickSort(smaller):x:QuickSort(bigger)]
//                      where smaller = [a | a <- xs, a <= x]
//                            bigger  = [b | b <- xs, b > x]
func QuickSort(items []int) []int {
	size := len(items)
	if size <= 1 {
		return items[:]
	}

	x, xs := items[0], items[1:]
	// smaller := []int{}
	// bigger := []int{}

	// for _, a := range xs {
	// 	if a <= x {
	// 		smaller = append(smaller, a)
	// 	} else {
	// 		bigger = append(bigger, a)
	// 	}
	// }

	smaller := filter(xs, small, x)
	bigger := filter(xs, big, x)

	sorted := append(QuickSort(smaller), x)
	sorted = append(sorted, QuickSort(bigger)...)

	return sorted
}

func small(x, a int) bool {
	if x <= a {
		return true
	}
	return false
}

func big(x, a int) bool {
	if x > a {
		return true
	}
	return false
}

type condFunc func(int, int) bool

func filter(items []int, cond condFunc, a int) []int {
	if len(items) == 0 {
		return []int{}
	}

	x, xs := items[0], items[1:]

	if cond(x, a) {
		return append([]int{x}, filter(xs, cond, a)...)
	}

	return filter(xs, cond, a)
}

// QuickSortInPlace sort in place
func QuickSortInPlace(items []int) {
	quickSortInPlace(items, 0, len(items)-1)
}

func quickSortInPlace(items []int, left, right int) {
	size := right - left + 1
	if size <= 1 {
		return
	}
	fmt.Printf("quicksort(items, %d, %d)\n", left, right)

	pivot := divide(items, left, right)
	quickSortInPlace(items, left, pivot-1)
	quickSortInPlace(items, pivot+1, right)
}

func divide(items []int, left, right int) int {
	fmt.Printf("divide(items, %d, %d)\n", left, right)
	p := left
	pivot := items[left]
	low := left + 1
	high := right

	for low <= high {
		for low <= right && items[low] <= pivot {
			low++
		}

		for left+1 <= high && pivot <= items[high] {
			high--
		}

		if low < high {
			items[low], items[high] = items[high], items[low]
			fmt.Println(items)
		}
	}

	if left < high {
		items[left], items[high] = items[high], items[left]
		fmt.Println(items)
		p = high
	}

	fmt.Println("pivot:", p)
	return p
}
