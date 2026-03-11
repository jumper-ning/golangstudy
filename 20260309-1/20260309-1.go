// 变量、常量、基本类型
package main

import "fmt"

func main() {
    // 1. 变量定义（类比 Shell 的 node_name="node-1"）
    var nodeName string = "node-1" // 完整声明
	var svctype string = "nodePort"
    // svctype := "clusterip"
    // var podCount int = 10
    podCount := 10                // 短声明（运维常用）
    var isRunning bool = true     // 布尔类型

    // 2. 常量定义（固定值，不可修改）
    const namespace = "default"

    // 3. 打印变量
    fmt.Println("节点名称是字符串：", nodeName)
    fmt.Println("Pod 数量是整数：", podCount)
    fmt.Println("是否运行是布尔值：", isRunning)
    fmt.Println("命名空间是常量：", namespace)
    fmt.Println("svc类型：", svctype)

    // 4. 修改变量（常量不能改）
    podCount = 15
    fmt.Println("修改后 Pod 数量：", podCount)
}