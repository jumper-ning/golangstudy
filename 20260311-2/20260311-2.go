// flag 命令行参数
package main

import "fmt"
import "flag"

func main() {
	// 定义命令行参数（类比 shell 的while getopts "n:p:" opt; do ...）
	// -n 指定命名空间，默认default
	namespace := flag.String("n", "default", "指定 K8s 命名空间")
	// -p 指定 Pod 名称（必填）
	podName := flag.String("p", "", "指定 Pod 名称（必填）")
	// -v 显示版本
	version := flag.Bool("v", false, "显示版本号")

	// 解析参数
	flag.Parse()

	// 处理版本参数
	if *version {
		fmt.Println("k8s-pod-tool v1.0.0")
		return
	}

	// 检查必填参数
	if *podName == "" {
		fmt.Println("错误: 必须通过 -p 指定 Pod 名称")
		// 打印帮助信息
		flag.Usage()
		return
	}

	// 业务逻辑: 模拟查询 Pod 信息
	fmt.Printf("查询命名空间 %s 下的 Pod %s 信息...\n", *namespace, *podName)
	fmt.Printf("Pod %s/%s 状态: Running\n", *namespace, *podName)
}