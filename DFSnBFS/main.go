package main

import "fmt"

// DFS
func DFS(graph map[string][]string, start string, visited map[string]bool) {
	visited[start] = true
	fmt.Print(start, " ")
	for _, v := range graph[start] {
		if !visited[v] {
			DFS(graph, v, visited)
		}
	}
}

// BFS
func BFS(graph map[string][]string, start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	visited[start] = true
	for len(queue) != 0 {
		s := queue[0]
		queue = queue[1:]
		fmt.Print(s, " ")
		for _, v := range graph[s] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
}

func main() {
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"A", "D", "E"}
	graph["C"] = []string{"A", "F"}
	graph["D"] = []string{"B"}
	graph["E"] = []string{"B", "F"}
	graph["F"] = []string{"C", "E"}

	fmt.Println("DFS:")
	visited := make(map[string]bool)
	DFS(graph, "A", visited)
	fmt.Println("\nBFS:")
	BFS(graph, "A")
}
