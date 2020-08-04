package main

import (
	"fmt"
)

//207. 课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacencyList, state := getAdjacencyList(prerequisites)
	for i := range adjacencyList {
		if hasRing(i, adjacencyList, &state) {
			return false
		}
	}
	return true
}

// 生成邻接表及状态表
func getAdjacencyList(prerequisites [][]int) (adjacencyList map[int]map[int]int, state map[int]int) {
	adjacencyList = make(map[int]map[int]int)
	state = make(map[int]int)
	for i := 0; i < len(prerequisites); i++ {
		if _, ok := adjacencyList[prerequisites[i][0]]; !ok {
			adjacencyList[prerequisites[i][0]] = make(map[int]int)
		}
		adjacencyList[prerequisites[i][0]][prerequisites[i][1]] = 1
		state[prerequisites[i][0]] = 0
	}
	return adjacencyList, state
}

// 监测是否有环
func hasRing(current int, adjacencyList map[int]map[int]int, state *map[int]int) bool {
	// 0 未监测，1 当前已监测，-1 无环
	// 优先监测状态表
	if v, ok := (*state)[current]; ok {
		if v == 1 {
			return true
		} else if v == -1 {
			return false
		} else {
			(*state)[current] = 1
		}
	} else {
		return false
	}
	// 若状态表中不存在结果，则遍历邻接表检查是否有环
	if ls, ok := adjacencyList[current]; ok {
		for next := range ls {
			if hasRing(next, adjacencyList, state) {
				return true
			}
		}
		(*state)[current] = -1
		return false
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}), true)
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}), false)
	fmt.Println(canFinish(4, [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}), false)
}
