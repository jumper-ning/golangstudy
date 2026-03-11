// Printf、Println
package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
)

func main() {
	// 使用exec.Command, 通过 sh -c 参数传递完整的命令字符串
	cmd := exec.Command("sh", "-c", "kubectl get pod --no-headers | wc -l ")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("执行命令失败：%v, 输出：%s", err, output)
	}
	podNum := strings.TrimSpace(string(output))
	fmt.Println("output的值是: ", output)
	fmt.Println("podNum的值是: ", podNum)
	fmt.Println("len(podNum)的值是：", len(podNum))
	fmt.Println("output的参数类型是: ", reflect.TypeOf(output))
	fmt.Println("podNum的参数类型是: ", reflect.TypeOf(podNum))
	fmt.Println("len(podNum)的参数类型是：", reflect.TypeOf(len(podNum)))
	fmt.Printf("outputx 的类型是: %T\n", output)
	fmt.Printf("podNum 的类型是: %T\n", podNum)
	fmt.Printf("len(podNum) 的类型是: %T\n", len(podNum))
}
