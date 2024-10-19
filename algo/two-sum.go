package main

import "fmt"

func twoSum(nums []int, target int) []int {
    // Create a map to store the indices of the numbers
    numMap := make(map[int]int)
    
    // Iterate over the array
    for i, num := range nums {
        // Calculate the complement that would sum to the target
        complement := target - num
        
        // Check if the complement exists in the map
        if idx, ok := numMap[complement]; ok {
            return []int{idx, i} // Return the indices
        }
        
        // Store the number and its index in the map
        numMap[num] = i
    }
    
    // Return an empty slice if no solution is found
    return []int{}
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result) // Output: [0, 1]
}
