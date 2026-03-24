// goroutine 基础
package main

import (
	"fmt"
	"time"
)

// 运维场景：检查单个节点的状态（模拟耗时操作）
func  checkNode(node string) {
	// 模拟网络请求耗时 1 秒
	time.Sleep(1 * time.Second)
	fmt.Printf("节点 %s 检查完成: 状态正常\n", node)
}

func main() {
	nodes := []string{"node-1", "node-2", "node-3"}
	fmt.Println("=== 串行检查（耗时久） ===")
	// 串行执行（类比 Shell 的for node in ...;do check $node; done）
	start := time.Now()
	for _, node := range nodes {
		checkNode(node)
	}
	fmt.Printf("串行检查耗时: %v\n", time.Since(start))


	fmt.Println("\n=== 并行检查(goroutine, 耗时短) ===")
	// 并发执行（类比 Shell 的 for node in ...;do check $node & done）
	start = time.Now()
	for _, node := range nodes {
		go checkNode(node) // 加 go 关键字， 开启 goroutine
	}

	// 等待所有 goroutine 完成（临时方案，后续用 WaitGroup）
	time.Sleep(2 * time.Second)
	fmt.Printf("并发检查耗时: %v\n", time.Since(start))
}
