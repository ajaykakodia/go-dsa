package sorting

type QuickSort struct {
	arr []int
}

func (q *QuickSort) partition(li, ri int) int {
	pivotIndex := ri

	pivot := q.arr[ri]

	ri = ri - 1

	for {

		for q.arr[li] < pivot {
			li += 1
		}

		for ri > -1 && q.arr[ri] > pivot {
			ri -= 1
		}

		if li >= ri {
			break
		}

		q.arr[li], q.arr[ri] = q.arr[ri], q.arr[li]
		li += 1
	}

	q.arr[li], q.arr[pivotIndex] = q.arr[pivotIndex], q.arr[li]
	return li
}

func (q *QuickSort) SortArray() {
	q.sort(0, len(q.arr)-1)
}

func (q *QuickSort) SetArray(arr []int) {
	q.arr = arr
}

func (q *QuickSort) SortedArray() []int {
	return q.arr
}

func (q *QuickSort) sort(li, ri int) {
	if ri-li <= 0 {
		return
	}

	pivotIndex := q.partition(li, ri)

	q.sort(li, pivotIndex-1)
	q.sort(pivotIndex+1, ri)
}

func (q *QuickSort) QuickSelect(kthNum int) int {
	return q.selectElement(kthNum-1, 0, len(q.arr)-1)
}

// func QuickSelect(arr []int, kthNumber int) int {
// 	arrQuick = arr
// 	return quickSelect(kthNumber, 0, len(arr)-1)
// }

func (q *QuickSort) selectElement(kthNumber, li, ri int) int {

	if ri-li <= 0 {
		return q.arr[li]
	}

	pi := q.partition(li, ri)
	if kthNumber < pi {
		return q.selectElement(kthNumber, li, pi-1)
	} else if kthNumber > pi {
		return q.selectElement(kthNumber, pi+1, ri)
	}
	return q.arr[pi]
}
