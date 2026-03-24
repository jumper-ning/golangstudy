package tools

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

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