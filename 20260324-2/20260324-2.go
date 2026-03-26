// sync.WaitGroup 等待组
package main

import (
	"fmt"
	"sync"
	"time"
)

// 检查节点状态（接受 WaitGroup 指针）
// wg *sync.WaitGroup：指针类型的等待组（必须传指针，否则是值拷贝，无法修改原计数器） wg 是自定义的变量名称
func checkNode(node string, wg *sync.WaitGroup) {
	// defer：函数退出前一定会执行（无论正常结束 / 异常）
	// wg.Done()：任务完成后，将 WaitGroup 计数器 -1
	// 必须写在函数开头，保证函数任何情况下退出，都能释放计数器
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("节点 %s 检查完成: 状态正常\n", node)
}

func main() {
	nodes := []string{"node-1", "node-2", "node-3", "node-4"}
	// 声明一个零值的 sync.WaitGroup，用于管理并发任务，跟踪所有 goroutine 是否执行完成。
	var wg sync.WaitGroup // 定义 WaitGroup

	fmt.Println("=== 并发检查 (WaitGroup) ===")
	start := time.Now()

	for _, node := range nodes {
		wg.Add(1) // 计数器 +1 （必须在 goruntine 启动前）告诉 WaitGroup：新增一个等待任务，计数器 +1
		// go 函数名()：启动一个 goroutine（轻量级线程）
		// 传入 &wg：必须传地址，保证所有 goroutine 操作同一个计数器
		go checkNode(node, &wg)
	}
	// 阻塞主 goroutine，直到计数器变为 0
	// 主程序会一直等在这里，直到所有节点检查完成
	wg.Wait() // 等待所有任务完成（计数器归 0 ）

	fmt.Printf("并发检查耗时: %v\n", time.Since(start))
	fmt.Println("所有节点检查完成！")
}

// 核心知识点总结
// sync.WaitGroup 作用等待一组 goroutine 全部执行完成，再继续主程序。
// 三个核心方法
// wg.Add(n)：计数器 +n（必须在 goroutine 前调用）
// wg.Done()：计数器 -1（配合 defer 使用）
// wg.Wait()：阻塞直到计数器 = 0
// 必须传指针&wg 传给 goroutine，否则是值拷贝，计数器不生效。
// 并发优势4 个任务各耗时 1 秒，并发执行总耗时 ≈ 1 秒。
