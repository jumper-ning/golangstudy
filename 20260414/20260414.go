// 整合 goroutine + WaitGroup + channel + 错误处理

package main

import (
	"fmt"
	"sync"
	"time"
)

// 节点状态结构体
type NodeStatus struct {
	Node   string  // 节点名
	CPU    string  // CPU使用率
	Memory string  // 内存使用率
	Disk   string  // 磁盘使用率
	Error  error   // 错误信息
}

// 检查节点 CPU 使用率
func checkCPU(node string, status *NodeStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	// 模拟结果
	status.CPU = fmt.Sprintf("%d%%",10+int(node[5]-'0')*10)
}

// 检查节点内存使用率
func checkMemory(node string, status *NodeStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(600 * time.Millisecond)
	status.Memory = fmt.Sprintf("%d%%", 20+int(node[5]-'0')*10)
}

// 检查节点磁盘使用率
func checkDisk(node string, status *NodeStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(700 * time.Millisecond)
	// 模拟 node-3 磁盘检查失败
	if node == "node-3" {
		status.Error = fmt.Errorf("磁盘检查超时")
	}
	status.Disk = fmt.Sprintf("%d%%", 30+int(node[5]-'0')*10)
}

// 检查单个节点所有指标
func checkNodeAll(node string, resultChan chan NodeStatus) {
	status := NodeStatus{Node: node}
	var wg sync.WaitGroup

	// 并发检查 CPU/内存/磁盘
	wg.Add(3)
	go checkCPU(node, &status, &wg)
	go checkMemory(node, &status, &wg)
	go checkDisk(node, &status, &wg)

	wg.Wait()
	resultChan <- status
}

func main() {
	nodes := []string{"node-1", "node-2", "node-3", "node-4"}
	resultChan := make(chan NodeStatus, len(nodes))
	var wg sync.WaitGroup

	fmt.Println("开始并发检查节点状态...")
	start := time.Now()

	// 开启 goroutine 检查每个节点
	for _, node := range nodes {
		wg.Add(1)
		go func (n string)  {
			defer wg.Done()
			checkNodeAll(n, resultChan)
		}(node)
	}

	wg.Wait()
	close(resultChan)

	// 打印汇总结果
	fmt.Println("\n=== 节点状态巡检汇总 ===")
	fmt.Printf("总耗时: %v\n", time.Since(start))
	for status := range resultChan {
		if status.Error != nil {
			fmt.Printf("节点 %s: 检查失败 - %v\n", status.Node, status.Error)
			continue
		}
		fmt.Printf("节点 %s: CPU=%s, 内存=%s, 磁盘=%s\n",
	        status.Node, status.CPU, status.Memory, status.Disk)
	}
}