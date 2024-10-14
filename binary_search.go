package main

import (
    "fmt"
)

// BinarySearch performs binary search on a sorted slice.
// It returns the index of the target if found, otherwise -1.
func BinarySearch(arr []int, target int) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := low + (high-low)/2 // To avoid potential overflow

        if arr[mid] == target {
            return mid // Target found, return the index
        } else if arr[mid] < target {
            low = mid + 1 // Search in the right half
        } else {
            high = mid - 1 // Search in the left half
        }
    }

    return -1 // Target not found
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
    target := 7

    result := BinarySearch(arr, target)
    
    if result != -1 {
        fmt.Printf("Element %d found at index %d\n", target, result)
    } else {
        fmt.Printf("Element %d not found in the array\n", target)
    }
}
