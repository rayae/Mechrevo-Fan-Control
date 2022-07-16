package dll

import (
	"Mechrevo-16Pro-FanCtrl/fan"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"syscall"
)

/**
 * @Author: Bave Lee
 * @Date: 2022/7/16 8:56
 * @Desc:
 */

type Control struct{}

var initialized = false
var wrRdIO *syscall.LazyDLL
var setReg *syscall.LazyProc

//go:embed WrRdIO.dll
var dllByte []byte

func initialize() {
	saveDir, exists := os.LookupEnv("AppData")
	if !exists {
		executable, err := os.Executable()
		if err == nil {
			saveDir = path.Dir(executable)
		} else {
			saveDir = "./"
		}
	}
	dllPath := path.Join(saveDir, "WrRdIO.dll")
	err := ioutil.WriteFile(dllPath, dllByte, fs.ModePerm)
	if err != nil {
		fmt.Printf("无法解压文件: %s 到 %v\n", "WrRdIO.dll", dllPath)
		return
	}
	wrRdIO = syscall.NewLazyDLL(dllPath)
	setReg = wrRdIO.NewProc("_SetRegisterValue@40")
	if wrRdIO == nil || setReg == nil {
		fmt.Printf("无法加载 WrRdIO.dll\n")
		return
	}
	initialized = true
}

func invoke(args ...uint32) {
	if !initialized {
		initialize()
	}
	var params []uintptr
	for _, arg := range args {
		params = append(params, uintptr(arg))
	}
	ret1, ret2, lastErr := setReg.Call(params...)
	if false {
		fmt.Printf("SetRegisterValue ret1=%v ret2=%v lastErr=%v\n", ret1, ret2, lastErr)
	}
}

func (Control) SetRegisterValue(code fan.OperationCode) {
	invoke(code.Addr, code.Offset, code.Opcode, code.Fan, code.Speed, 6, 7, 8, 1, 2)
}

func (Control) GetRegisterValue(code fan.OperationCode) interface{} {
	return nil
}
