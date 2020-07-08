package sort

func InsertSort(a []int) {
	l := len(a)

	if l == 1 {
		return
	}

	InsertSort(a[1:])

	for i := 0; i < l-1; i++ {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
		} else {
			break
		}
	}
}
