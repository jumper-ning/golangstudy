// YAML 解析（gopkg.in/yaml.v2）
package main

import (
	"fmt"
	"os"
	"bytes"
	"path/filepath"

	"gopkg.in/yaml.v3"  // 空行为了将标准库和第三方库区分
)

// 定义结构体（对应 K8s Deployment YAML 核心字段）
type Deployment struct {
	ApiVersion string `yaml:"apiVersion"` 
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name         string `yaml:"name"`
		Namespace    string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Replicas int `yaml:"replicas"`
		Template struct {
			Spec struct {
				Containers []struct {
					Name  string `yaml:"name"`
					Image string `yaml:"image"`
				} `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}






// 1. 编码器（Encoder）是什么？
// yaml.NewEncoder() 是 yaml.v3 提供的流式序列化工具
// 标准库 yaml.Marshal = 一次性输出（不能改缩进）
// Encoder = 可配置、可控制、可自定义格式（能改缩进）
// 2. 最关键一行：enc.SetIndent(indent)
// 这就是实现 MarshalIndent 的灵魂
// 作用：设置 YAML 嵌套层级的缩进空格数
// 你可以传：2（行业默认）、4（企业常用） 、8（极少用）
// 3. bytes.Buffer 是干嘛的？
// 它是一块内存缓冲区，编码器把 YAML 写到这里，最后通过 buf.Bytes() 直接拿到完整格式化后的 YAML 字符串
// 相当于：不输出到文件，只在内存里生成格式化 YAML
// 4. 为什么要 enc.Close()？
// 必须调用，作用：刷新缓冲区，保证所有数据都写入完成，不调用会导致 YAML 不完整、格式错乱


// YamlMarshalIndent 模拟 MarshalIndent，返回带缩进的 YAML 字节数组
// MarshalIndent 模拟 json.MarshalIndent
// v: 要序列化的数据
// indent: 缩进空格数（2/4/8）
// interface{} 是 Go 中的一个特殊接口类型，它是所有类型的超集，允许你以一种通用的方式处理多种类型的数据。
func YamlMarshalIndent(v interface{}, indent int) ([]byte, error) {
	// 创建一个内存缓冲区（相当于一个内存文件）
	var buf bytes.Buffer
	// 创建 YAML 编码器，把数据写到缓冲区
	encoder := yaml.NewEncoder(&buf)
	// 核心 API：设置缩进（这就是实现格式化的关键）
	encoder.SetIndent(indent)
	// 编码器把结构体转成 YAML
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	_ = encoder.Close()
	return buf.Bytes(), nil
}


func main() {
	// 1. 读取 K8s YAML 文件（先创建一个 deployment.yaml）
	yamlContent := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deploy
  namespace: default
spec:
  selector:
    matchLabels:
      app: nginx-deploy
  replicas: 2
  template:
    spec:
      containers:
        - name: nginx
          image: nginx:1.21`
	
	
	// 获取windows系统变量USERPROFILE 当前用户的家目录
	winuser := os.Getenv("USERPROFILE")
	dir := filepath.Join(winuser, "Desktop", "deployment.yaml")
	// 写入测试 YAML 文件
	err := os.WriteFile(dir, []byte(yamlContent), 0644)
	if err != nil {
		fmt.Println("写入 YAML 文件失败", err)
		return
	}


	// 读取 YAML 文件内容
	data, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println("读取 YAMl 文件失败", err)
		return
	}

	// 2. 解析 YAMl 文件到结构体
	var dep Deployment
	err = yaml.Unmarshal(data, &dep)
	if err != nil {
		fmt.Println("解析 YAML 文件失败", err)
		return
	}
	fmt.Printf("解析后的 Deployment: \n名称: %s\n副本数: %d\n镜像: %s\n",
        dep.Metadata.Name, dep.Spec.Replicas, dep.Spec.Template.Spec.Containers[0].Image)
	
	
	// 3. 修改 YAMl 内容（运维常用：修改镜像、副本数）
	dep.Spec.Replicas = 3                                     // 副本数改为 3
	dep.Spec.Template.Spec.Containers[0].Image = "nginx:1.25" // 镜像升级


	// 4. 生成新的 YAML 内容
	newYAML, err := YamlMarshalIndent(dep, 2)
	if err != nil {
		fmt.Println("生成 YAML 失败:", err)
		return
	}
     
	dir2 := filepath.Join(winuser, "Desktop", "new_deployment.yaml")
	// 5. 写入新的 YAML 文件
	err = os.WriteFile(dir2, newYAML, 0644)
	if err != nil {
		fmt.Println("写入新 YAMl 文件失败: ", err)
		return
	}

	fmt.Println("\n修改后的 YAML 已写入 new_deployment.yaml")
	fmt.Println("新副本数: 3, 新镜像: nginx:1.25")
}
