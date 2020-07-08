package sort

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	i, j := 0, len(arr)-1
	target := arr[0]

	for i < j {
		for arr[j] >= target && j > i {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]

		for arr[i] <= target && i < j {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	QuickSort(arr[:i])
	QuickSort(arr[i+1:])
}
