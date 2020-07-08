package sort

func HeapSort(a []int) {
	l := len(a)

	if l <= 1 {
		return
	}

	buildHeap(a)
	HeapSort(a[1:])
}

func buildHeap(a []int) {
	l := len(a)

	if l <= 1 {
		return
	}

	if l > 2 {
		for i := (l - 1) / 2; i > 0; i-- {
			if a[2*i] < a[i] {
				a[2*i], a[i] = a[i], a[2*i]
			}
			if 2*i+1 < l && a[2*i+1] < a[i] {
				a[2*i+1], a[i] = a[i], a[2*i+1]
			}
		}
	}

	if a[0] > a[1] {
		a[0], a[1] = a[1], a[0]
	}
}
