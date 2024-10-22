package main

import (
	"fmt"
	"sort"
)

// Function to find 3-sum numbers that match the target
func threeSum(nums []int, target int) [][]int {
	sort.Ints(nums) // Sort the array to simplify finding triplets
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Avoid duplicates for the first element of the triplet
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				// Found a triplet
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Avoid duplicates for the second element of the triplet
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Avoid duplicates for the third element of the triplet
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	// Example array and target sum
	nums := []int{-1, 0, 1, 2, -1, -4}
	target := 0

	// Call the threeSum function
	result := threeSum(nums, target)

	// Print the result
	fmt.Println("Triplets that sum to target:", result)
}
