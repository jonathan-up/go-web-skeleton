package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"skeleton/config"
)

func Init() {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败 -> %v\n", err))
	}

	err = yaml.Unmarshal(file, &config.YAML)
	if err != nil {
		panic(fmt.Sprintf("解析配置文件失败 -> %v\n", err))
	}
}
