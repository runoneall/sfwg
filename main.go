package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	doParse()
}

func doView(fn string) {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		fmt.Println("配置文件不存在")
	} else {
		content, err := os.ReadFile(fn)
		if err != nil {
			fmt.Println("读取配置文件失败", err)
		}

		fmt.Println(string(content))
	}
}

func doWGHelp() {
	output, err := exec.Command("curl", "sf/wg").Output()
	if err != nil {
		fmt.Println("获取帮助失败", err)
	} else {
		fmt.Println(string(output))
	}
}

func doShowStatus() {
	output, err := exec.Command("curl", "sf/wg/show").Output()
	if err != nil {
		fmt.Println("获取状态失败", err)
	} else {
		fmt.Println(string(output))
	}
}
