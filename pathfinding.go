package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type Node struct {
	Point
	Parent   *Node
	Cost     float64
	Priority float64
	Index    int
}

// PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// The priority is based on the A* heuristic (f = g + h)
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := x.(*Node)
	n.Index = len(*pq)
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(node *Node, cost, priority float64) {
	node.Cost = cost
	node.Priority = priority
	heap.Fix(pq, node.Index)
}

// ManhattanDistance is a common heuristic for grid-based movement.
func ManhattanDistance(a, b Point) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

// Neighbors returns the 4 neighbors (N, S, E, W) of a given node in the grid.
func Neighbors(p Point, grid [][]int) []Point {
	neighbors := []Point{}
	dirs := []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // Up, Down, Left, Right

	for _, d := range dirs {
		neighbor := Point{p.X + d.X, p.Y + d.Y}
		if neighbor.X >= 0 && neighbor.X < len(grid) && neighbor.Y >= 0 && neighbor.Y < len(grid[0]) {
			if grid[neighbor.X][neighbor.Y] == 0 { // assuming 0 means walkable
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}

func AStar(start, goal Point, grid [][]int) []*Node {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	startNode := &Node{
		Point:    start,
		Cost:     0,
		Priority: ManhattanDistance(start, goal),
	}
	heap.Push(&pq, startNode)

	visited := make(map[Point]bool)
	cameFrom := make(map[Point]*Node)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)

		if current.Point == goal {
			return reconstructPath(current)
		}

		visited[current.Point] = true

		for _, neighbor := range Neighbors(current.Point, grid) {
			if visited[neighbor] {
				continue
			}

			newCost := current.Cost + 1 // assuming uniform cost grid
			neighborNode := &Node{
				Point: neighbor,
				Cost:  newCost,
				Priority: newCost + ManhattanDistance(neighbor, goal),
				Parent: current,
			}

			heap.Push(&pq, neighborNode)
			cameFrom[neighbor] = neighborNode
		}
	}
	return nil
}

// reconstructPath reconstructs the path from the goal node to the start node.
func reconstructPath(node *Node) []*Node {
	var path []*Node
	for node != nil {
		path = append([]*Node{node}, path...)
		node = node.Parent
	}
	return path
}

func main() {
	// Define a simple grid: 0 is walkable, 1 is blocked
	grid := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}

	start := Point{0, 0}
	goal := Point{4, 4}

	path := AStar(start, goal, grid)

	if path != nil {
		fmt.Println("Path found:")
		for _, node := range path {
			fmt.Printf("(%d, %d) -> ", node.X, node.Y)
		}
		fmt.Println("Goal")
	} else {
		fmt.Println("No path found.")
	}
}
