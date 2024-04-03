package sorting

// InsertionSort
// Builds the final sorted array one item at a time by repeatedly inserting the next element into the sorted portion of the array.
// Time Complexity: O(n^2) in the worst and average case, and O(n) in the best case (when the list is nearly sorted).
func InsertionSort(data []int) {
	// given array [0..n]
	for i := 1; i < len(data); i++ { // walk unsorted subarray I from [1..n]
		for j := i; j > 0; j-- { // walk sorted subarray J in reverse from [i..0]
			if data[j-1] >= data[j] {
				temp := data[j-1]
				data[j-1] = data[j]
				data[j] = temp
			}
		}
	}
}
