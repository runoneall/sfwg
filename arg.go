package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var parser = argparse.NewParser(
	"sfwg",
	"segfault 上的 wgcf wireguard 管理程序",
)

var isHelp = parser.Flag("h", "help", &argparse.Options{
	Required: false,
	Help:     "显示帮助信息",
})

var configFile = parser.String("", "use", &argparse.Options{
	Required: false,
	Help:     "配置文件路径",
	Default:  "wgcf-profile.conf",
})

var isView = parser.Flag("", "view", &argparse.Options{
	Required: false,
	Help:     "查看配置文件内容",
})

var isWGHelp = parser.Flag("", "wg-help", &argparse.Options{
	Required: false,
	Help:     "查看 wireguard 命令的帮助信息",
})

var isShowStatus = parser.Flag("s", "status", &argparse.Options{
	Required: false,
	Help:     "查看 wireguard 状态",
})

var isWGDown = parser.Flag("d", "down", &argparse.Options{
	Required: false,
	Help:     "关闭 wireguard 连接",
})

var isWGUp = parser.Flag("u", "up", &argparse.Options{
	Required: false,
	Help:     "开启 wireguard 连接",
})

var isGenWGCFProfile = parser.Flag("g", "wgcf-g", &argparse.Options{
	Required: false,
	Help:     "自动调用 wgcf",
})

var isAutoDownloadWGCF = parser.Flag("", "wgcf-d", &argparse.Options{
	Required: false,
	Help:     "自动下载 wgcf",
})

func doParse() {
	parser.DisableHelp()

	if err := parser.Parse(os.Args); err != nil {
		fmt.Println(parser.Usage(err))
	}

	if len(os.Args) == 1 {
		fmt.Println(parser.Usage(nil))
		return
	}

	if *isHelp {
		fmt.Println(parser.Usage(nil))
	}

	if *isView {
		doView(*configFile)
	}

	if *isWGHelp {
		doWGHelp()
	}

	if *isShowStatus {
		doShowStatus()
	}

	if *isWGDown {
		doWGDown()
	}

	if *isWGUp {
		doWGUp(*configFile)
	}

	if *isGenWGCFProfile {
		doGenWGCFProfile()
	}

	if *isAutoDownloadWGCF {
		doAutoDownloadWGCF()
	}
}
