package main

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Find the distance between two numbers
func distance(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	// Read the data
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Load the heaps
	xs := &IntHeap{}
	ys := &IntHeap{}
	heap.Init(xs)
	heap.Init(ys)

	for _, record := range records {
		x, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		heap.Push(xs, x)

		y, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatal(err)
		}
		heap.Push(ys, y)
	}

	// n => [A.count(n), B.count(n)]
	counts := map[int][]int{}

	// Process the data
	total_distance := 0
	for xs.Len() > 0 && ys.Len() > 0 {
		x := heap.Pop(xs).(int)
		y := heap.Pop(ys).(int)

		_, ok := counts[x]
		if !ok {
			counts[x] = []int{0, 0}
		}

		_, ok = counts[y]
		if !ok {
			counts[y] = []int{0, 0}
		}

		counts[x][0]++
		counts[y][1]++

		total_distance = total_distance + distance(x, y)
	}

	total_similarity_score := 0
	for n, n_counts := range counts {
		total_similarity_score += n * n_counts[0] * n_counts[1]
	}

	fmt.Printf("total distance: %d\ntotal similarity score: %d\n",
		total_distance,
		total_similarity_score,
	)
}
