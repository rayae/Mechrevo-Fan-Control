package main

import (
	"github.com/bavelee/mfc/action"
	_ "github.com/bavelee/mfc/cfg"
	"github.com/bavelee/mfc/util"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
	"runtime/debug"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/17 19:56
 * @Desc:
 */

func main() {
	util.EnsureRunAsSudo()
	logger := util.SetupGlobalLogger()
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				zap.S().Error(err)
				zap.S().Error(string(debug.Stack()))
			}
		} else {
			_ = logger.Sync()
		}
	}()
	app := &cli.App{
		Name:  "mfc",
		Usage: "机械革命无界 16/16 Pro 风扇控制程序",
		Authors: []*cli.Author{
			{
				Name: "bavelee",
			},
		},
		Action: action.Daemon,
		Commands: []*cli.Command{
			{
				Name:        "start",
				Action:      action.Start,
				Description: "启动服务",
			},
			{
				Name:        "daemon",
				Action:      action.Daemon,
				Description: "启动守护进程",
			},
			{
				Name:        "run",
				Action:      action.RunSysTray,
				Description: "托盘程序",
			},
			{
				Name:        "restart",
				Action:      action.ReStart,
				Description: "重启服务",
			},
			{
				Name:        "stop",
				Action:      action.Stop,
				Description: "停止服务",
			},
			{
				Name:        "install",
				Action:      action.Install,
				Description: "安装系统服务",
			},
			{
				Name:        "uninstall",
				Action:      action.Uninstall,
				Description: "卸载系统服务",
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		zap.S().Error(err)
		os.Exit(1)
	}
	os.Exit(0)
}
