// if/else 条件判断
package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	namespace := "dev"
	cmd := exec.Command("kubectl", "get", "pods", "-n", namespace, "-o", "name", "--no-headers")
	output, err := cmd.CombinedOutput()
	if err != nil {
		// 错误处理，对应 Shell 的 $? 判断
		fmt.Printf("执行命令失败: %v, 输出: %s", err, output)
	}
	podResult := strings.Split(strings.TrimSpace(string(output)), "\n")
	podCount := len(podResult)
	fmt.Println(podResult, "\npod数量: ", podCount)
	maxPods := 20

	if podCount > maxPods {
		fmt.Printf("警告：pod数量 %d 已经超过了阈值 %d", podCount, maxPods)
	} else if podCount == maxPods {
		fmt.Println("提示：pod数量和阈值相等")
	} else {
		fmt.Println("pod数量在阈值范围内")
	}
	getNodeStatus()
}

// 布尔值判断
func getNodeStatus() {
	var nodeOnline bool

	// cmd := exec.Command("kubectl", "get", "nodes", "--no-headers")
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	// 错误处理，对应 Shell 的 $? 判断
	// 	fmt.Printf("执行命令失败: %v, 输出: %s", err, output)
	// }
	// nodeStatus := strings.Split(strings.TrimSpace(string(output)), "\n")
	nodeStatus := []string{"Ready", "Ready", "Ready"}
	for _, status := range nodeStatus {
		if status == "Ready" {
			nodeOnline = true
		} else {
			nodeOnline = false
		}
		// nodeOnline := false
		if !nodeOnline {
			fmt.Println("警告节点离线")
		} else {
			fmt.Println("提示：节点在线")
		}
	}

}
