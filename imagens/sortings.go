package main

func merge(firstHalf []int, secondHalf []int) []int {
	var result []int = []int{}
	var firstHalfPointer = 0
	var secondHalfPointer = 0

	for firstHalfPointer < len(firstHalf) && secondHalfPointer < len(secondHalf) {
		if firstHalf[firstHalfPointer] <= secondHalf[secondHalfPointer] {
			result = append(result, firstHalf[firstHalfPointer])
			firstHalfPointer++
		} else {
			result = append(result, secondHalf[secondHalfPointer])
			secondHalfPointer++
		}
	}

	for firstHalfPointer < len(firstHalf) {
		result = append(result, firstHalf[firstHalfPointer])
		firstHalfPointer++
	}

	for secondHalfPointer < len(secondHalf) {
		result = append(result, secondHalf[secondHalfPointer])
		secondHalfPointer++
	}

	return result
}

func MergeSort(list []int) []int {

	if len(list) <= 1 {
		return list
	}

	var firstHalf = list[:len(list)/2]
	var secondHalf = list[len(list)/2:]

	firstHalf = MergeSort(firstHalf)
	secondHalf = MergeSort(secondHalf)

	return merge(firstHalf, secondHalf)
}

func QuickSort(list []int, start int, end int) []int {

	var pivot = list[(start+end)/2]
	var startPointer = start
	var endPointer = end
	var aux int

	for startPointer <= endPointer {
		for list[startPointer] < pivot {
			startPointer++
		}
		for list[endPointer] > pivot {
			endPointer--
		}

		if startPointer <= endPointer {
			aux = list[startPointer]
			list[startPointer] = list[endPointer]
			list[endPointer] = aux

			startPointer++
			endPointer--
		}
	}

	if start < endPointer {
		list = QuickSort(list, start, endPointer)
	}

	if startPointer < end {
		list = QuickSort(list, startPointer, end)
	}

	return list
}
