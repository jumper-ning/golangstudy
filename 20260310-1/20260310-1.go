// 切片 slice、map
package main

import "fmt"

func main() {
	// 1. 切片 （类比shell 的nodes=("node-1" "node-2")）
	fmt.Println("=== 切片操作 ===")
	pods := []string{"nginx-01", "redis-01", "mysql-01"}
	// 添加元素
	pods = append(pods, "kafka-01")
	// 遍历
	for _, pod := range pods {
		fmt.Println("Pod名称: ", pod)
	}

	// 2. map（类比 shell 的 declare -A node_disk=([node-1]="80%" [node-2]="70%") -A 声明‌关联数组‌（键值对形式）‌）
	fmt.Println("\n=== Map 操作 ===")
	nodeDisk := map[string]string{
		"node-1": "80%",
		"node-2": "70%",
		"node-3": "60%",
	}
	// 添加键值对
	nodeDisk["node-4"] = "90%"
	// 删除键值对
	delete(nodeDisk, "node-2")
	// 遍历
	for node, usage := range nodeDisk {
		fmt.Printf("节点 %s 的磁盘使用率是 %s\n", node, usage)
	}
}
