//go:build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取 language-models 目录下的所有子目录
	dirs, err := os.ReadDir("language-models")
	if err != nil {
		fmt.Printf("读取 language-models 目录失败: %v\n", err)
		os.Exit(1)
	}

	// 遍历每个子目录
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		langCode := dir.Name()
		dirPath := fmt.Sprintf("language-models/%s", langCode)
		
		// 创建 embed.go 文件
		filePath := fmt.Sprintf("%s/embed.go", dirPath)
		
		// 生成文件内容
		content := fmt.Sprintf(`package %s

import (
	"embed"
	"github.com/pemistahl/lingua-go"
)

//go:embed *.zip
var model embed.FS

func init() {
	lingua.Register("%s", model)
}
`, langCode, langCode)

		// 写入文件
		err = os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Printf("写入文件 %s 失败: %v\n", filePath, err)
			continue
		}
		
		fmt.Printf("成功创建文件: %s\n", filePath)
	}
}
