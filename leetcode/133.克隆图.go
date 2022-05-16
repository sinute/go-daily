package leetcode

/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodes := make(map[int]*Node)
	return dfs133(nodes, node)
}

func dfs133(nodes map[int]*Node, node *Node) *Node {
	v, ok := nodes[node.Val]
	if !ok {
		nodes[node.Val] = &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
		for i := range node.Neighbors {
			nodes[node.Val].Neighbors = append(nodes[node.Val].Neighbors, dfs133(nodes, node.Neighbors[i]))
		}
		return nodes[node.Val]
	}
	return v
}

// @lc code=end
