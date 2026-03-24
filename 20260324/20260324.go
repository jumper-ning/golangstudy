// 整合本周知识点：flag + exec + 文件 + YAML
package main

import (
	//"flag"
	"flag"
	"fmt"
	"os"

	// "os/exec"

	"github.com/jumper-ning/golangstudy/tools"
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
				Containers [] struct {
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

	// 解析 YAML 文件
	var dep Deployment
	err = yaml.Unmarshal(data, &dep)
	if err != nil {
		return err
	}

	// 修改镜像
	dep.Spec.Template.Spec.Containers[0].Image = newImage

	// 生成新的 YAML
	newYAML, err := tools.YamlMarshalIndent(dep, 2)
	if err != nil {
		return err
	}

	fmt.Println("新的yaml 文件内容是: \n", string(newYAML))
	return nil
}


func main() {
	// 定义命令行参数
	yamlFile := flag.String("f" , "", "Deployment YAML 文件路径（必填）")
	newImage := flag.String("i", "", "新镜像版本（必填）")
	flag.Parse()

	// 检查参数
	if *yamlFile == "" || *newImage == "" {
		fmt.Println("错误：必须指定 -f YAML 文件 和 -i 新镜像")
		flag.Usage()
		return
	}

	// 执行镜像升级
	fmt.Printf("开始升级 %s 的镜像为 %s ...\n", *yamlFile, *newImage)
	err := upgradeDepImage(*yamlFile, *newImage)
	if err != nil {
		fmt.Println("升级失败: ", err)
		os.Exit(1)
	}

	fmt.Println("升级成功！")
}

