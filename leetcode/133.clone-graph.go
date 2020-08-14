package main

import "fmt"

//133. 克隆图
/**
 * Definition for a Node.
 */
type Node struct {
    Val       int
    Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    nodes := make(map[int]*Node)
    return dfs(nodes, node)
}

func dfs(nodes map[int]*Node, node *Node) *Node {
    v, ok := nodes[node.Val]
    if !ok {
        nodes[node.Val] = &Node{
            Val:       node.Val,
            Neighbors: []*Node{},
        }
        for i := range node.Neighbors {
            nodes[node.Val].Neighbors = append(nodes[node.Val].Neighbors, dfs(nodes, node.Neighbors[i]))
        }
        return nodes[node.Val]
    }
    return v
}

func main() {
    node1 := &Node{
        Val:       1,
        Neighbors: nil,
    }
    node2 := &Node{
        Val:       2,
        Neighbors: nil,
    }
    node3 := &Node{
        Val:       3,
        Neighbors: nil,
    }
    node4 := &Node{
        Val:       4,
        Neighbors: nil,
    }
    node1.Neighbors = []*Node{node2, node4}
    node2.Neighbors = []*Node{node1, node3}
    node3.Neighbors = []*Node{node2, node4}
    node4.Neighbors = []*Node{node1, node3}

    n := cloneGraph(node1)
    fmt.Println(n)
}
