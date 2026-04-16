// 并发能力巩固
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
	"flag"
)

// 全局变量：日志文件句柄、互斥锁
var (
	logFile *os.File
	mutex   sync.Mutex
)

// 初始化日志文件
func initLog() error {
	winuser := os.Getenv("USERPROFILE")
	logpath := filepath.Join(winuser, "Desktop", "tools.log")
	f, err := os.OpenFile(logpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	logFile = f
	return nil
}

// 并发写入日志（不加锁会导致内容混乱）
func writeLog(content string, wg *sync.WaitGroup) {
	defer wg.Done()

	// 加锁（类比 Shell 的 flock 锁定文件）
	mutex.Lock()
	defer mutex.Unlock() // 函数结束解锁

	// 写入日志 
	// fmt.Sprintf 会先在内存里生成一个完整字符串
	// fmt.Fprintf 直接把格式化内容写到 io.Writer 里，不生成中间字符串，少一次字符串分配
	// _, err := logFile.WriteString(fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), content))
	_, err := fmt.Fprintf(logFile, "[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), content)
	if err != nil {
		fmt.Println("写入日志失败", err)
		return
	}
	fmt.Println("写入日志文件成功", content)
}

// 节点状态结构体
type NodeStatus struct {
	Node    string  // 节点名
	CPU     string  // CPU使用率
	Memory  string  // 内存使用率
	Disk    string  // 磁盘使用率
	Error   error   // 错误信息
}

// 检查节点 CPU 使用率

func checkCPU(node string, status *NodeStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	// 模拟结果
	status.CPU = fmt.Sprintf("%d%%", 10+int(node[5]-'0')*10)
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

// splitComma 切割逗号分隔的字符串
func splitCommand(s string) []string {
	var res []string
	start := 0
	for i, c := range s {
		if c == ',' {
			res = append(res, s[start:i])
			start = i + 1
		}
	}
	res = append(res, s[start:])
	return res
}

func main() {
	// 1. 定义一个字符串切片，接收 -nodes 参数，逗号分隔
	nodeList := flag.String("nodes", "", "需要检查的节点，逗号分隔，例如：node-1,node-2,node-3")
	// 2. 解析命令行参数
	flag.Parse()

	// 3. 判断是否传入参数
	if *nodeList == "" {
		fmt.Println("错误: 必须通过 --nodes 指定节点列表")
		// 打印帮助信息
		flag.Usage()
		return
	}
	// 4. 把逗号分隔的字符串切分成切片
	nodes := splitCommand(*nodeList)
	

	// 初始化日志文件
	err := initLog()
	if err != nil {
		fmt.Println("初始化日志失败", err)
		return
	}

	defer logFile.Close()

	var wg sync.WaitGroup

	// nodes := []string{"node-1", "node-2", "node-3", "node-4"}
	resultChan := make(chan NodeStatus, len(nodes))

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

	

	// ========== 新增：日志等待组 ==========
	var logWg sync.WaitGroup

	// 打印汇总结果
	logWg.Add(1)
	writeLog("=== 节点状态巡检开始 ===", &logWg)
	logWg.Wait()

	for status := range resultChan {
		var content string
		if status.Error != nil {
			// 构造错误日志内容
			content = fmt.Sprintf("节点 %s: 检查失败 - %v", status.Node, status.Error)
			fmt.Printf("节点 %s: 检查失败 - %v\n", status.Node, status.Error)
		} else {
			// 构造正常状态日志内容
			content = fmt.Sprintf("节点 %s: CPU=%s, 内存=%s, 磁盘=%s",
	        status.Node, status.CPU, status.Memory, status.Disk)
			fmt.Printf("节点 %s: CPU=%s, 内存=%s, 磁盘=%s\n",
	        status.Node, status.CPU, status.Memory, status.Disk)
		}
		logWg.Add(1)
		writeLog(content, &logWg)
	}
	//
	logWg.Wait()

	// 
	logWg.Add(1)
	writeLog(fmt.Sprintf("所有节点巡检完成，总耗时：%v", time.Since(start)), &logWg)
	logWg.Wait()
}

