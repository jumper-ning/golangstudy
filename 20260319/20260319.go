// 理解 JSON 序列化 / 反序列化
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// 定义结构体（对应 K8s Pod 信息的 JSON 结构）
type Pod struct {
	// json:"name" 表示：当解析 JSON 时，JSON 中的 name 字段会映射到结构体的 Name 字段；当生成 JSON 时，结构体的 Name 字段会输出为 JSON 的 name 字段；
	Name      string `json:"name"`      // JSON 字段名 
	Namespace string `json:"namespace"` // JSON 字段名
	Status    string `json:"status"`    // JSON 字段名
	CPU       string `json:"cpu"`       // JSON 字段名
	Memory    string `json:"memory"`    // JSON 字段名
}

func main() {
	// 场景1：JSON 字符串解析为结构体（类比 Shell jq 解析JSON）
	fmt.Println("=== JSON 解析 ===")
	// 定义变量 jsonStr，使用反引号 ` 定义多行字符串，内容是一个标准的 JSON 字符串
	jsonStr := `{
	    "name": "nginx-01",
	    "namespace": "default",
	    "status": "Running",
	    "cpu": "100m",
	    "memory": "256Mi"
	}`
    // 声明一个类型为 Pod 的变量 pod，初始值是 Pod 结构体的零值（所有字段都是对应类型的零值，比如字符串字段为空字符串 ""）
	// 这个变量用于存储后续解析 JSON 后的结果
	var pod Pod
	// 反序列化（JSON->结构体）
	// json.Unmarshal：JSON 反序列化函数，作用是将 JSON 字节数据解析为 Go 结构体；
    // []byte(jsonStr)：将字符串 jsonStr 转换为字节切片（json.Unmarshal 要求输入是字节切片）；
    // &pod：传递 pod 变量的地址（指针），因为 json.Unmarshal 需要修改 pod 变量的内容，必须传指针才能修改原变量；
	err := json.Unmarshal([]byte(jsonStr), &pod)
	if err != nil {
		fmt.Println("解析 JSON 失败: ", err)
		return
	}
	fmt.Printf("解析后的 Pod 信息: \n名称: %s\n命名空间: %s\n状态: %s\nPod 的 CPU: %s\nPod 的 Memory: %s\n", pod.Name, pod.Namespace, pod.Status, pod.CPU, pod.Memory)


	// 场景2：结构体生成 JSON 字符串（类比 Shell 的 echo '{"name":"xxx"}'）
	fmt.Println("\n=== JSON 生成===")
	// 初始化 Pod 结构体
	newPod := Pod{
		Name:      "redis-01",
		Namespace: "default",
		Status:     "Pending",
		CPU:        "200m",
		Memory:     "512Mi",
	}

	// 序列化（结构体-> JSON，Indent 格式化输出）
	// := 只能在函数内部使用，并且要求左侧至少有一个新声明的变量。如果变量已经存在，则不能使用 := 来重新声明，否则会引发编译错误
	jsonData, err := json.MarshalIndent(newPod, "", " ")
	if err != nil {
		fmt.Println("生成 JSON 失败: ", err)
	}
	fmt.Println("生成的 JSON 配置: ")
	fmt.Println(string(jsonData))
     
	// 写入 JSON 文件
    // 获取windows系统变量USERPROFILE 当前用户的家目录
	winuser := os.Getenv("USERPROFILE")
	dir := filepath.Join(winuser, "Desktop", "pod_config.json")
	// := 相对的是 =，它仅用于赋值操作，不能用于变量声明
	err = os.WriteFile(dir, jsonData, 0644)
	if err != nil {
		fmt.Println("写入 JSON 文件失败: ", err)
	}
	fmt.Println("写入 JSON 文件成功")
}