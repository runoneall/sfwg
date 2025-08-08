package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

func main() {
	doParse()
}

func isExist(fn string) bool {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return false
	}
	return true
}

func doView(fn string) {
	if !isExist(fn) {
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
	runCmd([]string{"curl", "sf/wg"})
}

func doShowStatus() {
	runCmd([]string{"curl", "sf/wg/show"})
}

func doWGDown() {
	runCmd([]string{"curl", "sf/wg/down"})
}

func doWGUp(fn string) {
	if !isExist(fn) {
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
	profileEndpoint, err := net.ResolveTCPAddr("tcp", peerSec.Key("Endpoint").String())

	if err != nil {
		fmt.Println("解析 endpoint 失败", err)
		return
	}

	runCmd([]string{
		"curl", "sf/wg/up",
		"-d", fmt.Sprintf("endpoint=%s:%d", profileEndpoint.IP.String(), profileEndpoint.Port),
		"-d", fmt.Sprintf("PublicKey=%s", peerSec.Key("PublicKey").String()),
		"-d", fmt.Sprintf("PrivateKey=%s", interfaceSec.Key("PrivateKey").String()),
		"-d", fmt.Sprintf("Address=%s", profileAddress[0]),
		"-d", fmt.Sprintf("Addres6=%s", profileAddress[1]),
		"-d", "name=sfwg",
	})
}

func doGenWGCFProfile() {
	if !isExist("wgcf") {
		fmt.Println("找不到 wgcf 工具, 请置于当前目录下")
		return
	}

	runCmd([]string{"./wgcf", "register", "--accept-tos"})
	runCmd([]string{"./wgcf", "generate"})
}
