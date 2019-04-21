package sort

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
