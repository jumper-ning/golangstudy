// 函数定义、参数、返回值
package main

import "fmt"

// 运维场景：函数1 - 计算磁盘使用率（接收已用/总容量，返回使用率）
// 类比 Shell 的 function calc_disk_usage() { ... }

func calcDiskUsage(used float64, total int) float64 {
	return float64(used) / float64(total) * 100
}

// 运维场景：函数2 - 检查 Pod 状态（接收状态，返回是否正常+提示）
func checkPodStatus(status string) (bool, string) {
	if status == "Running" {
		return true, "Pod 运行正常"
	} else {
		return false, fmt.Sprintf("Pod 异常, 状态: %s", status)
	}
}

func main() {
	// 调用函数1
	usage := calcDiskUsage(79.3, 100)
	fmt.Printf("磁盘使用率: %.2f%%\n", usage)

	// 调用函数2
	isOk, msg := checkPodStatus("Error")
	fmt.Println("Pod 检查结果: ", isOk, "，提示: ", msg)

	isOk2, msg2 := checkPodStatus("Running")
	fmt.Println("Pod 检查结果: ", isOk2, "，提示: ", msg2)


}
