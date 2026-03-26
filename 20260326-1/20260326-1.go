// sync.Mutex 锁（简单使用）
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 全局变量：日志文件句柄、互斥锁
var (
	logFile *os.File
	mutex   sync.Mutex
)

// 初始化日志文件
func initLog() error {
	winuser := os.Getenv("USERPROFILE")
	logfile := filepath.Join(winuser, "Desktop", "concurrent.log")
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
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
		fmt.Println("写入日志失败: ", err)
		return
	}
	fmt.Println("写入日志成功: ", content)
 }

 func main () {
	// 初始化日志文件
	err := initLog()
	if err != nil {
		fmt.Println("初始化日志失败: ", err)
		return
	}

	defer logFile.Close()

	var wg sync.WaitGroup

	// 开启 5 个 goroutine 并发写入日志
	for i :=1; i <= 5; i ++ {
		wg.Add(1)
		content := fmt.Sprintf("节点检查任务 %d 完成", i)
		go writeLog(content, &wg)
	}

	wg.Wait()
	fmt.Println("所有日志写入完成！")

 }