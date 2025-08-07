package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/ini.v1"
)

func main() {
	doParse()
}

func isConfExist(fn string) bool {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return false
	}
	return true
}

func doView(fn string) {
	if !isConfExist(fn) {
		fmt.Println("配置文件不存在")
		return
	}

	content, err := os.ReadFile(fn)
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return
	}

	fmt.Println(string(content))
}

func doWGHelp() {
	output, err := exec.Command("curl", "sf/wg").Output()
	if err != nil {
		fmt.Println("获取帮助失败", err)
		return
	}

	fmt.Println(string(output))
}

func doShowStatus() {
	output, err := exec.Command("curl", "sf/wg/show").Output()
	if err != nil {
		fmt.Println("获取状态失败", err)
		return
	}

	fmt.Println(string(output))
}

func doWGDown() {
	output, err := exec.Command("curl", "sf/wg/down").Output()
	if err != nil {
		fmt.Println("关闭 wireguard 失败", err)
		return
	}

	fmt.Println(string(output))
}

func doWGUp(fn string) {
	if !isConfExist(fn) {
		fmt.Println("配置文件不存在")
	}

	cfg, err := ini.Load(fn)
	if err != nil {
		fmt.Println("加载配置文件失败", err)
		return
	}

	interfaceSec := cfg.Section("Interface")
	peerSec := cfg.Section("Peer")

	profileAddress := strings.Split(interfaceSec.Key("Address").String(), ", ")

	doWGDown()
	output, err := exec.Command(
		"curl", "sf/wg/up",
		"-d", fmt.Sprintf("'endpoint=%s'", peerSec.Key("Endpoint").String()),
		"-d", fmt.Sprintf("'PublicKey=%s'", peerSec.Key("PublicKey").String()),
		"-d", fmt.Sprintf("'PrivateKey=%s'", interfaceSec.Key("PrivateKey").String()),
		"-d", fmt.Sprintf("'Address=%s'", profileAddress[0]),
		"-d", fmt.Sprintf("'Addres6=%s'", profileAddress[1]),
		"-d", "'name=sfwg-auto-up'",
	).Output()

	if err != nil {
		fmt.Println("启动 wireguard 失败", err)
		return
	}

	fmt.Println(string(output))
}
