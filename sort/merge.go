package sort

func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	m := n / 2
	return merge(MergeSort(arr[m:]), MergeSort(arr[:m]))
}

func merge(arr1 []int, arr2 []int) []int {
	l1, l2 := len(arr1), len(arr2)
	l := l1 + l2
	arr := make([]int, l)

	m, n := 0, 0
	for i := 0; i < l; i++ {
		if m < l1 && n < l2 {
			if arr1[m] < arr2[n] {
				arr[i] = arr1[m]
				m++
			} else {
				arr[i] = arr2[n]
				n++
			}
		} else if m < l1 {
			arr[i] = arr1[m]
			m++
		} else if n < l2 {
			arr[i] = arr2[n]
			n++
		}
	}

	return arr
}
