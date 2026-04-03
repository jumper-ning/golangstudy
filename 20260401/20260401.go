// 整合 goroutine + WaitGroup + exec

package main

import (
	"fmt"
	"os/exec"
	"sync"
)

// 并发执行 kubectl 命令， 检查 Pod 状态
func checkPod (podName string, namespace string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()

	// 构建命令：kubectl  get pod <pod> -n <ns> -o jsonpath='{.status.phase}'
	cmd := exec.Command("kubectl", "get", "pod", podName, "-n", namespace, "-o", "jsonpath={.status.phase}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		resultChan <- fmt.Sprintf("Pod %s/%s: 检查失败 - %v, 输出：%s", namespace, podName, err, output)
		return
	}

	status := string(output)
	resultChan <- fmt.Sprintf("Pod %s/%s: 状态 - %s", namespace, podName, status)
}

func main() {
	// 待检查的 Pod 列表
	pods := []struct {
		name      string
		namespace string	
	}{
		{"nginx-01", "default"},
		{"redis-01", "default"},
		{"coredns-0", "kube-system"},
		{"calico-node-xxxx", "kube-system"}, // 替换为实际节点名
	}


	var wg sync.WaitGroup
	resultChan := make(chan string, len(pods))

	// 开启goroutine 并发检查
	for _, p := range pods {
		wg.Add(1)
		go checkPod(p.name, p.namespace, &wg, resultChan)
	}

	// 等待所有检查完成
	wg.Wait()
	close(resultChan)

	// 打印结果
	fmt.Println("=== Pod 状态检查结果 ===")
	for result := range resultChan {
		fmt.Println(result)
	}
}
