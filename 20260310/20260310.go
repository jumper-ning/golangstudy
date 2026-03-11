// for 循环、range
package main

import "fmt"

func main() {
	// 基础循环（类比shell 的for ((i=1; i<=5; i++))）
	fmt.Println("==== 循环打印 1-5 ====")
	for i := 1; i <= 5; i++ {
		fmt.Println("循环次数: ", i)
	}

	// // 2. 运维场景：遍历节点列表（类比 Shell 的 for node in node-1 node-2 node-3）
	fmt.Println("\n=== 遍历节点列表 ===")
	nodes := []string{"node1", "node2", "node3"}
	for index, node := range nodes {
		fmt.Printf("索引: %d, 节点名称: %s\n", index, node)
	}

	// 3. 类似while循环
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 4. 无限循环（需手动终止，运维少用）
	// for {
	// 	fmt.Println("无限循环")
	// 	break // 终止循环
	// }
}
