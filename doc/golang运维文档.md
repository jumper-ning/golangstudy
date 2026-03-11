Go 语言 4 周每日学习清单 + 配套示例代码
适用：会 Shell、K8s 运维、Windows + VSCode目标：每天 1 小时，代码直接复制就能跑，学完掌握 Go 通用编程能力
🎯 通用规则
每天：20 分钟学习 + 40 分钟写代码
代码：复制到 VSCode → 保存为 dayX.go → 执行 go run dayX.go
必须：运行成功、修改参数测试、记录问题
📅 第 1 周：Go 基础入门（打通语法）
第 1 天：环境搭建 + 第一个程序
任务清单
 安装 Go，执行 go env -w GOPROXY=https://goproxy.cn,direct
 VSCode 安装 Go 插件，执行 Go: Install/Update Tools 全选安装
 编写并运行 hello.go
 打卡：运行成功输出 Hello Go
示例代码（day1.go）
go
运行
package main

import "fmt"

func main() {
    // 第一个 Go 程序，类比 Shell 的 echo "Hello Go"
    fmt.Println("Hello Go")
    // 运维专属：打印 K8s 提示
    fmt.Println("K8s 运维 Go 学习开始 🚀")
}
第 2 天：变量、常量、基本类型
任务清单
 理解变量（var / 短声明）、常量（const）
 练习 int/string/bool 类型定义
 修改变量值，观察输出变化
示例代码（day2.go）
go
运行
package main

import "fmt"

func main() {
    // 1. 变量定义（类比 Shell 的 node_name="node-1"）
    var nodeName string = "node-1" // 完整声明
    podCount := 10                // 短声明（运维常用）
    var isRunning bool = true     // 布尔类型

    // 2. 常量定义（固定值，不可修改）
    const namespace = "default"

    // 3. 打印变量
    fmt.Println("节点名称：", nodeName)
    fmt.Println("Pod 数量：", podCount)
    fmt.Println("是否运行：", isRunning)
    fmt.Println("命名空间：", namespace)

    // 4. 修改变量（常量不能改）
    podCount = 15
    fmt.Println("修改后 Pod 数量：", podCount)
}
第 3 天：if/else 条件判断
任务清单
 理解 if/else 语法
 练习：判断 Pod 数量、节点状态
 类比 Shell 的 if [] 写法
示例代码（day3.go）
go
运行
package main

import "fmt"

func main() {
    // 运维场景：判断 Pod 数量是否超标
    podCount := 25
    maxPods := 20

    // 条件判断（类比 Shell 的 if [ $podCount -gt $maxPods ]）
    if podCount > maxPods {
        fmt.Printf("警告：Pod 数量 %d 超过阈值 %d\n", podCount, maxPods)
    } else if podCount == maxPods {
        fmt.Println("提示：Pod 数量等于阈值")
    } else {
        fmt.Println("正常：Pod 数量在阈值范围内")
    }

    // 布尔值判断（节点是否在线）
    nodeOnline := false
    if !nodeOnline {
        fmt.Println("警告：节点离线，请检查！")
    }
}
第 4 天：for 循环、range
任务清单
 理解 for 循环（Go 唯一循环）、range 遍历
 练习：循环打印节点列表、遍历字符串
 类比 Shell 的 for i in {1..5}
示例代码（day4.go）
go
运行
package main

import "fmt"

func main() {
    // 1. 基础循环（类比 Shell 的 for ((i=1; i<=5; i++))）
    fmt.Println("=== 循环打印 1-5 ===")
    for i := 1; i <= 5; i++ {
        fmt.Println("循环次数：", i)
    }

    // 2. 运维场景：遍历节点列表（类比 Shell 的 for node in node-1 node-2 node-3）
    fmt.Println("\n=== 遍历节点列表 ===")
    nodes := []string{"node-1", "node-2", "node-3", "node-4"}
    for index, node := range nodes {
        fmt.Printf("索引：%d，节点名称：%s\n", index, node)
    }

    // 3. 无限循环（需手动终止，运维少用）
    // for {
    //     fmt.Println("无限循环")
    //     break // 终止循环
    // }
}
第 5 天：切片 slice、map
任务清单
 理解切片（动态数组）、map（键值对）
 练习：添加 / 删除 / 遍历切片和 map
 类比 Shell 数组 / 关联数组
示例代码（day5.go）
go
运行
package main

import "fmt"

func main() {
    // 1. 切片（类比 Shell 的 nodes=("node-1" "node-2")）
    fmt.Println("=== 切片操作 ===")
    pods := []string{"nginx-01", "redis-01", "mysql-01"}
    // 添加元素
    pods = append(pods, "kafka-01")
    // 遍历
    for _, pod := range pods {
        fmt.Println("Pod 名称：", pod)
    }

    // 2. map（类比 Shell 的 declare -A node_disk=([node-1]="80%" [node-2]="70%")）
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
        fmt.Printf("节点 %s 磁盘使用率：%s\n", node, usage)
    }
}
第 6 天：函数定义、参数、返回值
任务清单
 理解函数定义、参数、返回值
 练习：写计算磁盘使用率、检查 Pod 状态的函数
 类比 Shell 函数
示例代码（day6.go）
go
运行
package main

import "fmt"

// 运维场景：函数1 - 计算磁盘使用率（接收已用/总容量，返回使用率）
// 类比 Shell 的 function calc_disk_usage() { ... }
func calcDiskUsage(used, total int) float64 {
    return float64(used) / float64(total) * 100
}




func checkPodStatus(status string) (bool, string) {
    if status == "Running" {
        return true, "Pod 运行正常"
    } else {
        return false, fmt.Sprintf("Pod 异常，状态：%s", status)
    }
}

func main() {
    // 调用函数1
    usage := calcDiskUsage(80, 100)
    fmt.Printf("磁盘使用率：%.2f%%\n", usage)

    // 调用函数2
    isOk, msg := checkPodStatus("Error")
    fmt.Println("Pod 检查结果：", isOk, "，提示：", msg)

    isOk2, msg2 := checkPodStatus("Running")
    fmt.Println("Pod 检查结果：", isOk2, "，提示：", msg2)
}
第 7 天（周总结）：综合小脚本
任务清单
 复习本周所有语法
 运行并修改综合脚本
 打卡：代码运行成功
示例代码（day7.go）
go
运行
package main

import "fmt"

// 函数：检查节点状态
func checkNodeStatus(node string, podCount int) string {
    if podCount > 20 {
        return fmt.Sprintf("节点 %s：Pod 数量超标（%d），状态异常", node, podCount)
    } else {
        return fmt.Sprintf("节点 %s：Pod 数量正常（%d），状态正常", node, podCount)
    }
}

func main() {
    // 定义节点列表
    nodes := map[string]int{
        "node-1": 15,
        "node-2": 25,
        "node-3": 18,
    }

    // 遍历节点，检查状态
    fmt.Println("=== 节点状态巡检 ===")
    for node, count := range nodes {
        result := checkNodeStatus(node, count)
        fmt.Println(result)
    }
}
📅 第 2 周：运维核心能力（能用、够用）
第 8 天：error 错误处理
任务清单
 理解 error 类型、错误返回 / 判断
 练习：给函数加错误处理
 不使用 _ 忽略错误
示例代码（day8.go）
go
运行
package main

import (
    "errors"
    "fmt"
)

// 运维场景：获取 Pod 信息（不存在则返回错误）
func getPodInfo(podName string) (string, error) {
    // 模拟 Pod 列表
    validPods := []string{"nginx-01", "redis-01"}

    // 检查 Pod 是否存在
    for _, p := range validPods {
        if p == podName {
            return fmt.Sprintf("Pod %s 信息：运行中，CPU：10%", podName), nil
        }
    }

    // 返回自定义错误（类比 Shell 的 echo "error: xxx" && exit 1）
    return "", errors.New(fmt.Sprintf("Pod %s 不存在", podName))
}

func main() {
    // 调用函数，处理错误
    podInfo, err := getPodInfo("mysql-01")
    if err != nil {
        // 错误处理（运维必备：打印错误、记录日志）
        fmt.Println("错误：", err)
        // 可以在这里加日志写入、告警等逻辑
        return
    }

    // 无错误则打印信息
    fmt.Println("Pod 信息：", podInfo)
}
第 9 天：flag 命令行参数
任务清单
 理解 flag 包使用
 练习：写支持 -n（命名空间）、-p（Pod 名）的工具
 类比 Shell 的 getopts
示例代码（day9.go）
go
运行
package main

import (
    "flag"
    "fmt"
)

func main() {
    // 定义命令行参数（类比 Shell 的 while getopts "n:p:" opt; do ...）
    // -n 指定命名空间，默认 default
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
        fmt.Println("错误：必须通过 -p 指定 Pod 名称")
        flag.Usage() // 打印帮助信息
        return
    }

    // 业务逻辑：模拟查询 Pod 信息
    fmt.Printf("查询命名空间 %s 下的 Pod %s 信息...\n", *namespace, *podName)
    fmt.Printf("Pod %s/%s 状态：Running\n", *namespace, *podName)
}
运行测试
bash
运行
# 查看帮助
go run day9.go -h
# 显示版本
go run day9.go -v
# 查询 Pod
go run day9.go -n kube-system -p coredns-0
# 缺少必填参数（测试错误处理）
go run day9.go -n default
第 10 天：os/exec 执行系统命令
任务清单
 理解 os/exec 执行系统命令
 练习：执行 kubectl version、df 等命令
 获取命令输出、处理错误
示例代码（day10.go）
go
运行
package main

import (
    "fmt"
    "os/exec"
)

// 运维场景：执行系统命令（类比 Shell 的 cmd=$(kubectl version)）
func runCommand(cmd string, args ...string) (string, error) {
    // 构建命令
    c := exec.Command(cmd, args...)
    // 执行并获取输出（CombinedOutput 包含 stdout + stderr）
    output, err := c.CombinedOutput()
    if err != nil {
        return "", fmt.Errorf("命令执行失败：%v，输出：%s", err, output)
    }

    return string(output), nil
}

func main() {
    // 示例1：执行 kubectl version
    fmt.Println("=== 执行 kubectl version ===")
    kubectlOut, err := runCommand("kubectl", "version", "--short")
    if err != nil {
        fmt.Println("错误：", err)
    } else {
        fmt.Println(kubectlOut)
    }

    // 示例2：执行 df -h（Windows 替换为 dir）
    fmt.Println("=== 执行 df -h ===")
    // Windows 请替换为：runCommand("dir")
    dfOut, err := runCommand("df", "-h")
    if err != nil {
        fmt.Println("错误：", err)
    } else {
        fmt.Println(dfOut)
    }
}
第 11 天：文件读写（os/io）
任务清单
 理解文件读取、写入
 练习：读取 Pod 列表文件、写入巡检日志
 类比 Shell 的 cat/echo >
示例代码（day11.go）
go
运行
package main

import (
    "fmt"
    "os"
)

// 运维场景1：读取文件（类比 Shell 的 cat pods.txt）
func readFile(filePath string) (string, error) {
    content, err := os.ReadFile(filePath)
    if err != nil {
        return "", fmt.Errorf("读取文件失败：%v", err)
    }
    return string(content), nil
}

// 运维场景2：写入文件（类比 Shell 的 echo "xxx" > log.txt）
func writeFile(filePath, content string) error {
    // 0644 是文件权限（运维常用）
    err := os.WriteFile(filePath, []byte(content), 0644)
    if err != nil {
        return fmt.Errorf("写入文件失败：%v", err)
    }
    return nil
}

func main() {
    // 1. 写入测试文件（Pod 列表）
    podsContent := "nginx-01\nredis-01\nmysql-01\n"
    err := writeFile("pods.txt", podsContent)
    if err != nil {
        fmt.Println("错误：", err)
        return
    }
    fmt.Println("成功写入 pods.txt")

    // 2. 读取测试文件
    content, err := readFile("pods.txt")
    if err != nil {
        fmt.Println("错误：", err)
        return
    }
    fmt.Println("读取 pods.txt 内容：")
    fmt.Println(content)

    // 3. 写入巡检日志（追加模式，运维常用）
    logContent := "2026-03-09 10:00:00 - 节点巡检完成，所有 Pod 正常\n"
    // 打开文件（O_APPEND 追加，O_WRONLY 只写，O_CREATE 不存在则创建）
    f, err := os.OpenFile("k8s_check.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("错误：", err)
        return
    }
    defer f.Close() // 延迟关闭文件（运维必备：防止文件句柄泄露）
    _, err = f.WriteString(logContent)
    if err != nil {
        fmt.Println("错误：", err)
        return
    }
    fmt.Println("成功追加巡检日志到 k8s_check.log")
}
第 12 天：JSON 解析与生成
任务清单
 理解 JSON 序列化 / 反序列化
 练习：解析 Pod JSON 信息、生成 JSON 配置
 对接 K8s API 场景
示例代码（day12.go）
go
运行
package main

import (
    "encoding/json"
    "fmt"
)

// 定义结构体（对应 K8s Pod 信息的 JSON 结构）
type Pod struct {
    Name      string `json:"name"`      // JSON 字段名
    Namespace string `json:"namespace"` // JSON 字段名
    Status    string `json:"status"`    // JSON 字段名
    CPU       string `json:"cpu"`       // JSON 字段名
    Memory    string `json:"memory"`    // JSON 字段名
}

func main() {
    // 场景1：JSON 字符串解析为结构体（类比 Shell 的 jq 解析 JSON）
    fmt.Println("=== JSON 解析 ===")
    jsonStr := `{
        "name": "nginx-01",
        "namespace": "default",
        "status": "Running",
        "cpu": "100m",
        "memory": "256Mi"
    }`

    var pod Pod
    // 反序列化（JSON → 结构体）
    err := json.Unmarshal([]byte(jsonStr), &pod)
    if err != nil {
        fmt.Println("解析 JSON 失败：", err)
        return
    }
    fmt.Printf("解析后的 Pod 信息：\n名称：%s\n命名空间：%s\n状态：%s\n", pod.Name, pod.Namespace, pod.Status)

    // 场景2：结构体生成 JSON 字符串（类比 Shell 的 echo '{"name":"xxx"}'）
    fmt.Println("\n=== JSON 生成 ===")
    newPod := Pod{
        Name:      "redis-01",
        Namespace: "default",
        Status:    "Running",
        CPU:       "200m",
        Memory:    "512Mi",
    }

    // 序列化（结构体 → JSON，Indent 格式化输出）
    jsonData, err := json.MarshalIndent(newPod, "", "  ")
    if err != nil {
        fmt.Println("生成 JSON 失败：", err)
        return
    }
    fmt.Println("生成的 JSON 配置：")
    fmt.Println(string(jsonData))

    // 可选：写入 JSON 文件（运维常用）
    err = os.WriteFile("pod_config.json", jsonData, 0644)
    if err != nil {
        fmt.Println("写入 JSON 文件失败：", err)
    }
}
第 13 天：YAML 解析（gopkg.in/yaml.v3）
任务清单
 安装 yaml 包：go get gopkg.in/yaml.v3
 理解 YAML 序列化 / 反序列化
 练习：读取 / 修改 K8s YAML 配置
示例代码（day13.go）
go
运行
package main

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v3"
)

// 定义结构体（对应 K8s Deployment YAML 核心字段）
type Deployment struct {
    ApiVersion string `yaml:"apiVersion"`
    Kind       string `yaml:"kind"`
    Metadata   struct {
        Name      string `yaml:"name"`
        Namespace string `yaml:"namespace"`
    } `yaml:"metadata"`
    Spec struct {
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

func main() {
    // 1. 读取 K8s YAML 文件（先创建一个 deployment.yaml）
    yamlContent := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deploy
  namespace: default
spec:
  replicas: 2
  template:
    spec:
      containers:
      - name: nginx
        image: nginx:1.21`

    // 写入测试 YAML 文件
    err := os.WriteFile("deployment.yaml", []byte(yamlContent), 0644)
    if err != nil {
        fmt.Println("写入 YAML 文件失败：", err)
        return
    }

    // 读取 YAML 文件内容
    data, err := os.ReadFile("deployment.yaml")
    if err != nil {
        fmt.Println("读取 YAML 文件失败：", err)
        return
    }

    // 2. 解析 YAML 到结构体
    var dep Deployment
    err = yaml.Unmarshal(data, &dep)
    if err != nil {
        fmt.Println("解析 YAML 失败：", err)
        return
    }
    fmt.Printf("解析后的 Deployment：\n名称：%s\n副本数：%d\n镜像：%s\n",
        dep.Metadata.Name, dep.Spec.Replicas, dep.Spec.Template.Spec.Containers[0].Image)

    // 3. 修改 YAML 内容（运维常用：修改镜像、副本数）
    dep.Spec.Replicas = 3                          // 副本数改为 3
    dep.Spec.Template.Spec.Containers[0].Image = "nginx:1.25" // 镜像升级

    // 4. 生成新的 YAML 内容
    newYAML, err := yaml.MarshalIndent(dep, "", "  ")
    if err != nil {
        fmt.Println("生成 YAML 失败：", err)
        return
    }

    // 5. 写入新的 YAML 文件
    err = os.WriteFile("new_deployment.yaml", newYAML, 0644)
    if err != nil {
        fmt.Println("写入新 YAML 文件失败：", err)
        return
    }

    fmt.Println("\n修改后的 YAML 已写入 new_deployment.yaml")
    fmt.Println("新副本数：3，新镜像：nginx:1.25")
}
前置操作
bash
运行
# 安装 yaml 依赖包
go get gopkg.in/yaml.v3
第 14 天（周总结）：运维小工具整合
任务清单
 整合本周知识点：flag + exec + 文件 + YAML
 运行并修改综合工具
 打卡：工具可正常使用
示例代码（day14.go）
go
运行
package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"

    "gopkg.in/yaml.v3"
)

// 定义 Deployment 结构体
type Deployment struct {
    Metadata struct {
        Name string `yaml:"name"`
    } `yaml:"metadata"`
    Spec struct {
        Template struct {
            Spec struct {
                Containers []struct {
                    Image string `yaml:"image"`
                } `yaml:"containers"`
            } `yaml:"spec"`
        } `yaml:"template"`
    } `yaml:"spec"`
}

// 升级 Deployment 镜像
func upgradeDepImage(yamlFile, newImage string) error {
    // 读取 YAML 文件
    data, err := os.ReadFile(yamlFile)
    if err != nil {
        return err
    }

    // 解析 YAML
    var dep Deployment
    err = yaml.Unmarshal(data, &dep)
    if err != nil {
        return err
    }

    // 修改镜像
    dep.Spec.Template.Spec.Containers[0].Image = newImage

    // 生成新 YAML
    newYAML, err := yaml.Marshal(dep)
    if err != nil {
        return err
    }

    // 执行 kubectl apply
    cmd := exec.Command("kubectl", "apply", "-f", "-")
    cmd.Stdin = strings.NewReader(string(newYAML))
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("kubectl apply 失败：%v，输出：%s", err, output)
    }

    fmt.Println("执行结果：", string(output))
    return nil
}

func main() {
    // 定义命令行参数
    yamlFile := flag.String("f", "", "Deployment YAML 文件路径（必填）")
    newImage := flag.String("i", "", "新镜像版本（必填）")
    flag.Parse()

    // 检查参数
    if *yamlFile == "" || *newImage == "" {
        fmt.Println("错误：必须指定 -f YAML 文件 和 -i 新镜像")
        flag.Usage()
        return
    }

    // 执行镜像升级
    fmt.Printf("开始升级 %s 的镜像为 %s...\n", *yamlFile, *newImage)
    err := upgradeDepImage(*yamlFile, *newImage)
    if err != nil {
        fmt.Println("升级失败：", err)
        os.Exit(1)
    }

    fmt.Println("升级成功！")
}
运行测试
bash
运行
# 查看帮助
go run day14.go -h
# 升级镜像（需提前准备 deployment.yaml）
go run day14.go -f deployment.yaml -i nginx:1.25
📅 第 3 周：并发编程（真正编程能力）
第 15 天：goroutine 基础
任务清单
 理解 goroutine（轻量级线程）
 练习：开启多个 goroutine 并发执行任务
 类比 Shell 的 cmd &
示例代码（day15.go）
go
运行
package main

import (
    "fmt"
    "time"
)

// 运维场景：检查单个节点状态（模拟耗时操作）
func checkNode(node string) {
    // 模拟网络请求耗时 1 秒
    time.Sleep(1 * time.Second)
    fmt.Printf("节点 %s 检查完成：状态正常\n", node)
}

func main() {
    nodes := []string{"node-1", "node-2", "node-3", "node-4"}

    fmt.Println("=== 串行检查（耗时久）===")
    // 串行执行（类比 Shell 的 for node in ...; do check $node; done）
    start := time.Now()
    for _, node := range nodes {
        checkNode(node)
    }
    fmt.Printf("串行检查耗时：%v\n", time.Since(start))

    fmt.Println("\n=== 并发检查（goroutine，耗时短）===")
    // 并发执行（类比 Shell 的 for node in ...; do check $node & done）
    start = time.Now()
    for _, node := range nodes {
        go checkNode(node) // 加 go 关键字，开启 goroutine
    }

    // 等待所有 goroutine 完成（临时方案，后续用 WaitGroup）
    time.Sleep(2 * time.Second)
    fmt.Printf("并发检查耗时：%v\n", time.Since(start))
}
第 16 天：sync.WaitGroup 等待组
任务清单
 理解 WaitGroup（等待所有 goroutine 完成）
 练习：用 WaitGroup 替代 Sleep 等待
 类比 Shell 的 wait 命令
示例代码（day16.go）
go
运行
package main

import (
    "fmt"
    "sync"
    "time"
)

// 检查节点状态（接收 WaitGroup 指针）
func checkNode(node string, wg *sync.WaitGroup) {
    defer wg.Done() // 任务完成后，计数器-1（必须放在开头）
    time.Sleep(1 * time.Second)
    fmt.Printf("节点 %s 检查完成：状态正常\n", node)
}

func main() {
    nodes := []string{"node-1", "node-2", "node-3", "node-4"}
    var wg sync.WaitGroup // 定义 WaitGroup

    fmt.Println("=== 并发检查（WaitGroup）===")
    start := time.Now()

    for _, node := range nodes {
        wg.Add(1) // 计数器+1（必须在 goroutine 启动前）
        go checkNode(node, &wg)
    }

    wg.Wait() // 等待所有任务完成（计数器归 0）

    fmt.Printf("并发检查耗时：%v\n", time.Since(start))
    fmt.Println("所有节点检查完成！")
}
第 17 天：channel 通道基础
任务清单
 理解 channel（goroutine 间通信）
 练习：用 channel 接收 goroutine 执行结果
 理解 “同步” 概念
示例代码（day17.go）
go
运行
package main

import (
    "fmt"
    "time"
)

// 检查节点磁盘使用率（结果通过 channel 返回）
func checkDisk(node string, ch chan string) {
    time.Sleep(1 * time.Second)
    // 模拟磁盘使用率
    result := fmt.Sprintf("节点 %s 磁盘使用率：80%%", node)
    ch <- result // 将结果发送到通道
}

func main() {
    nodes := []string{"node-1", "node-2", "node-3"}
    // 创建通道（缓冲大小为 3，避免阻塞）
    ch := make(chan string, len(nodes))

    // 开启 goroutine 执行任务
    for _, node := range nodes {
        go checkDisk(node, ch)
    }

    // 接收通道结果
    fmt.Println("=== 节点磁盘检查结果 ===")
    for i := 0; i < len(nodes); i++ {
        result := <-ch // 从通道接收结果
        fmt.Println(result)
    }

    close(ch) // 关闭通道（可选，防止内存泄漏）
    fmt.Println("所有结果接收完成！")
}
第 18 天：sync.Mutex 锁（简单使用）
任务清单
 理解 Mutex（互斥锁，解决并发写入冲突）
 练习：并发写入日志时加锁
 知道什么时候加锁（共享资源写入）
示例代码（day18.go）
go
运行
package main

import (
    "fmt"
    "os"
    "sync"
    "time"
)

// 全局变量：日志文件句柄、互斥锁
var (
    logFile *os.File
    mutex   sync.Mutex
)

// 初始化日志文件
func initLog() error {
    f, err := os.OpenFile("concurrent.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    logFile = f
    return nil
}

// 并发写入日志（不加锁会导致内容混乱）
func writeLog(content string, wg *sync.WaitGroup) {
    defer wg.Done()

    // 加锁（类比 Shell 的 flock 锁定文件）
    mutex.Lock()
    defer mutex.Unlock() // 函数结束解锁

    // 写入日志
    _, err := logFile.WriteString(fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), content))
    if err != nil {
        fmt.Println("写入日志失败：", err)
        return
    }

    fmt.Println("日志写入成功：", content)
}

func main() {
    // 初始化日志文件
    err := initLog()
    if err != nil {
        fmt.Println("初始化日志失败：", err)
        return
    }
    defer logFile.Close()

    var wg sync.WaitGroup

    // 开启 5 个 goroutine 并发写入日志
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        content := fmt.Sprintf("节点检查任务 %d 完成", i)
        go writeLog(content, &wg)
    }

    wg.Wait()
    fmt.Println("所有日志写入完成！")
}
第 19 天：并发执行命令
任务清单
 整合 goroutine + WaitGroup + exec
 练习：并发执行 kubectl 命令，汇总结果
 运维场景：批量检查 Pod 状态
示例代码（day19.go）
go
运行
package main

import (
    "fmt"
    "os/exec"
    "sync"
)

// 并发执行 kubectl 命令，检查 Pod 状态
func checkPod(podName string, namespace string, wg *sync.WaitGroup, resultChan chan string) {
    defer wg.Done()

    // 构建命令：kubectl get pod <pod> -n <ns> -o jsonpath='{.status.phase}'
    cmd := exec.Command("kubectl", "get", "pod", podName, "-n", namespace, "-o", "jsonpath={.status.phase}")
    output, err := cmd.CombinedOutput()
    if err != nil {
        resultChan <- fmt.Sprintf("Pod %s/%s：检查失败 - %v，输出：%s", namespace, podName, err, output)
        return
    }

    status := string(output)
    resultChan <- fmt.Sprintf("Pod %s/%s：状态 - %s", namespace, podName, status)
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

    // 开启 goroutine 并发检查
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
第 20 天：综合并发小项目
任务清单
 整合 goroutine + WaitGroup + channel + 错误处理
 练习：并发检查多个节点的 CPU / 内存 / 磁盘
 输出清晰的汇总结果
示例代码（day20.go）
go
运行
package main

import (
    "fmt"
    "sync"
    "time"
)

// 节点状态结构体
type NodeStatus struct {
    Node   string // 节点名
    CPU    string // CPU 使用率
    Memory string // 内存使用率
    Disk   string // 磁盘使用率
    Error  error  // 错误信息
}

// 检查节点 CPU 使用率
func checkCPU(node string, status *NodeStatus, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(500 * time.Millisecond)
    // 模拟结果
    status.CPU = fmt.Sprintf("%d%%", 10+int(node[5]-'0')*10)
}

// 检查节点内存使用率
func checkMemory(node string, status *NodeStatus, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(600 * time.Millisecond)
    status.Memory = fmt.Sprintf("%d%%", 20+int(node[5]-'0')*10)
}

// 检查节点磁盘使用率
func checkDisk(node string, status *NodeStatus, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(700 * time.Millisecond)
    // 模拟 node-3 磁盘检查失败
    if node == "node-3" {
        status.Error = fmt.Errorf("磁盘检查超时")
        return
    }
    status.Disk = fmt.Sprintf("%d%%", 30+int(node[5]-'0')*10)
}

// 检查单个节点所有指标
func checkNodeAll(node string, resultChan chan NodeStatus) {
    status := NodeStatus{Node: node}
    var wg sync.WaitGroup

    // 并发检查 CPU/内存/磁盘
    wg.Add(3)
    go checkCPU(node, &status, &wg)
    go checkMemory(node, &status, &wg)
    go checkDisk(node, &status, &wg)

    wg.Wait()
    resultChan <- status
}

func main() {
    nodes := []string{"node-1", "node-2", "node-3", "node-4"}
    resultChan := make(chan NodeStatus, len(nodes))
    var wg sync.WaitGroup

    fmt.Println("开始并发检查节点状态...")
    start := time.Now()

    // 开启 goroutine 检查每个节点
    for _, node := range nodes {
        wg.Add(1)
        go func(n string) {
            defer wg.Done()
            checkNodeAll(n, resultChan)
        }(node)
    }

    wg.Wait()
    close(resultChan)

    // 打印汇总结果
    fmt.Println("\n=== 节点状态巡检汇总 ===")
    fmt.Printf("总耗时：%v\n", time.Since(start))
    for status := range resultChan {
        if status.Error != nil {
            fmt.Printf("节点 %s：检查失败 - %v\n", status.Node, status.Error)
            continue
        }
        fmt.Printf("节点 %s：CPU=%s，内存=%s，磁盘=%s\n",
            status.Node, status.CPU, status.Memory, status.Disk)
    }
}
第 21 天（周总结）：并发能力巩固
任务清单
 复习本周并发知识点
 修改综合项目，添加日志写入
 打卡：并发工具稳定运行
扩展要求
给 day20.go 添加日志写入功能（加 Mutex 锁）
支持命令行参数指定要检查的节点列表
输出结果同时写入文件和控制台
📅 第 4 周：项目实战 + 能力成型
第 22 天：Go Module 完整使用
任务清单
 理解 go mod init/get/tidy 命令
 练习：初始化项目、管理依赖
 规范项目结构
操作步骤 + 示例代码
bash
运行
# 1. 创建项目目录
mkdir k8s-ops-tool
cd k8s-ops-tool

# 2. 初始化 module
go mod init github.com/你的用户名/k8s-ops-tool

# 3. 创建 main.go
touch main.go

# 4. 安装依赖（如 yaml、resty 等）
go get gopkg.in/yaml.v3
go get github.com/go-resty/resty/v2

# 5. 整理依赖（删除未使用的）
go mod tidy
示例代码（main.go）
go
运行
package main

import (
    "fmt"

    "github.com/go-resty/resty/v2"
    "gopkg.in/yaml.v3"
)

func main() {
    // 验证依赖导入
    fmt.Println("=== Go Module 测试 ===")

    // yaml 测试
    data := map[string]string{"name": "k8s-ops-tool", "version": "1.0.0"}
    yamlData, _ := yaml.Marshal(data)
    fmt.Println("YAML 生成：", string(yamlData))

    // resty 测试（HTTP 客户端，运维调用 API 常用）
    client := resty.New()
    resp, err := client.R().Get("https://httpbin.org/get")
    if err != nil {
        fmt.Println("HTTP 请求失败：", err)
        return
    }
    fmt.Println("HTTP 响应：", resp.Status())
}
第 23 天：代码结构：多文件拆分
任务清单
 拆分代码为多文件：main.go/k8sutil.go/logutil.go
 理解包内函数调用
 类比 Shell 多脚本管理
项目结构
plaintext
k8s-ops-tool/
├── go.mod
├── go.sum
├── main.go       # 主程序、命令行参数
├── k8sutil.go    # K8s 相关工具函数
└── logutil.go    # 日志相关工具函数
代码文件
main.go
go
运行
package main

import (
    "flag"
    "fmt"
)

func main() {
    // 初始化日志
    InitLog("k8s-ops.log")

    // 命令行参数
    namespace := flag.String("n", "default", "命名空间")
    flag.Parse()

    // 调用 k8sutil 函数
    pods, err := ListPods(*namespace)
    if err != nil {
        ErrorLog("获取 Pod 列表失败：%v", err)
        return
    }

    InfoLog("成功获取 %s 命名空间下的 Pod 列表：%v", *namespace, pods)
    fmt.Println("Pod 列表：", pods)
}
k8sutil.go
go
运行
package main

import (
    "errors"
    "os/exec"
    "strings"
)

// 列出指定命名空间的 Pod 名称
func ListPods(namespace string) ([]string, error) {
    cmd := exec.Command("kubectl", "get", "pods", "-n", namespace, "-o", "name")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return nil, errors.New(string(output))
    }

    pods := strings.Split(strings.TrimSpace(string(output)), "\n")
    if len(pods) == 0 || pods[0] == "" {
        return nil, errors.New("未找到 Pod")
    }

    // 去掉 pod/ 前缀
    for i, p := range pods {
        pods[i] = strings.TrimPrefix(p, "pod/")
    }

    return pods, nil
}
logutil.go
go
运行
package main

import (
    "fmt"
    "os"
    "sync"
    "time"
)

var (
    logFile *os.File
    mutex   sync.Mutex
)

// 初始化日志文件
func InitLog(filePath string) {
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("初始化日志失败：", err)
        return
    }
    logFile = f
}

// 信息日志
func InfoLog(format string, v ...interface{}) {
    writeLog("INFO", format, v...)
}

// 错误日志
func ErrorLog(format string, v ...interface{}) {
    writeLog("ERROR", format, v...)
}

// 写入日志
func writeLog(level, format string, v ...interface{}) {
    mutex.Lock()
    defer mutex.Unlock()

    // 日志格式：[时间] [级别] 内容
    logLine := fmt.Sprintf("[%s] [%s] %s\n",
        time.Now().Format("2006-01-02 15:04:05"),
        level,
        fmt.Sprintf(format, v...),
    )

    // 打印到控制台
    fmt.Print(logLine)

    // 写入文件
    if logFile != nil {
        _, err := logFile.WriteString(logLine)
        if err != nil {
            fmt.Println("写入日志文件失败：", err)
        }
    }
}
第 24 天：日志输出：log 包
任务清单
 理解标准库 log 包使用
 练习：配置日志输出到文件 + 控制台、加时间戳
 让工具日志更 “专业”
示例代码（day24.go）
go
运行
package main

import (
    "log"
    "os"
)

func main() {
    // 1. 创建日志文件
    logFile, err := os.OpenFile("k8s-tool.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal("创建日志文件失败：", err)
    }
    defer logFile.Close()

    // 2. 配置日志：同时输出到控制台和文件
    multiWriter := io.MultiWriter(os.Stdout, logFile)
    log.SetOutput(multiWriter)

    // 3. 配置日志格式：时间 + 文件名 + 行号 + 内容
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    // 4. 输出不同级别日志（标准库 log 无级别，可自定义）
    log.Println("INFO: 工具启动")
    log.Printf("INFO: 开始处理命名空间 %s\n", "default")

    // 错误日志（Fatal 会退出程序）
    // log.Fatal("ERROR: 致命错误，退出程序")
    log.Println("ERROR: 非致命错误，继续执行")

    // 5. 自定义日志前缀
    log.SetPrefix("[K8S-TOOL] ")
    log.Println("INFO: 工具执行完成")
}
第 25 天：跨平台编译
任务清单
 理解 GOOS/GOARCH 环境变量
 练习：Windows 编译 Linux/macOS 二进制
 编译后的工具能在服务器运行
操作命令 + 示例
bash
运行
# 1. 查看当前系统
go env GOOS GOARCH

# 2. Windows 下编译 Linux 64 位二进制
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o k8s-ops-tool main.go

# 3. Windows 下编译 macOS 64 位二进制
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o k8s-ops-tool-darwin main.go

# 4. Windows 下编译 Windows 二进制
go build -o k8s-ops-tool.exe main.go

# 5. Linux 下编译 Windows 二进制
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o k8s-ops-tool.exe main.go

# 6. 验证二进制（Linux 下）
chmod +x k8s-ops-tool
./k8s-ops-tool -n default
编译脚本（build.bat，Windows 批处理）
bat
@echo off
echo 开始编译跨平台二进制...

:: 编译 Linux 64 位
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o bin/k8s-ops-tool-linux main.go

:: 编译 Windows 64 位
set GOOS=windows
set GOARCH=amd64
go build -o bin/k8s-ops-tool-windows.exe main.go

:: 编译 macOS 64 位
set GOOS=darwin
set GOARCH=amd64
go build -o bin/k8s-ops-tool-darwin main.go

echo 编译完成！二进制文件在 bin 目录下。
pause
第 26 天：开始完整项目（任选其一）
可选项目（选一个即可）
K8s Pod 列表 / 状态查询工具
支持 -n 指定命名空间
支持 -l 指定标签筛选
输出 Pod 名称、状态、IP、重启次数
结果可输出到控制台 / JSON/YAML 文件
批量执行命令工具
支持指定节点列表
并发执行命令
汇总执行结果
错误重试机制
节点信息采集巡检工具
并发采集节点 CPU / 内存 / 磁盘 / 网络
支持阈值告警
生成巡检报告（JSON / 文本）
日志记录
项目初始化步骤
创建项目目录：mkdir k8s-ops-tool && cd k8s-ops-tool
初始化 module：go mod init k8s-ops-tool
创建目录结构：
plaintext
k8s-ops-tool/
├── cmd/          # 命令行子命令
├── internal/     # 内部工具函数
├── pkg/          # 公共包
├── config/       # 配置文件
├── logs/         # 日志目录
└── main.go       # 主程序
编写需求文档（简单版）：列出工具要实现的功能、参数、输出格式
第 27 天：完善项目：参数、日志、错误处理
完善要求
命令行参数：支持 -h/--help、-v/--version、必填参数检查
日志：输出到文件 + 控制台、加时间戳 / 级别、错误日志单独标记
错误处理：所有可能出错的地方都处理（文件读写、命令执行、网络请求）
用户体验：友好的提示信息、进度反馈、颜色输出（可选）
测试：测试各种异常情况（参数错误、Pod 不存在、节点离线）
示例：完善后的 Pod 查询工具（main.go）
go
运行
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
    "strings"

    "gopkg.in/yaml.v3"
)

// 版本号
const version = "1.0.0"

// Pod 信息结构体
type PodInfo struct {
    Name      string `json:"name" yaml:"name"`
    Namespace string `json:"namespace" yaml:"namespace"`
    Status    string `json:"status" yaml:"status"`
    IP        string `json:"ip" yaml:"ip"`
    Restarts  string `json:"restarts" yaml:"restarts"`
}

// 初始化日志
func initLog() {
    // 创建 logs 目录
    if err := os.MkdirAll("logs", 0755); err != nil && !os.IsExist(err) {
        log.Fatal("创建 logs 目录失败：", err)
    }

    // 日志文件
    logFile, err := os.OpenFile("logs/k8s-pod-tool.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal("初始化日志文件失败：", err)
    }
    defer logFile.Close()

    // 同时输出到控制台和文件
    log.SetOutput(io.MultiWriter(os.Stdout, logFile))
    log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// 获取 Pod 详细信息
func getPodInfo(namespace, podName string) (*PodInfo, error) {
    // 获取 Pod 状态
    statusCmd := exec.Command("kubectl", "get", "pod", podName, "-n", namespace, "-o", "jsonpath={.status.phase}")
    statusOut, err := statusCmd.CombinedOutput()
    if err != nil {
        return nil, fmt.Errorf("获取 Pod 状态失败：%v，输出：%s", err, statusOut)
    }

    // 获取 Pod IP
    ipCmd := exec.Command("kubectl", "get", "pod", podName, "-n", namespace, "-o", "jsonpath={.status.podIP}")
    ipOut, err := ipCmd.CombinedOutput()
    if err != nil {
        return nil, fmt.Errorf("获取 Pod IP 失败：%v，输出：%s", err, ipOut)
    }

    // 获取重启次数
    restartCmd := exec.Command("kubectl", "get", "pod", podName, "-n", namespace, "-o", "jsonpath={.status.containerStatuses[0].restartCount}")
    restartOut, err := restartCmd.CombinedOutput()
    if err != nil {
        return nil, fmt.Errorf("获取重启次数失败：%v，输出：%s", err, restartOut)
    }

    return &PodInfo{
        Name:      podName,
        Namespace: namespace,
        Status:    string(statusOut),
        IP:        string(ipOut),
        Restarts:  string(restartOut),
    }, nil
}

// 列出指定命名空间的所有 Pod
func listPods(namespace string) ([]string, error) {
    cmd := exec.Command("kubectl", "get", "pods", "-n", namespace, "-o", "name")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return nil, fmt.Errorf("列出 Pod 失败：%v，输出：%s", err, output)
    }

    pods := strings.Split(strings.TrimSpace(string(output)), "\n")
    for i, p := range pods {
        pods[i] = strings.TrimPrefix(p, "pod/")
    }

    return pods, nil
}

func main() {
    // 初始化日志
    initLog()

    // 定义命令行参数
    namespace := flag.String("n", "default", "指定 K8s 命名空间")
    podName := flag.String("p", "", "指定 Pod 名称（为空则列出所有 Pod）")
    output := flag.String("o", "console", "输出格式：console/json/yaml")
    versionFlag := flag.Bool("v", false, "显示版本号")
    flag.Parse()

    // 显示版本号
    if *versionFlag {
        fmt.Printf("k8s-pod-tool v%s\n", version)
        return
    }

    log.Printf("开始查询 %s 命名空间的 Pod 信息...\n", *namespace)

    // 单个 Pod 查询
    if *podName != "" {
        podInfo, err := getPodInfo(*namespace, *podName)
        if err != nil {
            log.Fatal("查询失败：", err)
        }

        // 输出结果
        switch *output {
        case "json":
            jsonData, _ := json.MarshalIndent(podInfo, "", "  ")
            fmt.Println(string(jsonData))
        case "yaml":
            yamlData, _ := yaml.Marshal(podInfo)
            fmt.Println(string(yamlData))
        default:
            fmt.Printf("=== Pod 信息 ===\n")
            fmt.Printf("名称：%s\n", podInfo.Name)
            fmt.Printf("命名空间：%s\n", podInfo.Namespace)
            fmt.Printf("状态：%s\n", podInfo.Status)
            fmt.Printf("IP：%s\n", podInfo.IP)
            fmt.Printf("重启次数：%s\n", podInfo.Restarts)
        }
        return
    }

    // 列出所有 Pod
    pods, err := listPods(*namespace)
    if err != nil {
        log.Fatal("列出 Pod 失败：", err)
    }

    if len(pods) == 0 {
        log.Println("未找到任何 Pod")
        return
    }

    // 输出所有 Pod
    switch *output {
    case "json":
        jsonData, _ := json.MarshalIndent(pods, "", "  ")
        fmt.Println(string(jsonData))
    case "yaml":
        yamlData, _ := yaml.Marshal(pods)
        fmt.Println(string(yamlData))
    default:
        fmt.Printf("=== %s 命名空间的 Pod 列表（共 %d 个）===\n", *namespace, len(pods))
        for i, pod := range pods {
            fmt.Printf("%d. %s\n", i+1, pod)
        }
    }

    log.Println("查询完成！")
}
第 28 天（最终验收）：完整项目交付
验收标准
 代码结构清晰（多文件拆分）
 支持命令行参数（-h/-v/-n/-o 等）
 完善的日志输出（文件 + 控制台）
 完整的错误处理（无崩溃、友好提示）
 可跨平台编译（Windows/Linux 二进制）
 功能完整（如 Pod 查询工具能列出 / 查询 Pod 信息）
 编写 README.md（使用说明、参数说明、示例）
README.md 示例
markdown
# k8s-pod-tool
K8s Pod 信息查询工具（Go 实现）

## 功能
- 列出指定命名空间的所有 Pod
- 查询单个 Pod 的详细信息（状态、IP、重启次数）
- 支持 console/json/yaml 三种输出格式
- 完善的日志记录

## 编译
```bash
# Windows 编译 Linux 二进制
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o k8s-pod-tool main.go

# Windows 编译 Windows 二进制
go build -o k8s-pod-tool.exe main.go
使用说明
bash
运行
# 显示版本号
./k8s-pod-tool -v

# 列出 default 命名空间的 Pod
./k8s-pod-tool

# 列出 kube-system 命名空间的 Pod
./k8s-pod-tool -n kube-system

# 查询单个 Pod 信息（控制台输出）
./k8s-pod-tool -n default -p nginx-01

# 查询单个 Pod 信息（JSON 输出）
./k8s-pod-tool -n default -p nginx-01 -o json

# 查询单个 Pod 信息（YAML 输出）
./k8s-pod-tool -n default -p nginx-01 -o yaml
日志
日志文件位于 logs/k8s-pod-tool.log
plaintext

---

# ✅ 打卡格式（可直接复制）
日期：____今日任务：第__天完成情况：
 知识点学习
 代码编写
 运行成功
遇到问题：____
解决方法：____
plaintext

---

### 总结
1. 这份清单覆盖 Go 核心编程能力，且完全贴合 K8s 运维场景，代码可直接复制运行；
2. 每天 1 小时、4 周循序渐进，从基础语法到完整项目，最终能独立完成可交付的运维工具；
3. 核心学习逻辑：「复制跑通 → 修改适配 → 独立编写」，符合运维人员从 Shell 过渡到 Go 的学习习惯。