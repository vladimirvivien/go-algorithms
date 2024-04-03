package sorting

// BubbSort
// It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.
// The pass through the list is repeated until the list is sorted.
//
// Performance:
// Time Complexity: O(n^2) in the worst and average case,
// and O(n) in the best case (when the list is already sorted).
func BubbSort(data []int) {
	for i := range data {
		for j := i+1; j < len(data); j++{
			if data[j] < data[i] {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}
}
