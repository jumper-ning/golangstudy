// 第一周总结
package main

import (
	"fmt"
)

// 函数：检查节点状态
func checkNodesStatus(node string, podCount int) string {
	if podCount > 20 {
		gapCount := podCount - 20
		return fmt.Sprintf("节点: %s Pod 数量高过指标 (%d)，状态异常", node, gapCount)
	} else {
		gapCount := 20 - podCount 
		return fmt.Sprintf("节点: %s Pod 数量低于指标 (%d)，状态正常", node, gapCount)
	}
}

func main() {
	// 定义节点列表
	nodes := map[string]int{
		"node-1": 15,
		"node-2": 25,
		"node-3": 18,
	}

	// 遍历节点，检查状态
	fmt.Println("=== 节点状态巡检 ===")
	for node, count := range nodes {
		result := checkNodesStatus(node, count)
		fmt.Println(result)
	}
}