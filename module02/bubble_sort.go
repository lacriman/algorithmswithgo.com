package module02

import "sort"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
// ![A list containing the numbers 5, 4, 2, 3, 1, 0](/img/bubble_sort_01.png)

func BubbleSortInt(list []int) {
	swapped := true
	for swapped {
		swapped = false
		for i, j := 0, 1; j < len(list); i, j = i+1, j+1 {
			if list[i] > list[j] {
				swapped = true
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
}
