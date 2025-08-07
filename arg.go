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

var isShowStatus = parser.Flag("", "status", &argparse.Options{
	Required: false,
	Help:     "查看 wireguard 状态",
})

func doParse() {
	parser.DisableHelp()
	if err := parser.Parse(os.Args); err != nil {
		fmt.Println(parser.Usage(err))
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
}
