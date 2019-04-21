package sort

import "fmt"

// MergeSort : [] = []
//             as = merge(MergeSort(xs), MergeSort(ys))
//                  where xs is left half of as
//                        ys is right half of as
// merge [], xs = xs
//       xs, [] = xs
//       [x:xs], [y:ys] = if x <= y then [x:merge(xs, [y:ys])]
//                                  else [y:merge([x:xs], ys)]
func MergeSort(items []int) []int {
	size := len(items)
	if size <= 1 {
		return items[:]
	}
	mid := size / 2
	return merge(MergeSort(items[:mid]), MergeSort(items[mid:]))
}

func merge(left []int, right []int) (sorted []int) {
	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}

	x, xs := left[0], left[1:]
	y, ys := right[0], right[1:]

	if x <= y {
		return append([]int{x}, merge(xs, right)...)
	}
	return append([]int{y}, merge(left, ys)...)
}

// MergeSortInPlace sort in place
func MergeSortInPlace(a []int) {
	size := len(a)
	for i := 1; i < size; i = i * 2 {
		for j := 0; j <= size; j = j + i*2 {
			mergeInPlace(a, j, ((i+j)*2-1)/2, j+(i*2)-1)
		}
	}
}

func mergeInPlace(a []int, l, m, r int) {
	// fmt.Printf("merge(%d, %d, %d)\n", l, m, r)
	i := l
	j := m + 1
	size := len(a)
	if m >= size {
		return
	}
	if r >= size {
		r = size - 1
	}

	for i <= m && j <= r {
		if a[i] <= a[j] {
			i++
			continue
		}

		for s := j; s > i; s-- {
			a[s-1], a[s] = a[s], a[s-1]
			fmt.Println(a)
		}
		i++
		j++
		m++
	}
}
