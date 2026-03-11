// error 错误处理
package main

import (
	"fmt"
)

func getPodInfo(podName string) (string, error) {
	// 模拟pod列表
	vaildPods := []string{"nginx-01", "redis-01"}

	// 检查pod是否存在
	for _, p := range vaildPods {
		if p == podName {
			// Sprintf函数用于格式化字符串到一个新的字符串，而不是输出到标准输出，一般用于return
			// Printf函数用于格式化输出到标准输出（通常是控制台）。它类似于C语言中的printf函数
			// Println函数用于将数据输出到标准输出，并且会在输出的末尾自动添加一个换行符。它主要用于简单的打印操作。
			return fmt.Sprintf("Pod %s 信息: 运行中， CPU: 10%%", podName), nil
		}
	}
	// 返回自定义错误（类似于 shell 的 echo "error: xxx" && exit 1）
	// 输出错误信息为静态消息是推荐使用errors.new，如果错误信息包含动态数据，比如变量，推荐使用fmt.Errorf
	// fmt.Errorf可能会导致性能问题，errors.New性能可能会更好
	// return "", errors.New(fmt.Sprintf("Pod %s 不存在", podName))
	return "", fmt.Errorf("Pod %s 不存在", podName)
}

func main() {
	// 调用函数，处理错误
	pod := "redis-01"
	podInfo, err := getPodInfo(pod)
	if err != nil {
		//错误处理（运维必备：打印错误、记录日志）
		fmt.Println("错误: ", err)
		// 可以在这里加日志写入、告警等逻辑
		return
	}
	// 无错误泽打印信息
	fmt.Println("Pod 信息: ", podInfo)
}
