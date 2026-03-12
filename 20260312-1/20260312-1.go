// 文件读写（os/io）
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)


// 运维场景1：读取文件（类比 Shell 的 cat pods.txt）
func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("读取文件失败: %v", err)
	}
	return string(content), nil
}

// 运维场景2：写入文件（类比 Shell 的 echo "xxx" > log.txt
func writeFile(filePath, content string) error {
	// 0644 是文件权限（运维常用）
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}
	return nil
}

func getTimeStamp() (string) {
	// 获取当前时间
	now := time.Now()
	
	// 输出当前时间
	// fmt.Println("当前时间:", now)
	
	// 提取年月日时分秒
	year, month, day := now.Date()
	hour, minute, second := now.Clock()
	times := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	return times
}

func main() {
	// 1. 写入测试文件（Pod 列表）
	podsContent := "nginx01/redis-01/mysql-01\n"
	path := filepath.Join("C:\\", "Users", "lenovo", "Desktop", "pods.txt")
	err := writeFile(path, podsContent)
	if err != nil {
		fmt.Println("错误: ", err)
		return
	}
	fmt.Println("成功写入 ", path)

	// 2. 读取测试文件
	content, err := readFile(path)
	if err != nil {
		fmt.Println("错误: ", err)
		return
	}
	fmt.Println("读取 pods.txt 内容: ", content)

	// 3. 写入巡检日志（追加模式，运维常用）
	logContent :=  getTimeStamp() + " - 节点巡检完成，所有 Pod 正常\n"
	// 打开文件（O_APPEND 追加，O_WRONLY 只写，O_CREATE 不存在则创建）
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("错误: ", err)
		return
	}
	defer f.Close() // 延迟关闭文件（运维必备：防止文件句柄泄露）
    _, err = f.WriteString(logContent)
	if err != nil {
		fmt.Println("错误: ", err)
		return
	}
	fmt.Println("成功追加巡检日志到 ", path)

}