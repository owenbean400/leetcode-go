package main

func twoSum(nums []int, target int) []int {
	//Initalize starting values and sort array
	size := len(nums)
	sortedNum := mergeSort(nums) // This uses up most of the memory

	start, finish := size/2, size/2+1

	// Base case
	if size == 2 {
		return []int{0, 1}
	}

	// Move pointers to where the highest number before target between two numbers next to each other O(n)
	if sortedNum[start]+sortedNum[finish] > target {
		for sortedNum[start]+sortedNum[finish] > target && start >= 0 {
			start--
			finish--
		}
	} else if sortedNum[start]+sortedNum[finish] < target {
		for sortedNum[start]+sortedNum[finish] < target && finish < size {
			start++
			finish++
		}
		if sortedNum[start]+sortedNum[finish] > target && start >= 0 {
			start--
			finish--
		}
	}

	// Will now find the target, which is O(n^2), however, because the pointer are set to optimal point, the algorithm.
	for start >= 0 {
		savedFinish := finish
		for sortedNum[start]+sortedNum[finish] < target && finish < size-1 {
			finish += 1
		}
		// Target found, find index based on sorted array to none sorted
		if sortedNum[start]+sortedNum[finish] == target {
			first, second := 0, 0
			for sortedNum[start] != nums[first] && first < len(nums) {
				first += 1
			}
			for (sortedNum[finish] != nums[second] && second < len(nums)) || second == first {
				second += 1
			}
			if first < second {
				return []int{first, second}
			} else {
				return []int{second, first}
			}
		}
		finish = savedFinish
		start -= 1
	}

	return []int{0, 1}
}

// Merge sort O(n log(n))
func mergeSort(arr []int) []int {
	num := len(arr)

	if num == 1 {
		return arr
	}

	m := int(num / 2)

	left, right := make([]int, m), make([]int, num-m)
	for i := 0; i < num; i++ {
		if i < m {
			left[i] = arr[i]
		} else {
			right[i-m] = arr[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
