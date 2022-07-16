package main

import (
	"Mechrevo-16Pro-FanCtrl/cli"
	"Mechrevo-16Pro-FanCtrl/fan"
	"Mechrevo-16Pro-FanCtrl/fan/dll"
	"fmt"
	"os"
	"strconv"
	"time"
)

const VERSION = "v0.0.1-20220716"

var regCtrl fan.RegCtrl

func init() {
	regCtrl = fan.RegCtrl{RegisterControl: dll.Control{}}
}

var (
	Hung = cli.Arg{
		Name:  "--hung",
		Alias: "",
		Desc:  "涡轮加速并在程序退出时关闭(默认)",
		More:  false,
		Value: "",
	}
	Turbo = cli.Arg{
		Name:  "--turbo",
		Alias: "-t",
		Desc:  "左右风扇速度拉满",
		More:  false,
		Value: "",
	}
	Stop = cli.Arg{
		Name:  "--stop",
		Alias: "-s",
		Desc:  "关闭手动风扇控制",
		More:  false,
		Value: "",
	}
	Left = cli.Arg{
		Name:  "--left",
		Alias: "-l",
		Desc:  "设置左侧风扇转速",
		More:  true,
		Value: "40",
	}
	Right = cli.Arg{
		Name:  "--right",
		Alias: "-r",
		Desc:  "设置右侧风扇转速",
		More:  true,
		Value: "40",
	}
	Help = cli.Arg{
		Name:  "--help",
		Alias: "-h",
		Desc:  "显示帮助菜单",
		More:  false,
		Value: "",
	}
)

func main() {
	app := cli.App{
		Args:  &[]*cli.Arg{&Turbo, &Hung, &Stop, &Left, &Right, &Help},
		Title: fmt.Sprintf("fan-control : 适用于机械革命无界14/16 风扇控制(%s)\n\n参数 : \n", VERSION),
		Examples: []string{
			fmt.Sprintf("\t左右转速拉满 : fan-control --turbo\n"),
			fmt.Sprintf("\t关闭风扇控制 : fan-control --stop\n"),
			fmt.Sprintf("\t设置左侧风扇转速30%%,右风扇45%% : fan-control --left 30 --right 45\n"),
		},
	}.Build()
	err := app.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		app.ShowUsage()
		return
	}
	if Stop.Value != "" {
		regCtrl.TurnOffFanControl()
		return
	}
	if Help.Value != "" {
		app.ShowUsage()
		return
	}
	var leftRate, rightRate = 40, 40
	if Left.Value != "" {
		leftRate, err = strconv.Atoi(Left.Value)
		if err != nil {
			fmt.Printf("左侧风扇设置有误\n%v\n", err)
			return
		}
	}
	if Right.Value != "" {
		rightRate, err = strconv.Atoi(Right.Value)
		if err != nil {
			fmt.Printf("右侧风扇设置有误\n%v\n", err)
			return
		}
	}
	var turbo = Turbo.Value != ""
	var hungOn = !(Left.Present || Right.Present || turbo)
	if turbo || hungOn {
		leftRate = 100
		rightRate = 100
	}

	//启动风扇控制
	regCtrl.TurnOnFanControl()
	time.Sleep(time.Duration(50) * time.Millisecond)
	regCtrl.SetLeftFan(leftRate)
	time.Sleep(time.Duration(10) * time.Millisecond)
	regCtrl.SetRightFan(rightRate)
	if hungOn {
		fmt.Printf("芜湖，持续涡轮增压中...\n")
		fmt.Printf("关闭窗口或Ctrl-C将会终端涡轮增速模式...\n")
		regCtrl.SetConsoleCtrlHandler()
		select {}
	}
}
