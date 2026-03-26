// channel 通道基础、用 channel 接收 goroutine 执行结果
package main

import (
	"fmt"
	"time"
)

// 检查节点磁盘使用率(结果通过 channel 返回)
// chan 是一种类型，可以看作是一个通信机制，用于在不同的 goroutine 之间传递数据
func checkDisk(node string, ch chan string) {
	time.Sleep(1 * time.Second)
	// 模拟磁盘使用率
	result := fmt.Sprintf("节点 %s 磁盘使用率: 80%%", node)
	ch <- result // 将结果发送到通道
}

func main() {
	nodes := []string{"node-1", "node-2", "node-3"}
	// 通道必须通过 make() 函数初始化，创建通道（缓冲大小为 3，避免阻塞）
	ch := make(chan string, len(nodes))

	// 开启 goroutine 执行任务
	for _, node := range nodes {
		go checkDisk(node, ch)
	}

	// 接受通道结果
	fmt.Println("=== 节点磁盘检查结果 ===")
	for i := 0; i < len(nodes); i++ {
		result := <- ch //从通道接收结果
		fmt.Println(result)
	}

	close(ch) // 关闭通道 （可选，防止内存泄露）
	fmt.Println("所有结果接收完成！")
}