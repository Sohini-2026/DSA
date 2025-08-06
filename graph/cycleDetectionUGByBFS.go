package graph

import "fmt"

func IsCycleInUndirectedGraphByBFS(graph map[int][]int) bool {
	visited := make(map[int]bool)
	type NodeParent struct {
		node, parent int
	}
	queue := []NodeParent{}

	for node := range graph {
		if !visited[node] { //you guarantee that every node (and thus every component) is checked for cycles.
			queue = append(queue, NodeParent{node, -1})
			visited[node] = true

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				for _, neighbor := range graph[current.node] {
					fmt.Print("Visiting node:", current.node, "with parent:", current.parent, "\n", neighbor, "is a neighbor\n")
					if !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, NodeParent{neighbor, current.node})
					} else if neighbor != current.parent {
						fmt.Print("Cycle detected between node:", current.node, "and neighbor:", neighbor, "\n", current.parent, "is the parent\n")
						return true // Cycle found
					}
				}
			}
		}
	}
	return false // No cycle found
}

//While this BFS approach is less common for cycle detection in undirected graphs,
//it can be useful in certain scenarios, especially when you want to avoid recursion or when the graph is large and you want to manage memory usage more effectively.
//Note : Adjacency list : 1{2,5},2,{1,3},3{2,4},4{3,5},5{1,4}
//Process of execution :
// 	Visiting node:1with parent:-1
// 2is a neighbor
// Visiting node:1with parent:-1
// 5is a neighbor
// Visiting node:2with parent:1
// 1is a neighbor
// Visiting node:2with parent:1
// 3is a neighbor
// Visiting node:5with parent:1
// 4is a neighbor
// Visiting node:5with parent:1
// 1is a neighbor
// Visiting node:3with parent:2
// 2is a neighbor
// Visiting node:3with parent:2
// 4is a neighbor
// Cycle detected between node:3and neighbor:4
// 2is the parent
//Its BFS so 1{2,5} its going to node 2 and then node 5
