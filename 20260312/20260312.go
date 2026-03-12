// os/exec 执行系统命令

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// 运维场景: 执行系统命令（类似 Shell 的 cmd=$(kubectl version)）
func runCommand(cmd string, args ...string) (string, error) {
	// 构建命令
	c := exec.Command(cmd, args...)
	// 执行并获取输出（CombinedOutput 包含 stdout + stderr）
	output, err := c.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("命令执行失败: %v, 输出: %s", err, output)
	}
	return string(output), nil
}

// 示例3：合并成一个函数调用，将命令作为参数传进去
func defReturn(cmd string, args ...string) {
	// 步骤1：拼接命令和所有参数成一个完整的字符串
	// cmd 是主命令（如 kubectl），args 是可变参数列表（如 version --short）
	commandParts := append([]string{cmd}, args...) // 合并成 ["kubectl", "version", "--short"]
	commandStr := strings.Join(commandParts, " ")  // 拼接成 "kubectl version --short"
	fmt.Printf("=== 执行 %s ===\n", commandStr)
	commandOut, err := runCommand(cmd, args...)
	if err != nil {
		fmt.Println("错误: ", err)
	} else {
		fmt.Println(commandOut)
	}
}

func main() {
	defReturn("kubectl", "version", "--short")
	// // 示例1: 执行 kubectl version
	// fmt.Println("=== 执行 kubectl version ===")
	// kubectlOut, err := runCommand("kubectl", "version", "--short")
	// if err != nil {
	// 	fmt.Println("错误: ", err)
	// } else {
	// 	fmt.Println(kubectlOut)
	// }

	// // 示例2: 执行 df -h （windows 替换为 dir）
	// fmt.Println("=== 执行 df -h ===")
	// // Windows 请替换为：runCommand("dir")
	// dfOut, err := runCommand("df", "-h")
	// // dir 并不是一个独立的可执行文件（没有 dir.exe），而是 Windows 命令提示符（cmd.exe）内置的命令
	// // dfOut, err := runCommand("cmd", "/c", "dir")
	// if err != nil {
	// 	fmt.Println("错误", err)
	// } else {
	// 	fmt.Println(dfOut)
	// }
}
