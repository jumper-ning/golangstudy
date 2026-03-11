// Go 版：遍历打印指定命名空间的 Pod（类比 Shell：kubectl get pods -n default | awk 'NR>1{print $1}'）
package main

import (
	"fmt"
	"os/exec" // 执行系统命令，对应 Shell 的命令执行
	"strings"
)

// 封装成函数，对应 Shell 函数
func getPods(namespace string) ([]string, error) {
	// 执行 kubectl 命令，对应 Shell 的 kubectl 调用
	cmd := exec.Command("kubectl", "get", "pods", "-n", namespace, "-o", "name")
	output, err := cmd.CombinedOutput()
	var result []string
	if err != nil {
		// 错误处理，对应 Shell 的 $? 判断
		return nil, fmt.Errorf("执行命令失败: %v, 输出: %s", err, output)
	}

	// 处理输出，对应 Shell 的 awk/sedk
	pods := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, pod := range pods {
		parts := strings.Split(pod, "/")
		secondPart := parts[1]
		result = append(result, secondPart)
	}
	
	return result, nil
}

func main() {
	var ns string = "dev"
	pods, err := getPods(ns)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}

	// 循环遍历，对应 Shell 的 for 循环
	for _, pod := range pods {
		fmt.Println(ns,"命名空间下的Pod 名称:", pod)
	}
}