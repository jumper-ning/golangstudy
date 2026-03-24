// sync.WaitGroup 等待组
package main

import (
	"fmt"
	"sync"
	"time"
)

// 检查节点状态（接受 WaitGroup 指针）
func checkNode(node string, wg *sync.WaitGroup) {
	defer wg.Done() // 任务完成后， 计数器 -1（必须放在开头）
	time.Sleep(1 * time.Second)
	fmt.Printf("节点 %s 检查完成: 状态正常\n", node)
}

func main() {
	nodes := []string{"node-1", "node-2", "node-3", "node-4"}
	var wg sync.WaitGroup // 定义 WaitGroup

	fmt.Println("=== 并发检查 (WaitGroup) ===")
	start := time.Now()

	for _, node := range nodes {
		wg.Add(1) // 计数器 +1 （必须在 goruntine 启动前）
		go checkNode(node, &wg)
	}

	wg.Wait() // 等待所有任务完成（计数器归 0 ）

	fmt.Printf("并发检查耗时: %v\n", time.Since(start))
	fmt.Println("所有节点检查完成！")
}