package dll

import (
	_ "embed"
	"github.com/bavelee/mfc/cfg"
	"go.uber.org/zap"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"
)

type Control struct{}

var initialized = false
var wrRdIO *syscall.LazyDLL
var setReg *syscall.LazyProc
var getReg *syscall.LazyProc

//go:embed WrRdIO.dll
var dllByte []byte

func initialize() {
	dllPath := filepath.Join(cfg.Home, "WrRdIO.dll")
	err := os.WriteFile(dllPath, dllByte, fs.ModePerm)
	if err != nil {
		zap.S().Fatalf("无法解压文件: %s 到 %v\n", "WrRdIO.dll", dllPath)
	}
	wrRdIO = syscall.NewLazyDLL(dllPath)
	setReg = wrRdIO.NewProc("_SetRegisterValue@40")
	getReg = wrRdIO.NewProc("_GetRegisterValue@36")
	if wrRdIO == nil || setReg == nil || getReg == nil {
		zap.S().Fatalf("无法加载 WrRdIO.dll")
	}
	initialized = true
}

func invoke(callSet bool, args ...uint32) uintptr {
	if !initialized {
		initialize()
	}
	var params []uintptr
	for _, arg := range args {
		params = append(params, uintptr(arg))
	}
	var proc = getReg
	if callSet {
		proc = setReg
	}
	ret1, _, _ := proc.Call(params...)
	return ret1
}

func (Control) SetRegisterValue(code []uint32) {
	code = append(code, 6, 7, 8, 1, 2)
	invoke(true, code...)
}

func (Control) GetRegisterValue(code []uint32) uintptr {
	code = append(code, 6, 7, 8, 1, 1)
	return invoke(false, code...)
}
