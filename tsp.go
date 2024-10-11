package main

import (
	"fmt"
	"math"
)

// Function to calculate the total distance of a given path
func calculateDistance(path []int, dist [][]int) int {
	n := len(path)
	totalDist := 0
	for i := 0; i < n-1; i++ {
		totalDist += dist[path[i]][path[i+1]]
	}
	totalDist += dist[path[n-1]][path[0]] // Return to the start city
	return totalDist
}

// Function to generate all permutations of a slice
func generatePermutations(arr []int, l int, r int, permutations *[][]int) {
	if l == r {
		perm := make([]int, len(arr))
		copy(perm, arr)
		*permutations = append(*permutations, perm)
	} else {
		for i := l; i <= r; i++ {
			arr[l], arr[i] = arr[i], arr[l] // Swap
			generatePermutations(arr, l+1, r, permutations)
			arr[l], arr[i] = arr[i], arr[l] // Backtrack
		}
	}
}

// Brute-force solution to find the shortest path for the TSP
func tsp(dist [][]int) (int, []int) {
	n := len(dist)
	cities := make([]int, n)
	for i := 0; i < n; i++ {
		cities[i] = i
	}

	permutations := [][]int{}
	generatePermutations(cities, 0, n-1, &permutations)

	minDistance := math.MaxInt
	var bestPath []int

	// Check every permutation and calculate the total distance
	for _, path := range permutations {
		distance := calculateDistance(path, dist)
		if distance < minDistance {
			minDistance = distance
			bestPath = path
		}
	}

	return minDistance, bestPath
}

func main() {
	// Example distance matrix (symmetric)
	distances := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}

	minDistance, bestPath := tsp(distances)

	fmt.Printf("The shortest path has distance %d and follows this route: %v\n", minDistance, bestPath)
}
