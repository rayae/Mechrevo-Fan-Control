package util

import (
	"github.com/sqweek/dialog"
	"os"
	"syscall"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/18 11:03
 * @Desc:
 */

var (
	shell32DLL        = syscall.NewLazyDLL("shell32.dll")
	isUserAnAdminProc = shell32DLL.NewProc("IsUserAnAdmin")
)

func EnsureRunAsSudo() {
	if !isUserAdmin() {
		dialog.Message("本程序只能以管理员身份运行.").Title("错误").Error()
		os.Exit(1)
	}
}

func isUserAdmin() bool {
	var isAdmin bool
	ret, _, _ := isUserAnAdminProc.Call()
	if ret != 0 {
		isAdmin = true
	}
	return isAdmin
}
