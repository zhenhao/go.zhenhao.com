package sort

func QuickSort(arr []int) {
	i, j := 0, len(arr)-1

	if (i >= j) {
		return
	}

	p := arr[0]

	for i < j {
		for arr[j] >= p && j > i {
			j--
		}

		arr[i], arr[j] = arr[j], arr[i]

		for arr[i] <= p && i < j {
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	QuickSort(arr[:i])
	QuickSort(arr[i+1:])
}
