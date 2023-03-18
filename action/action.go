package action

import (
	"fmt"
	"github.com/bavelee/mfc/cfg"
	"github.com/bavelee/mfc/systray"
	"github.com/bavelee/mfc/winsvc"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func Install(context *cli.Context) error {
	//service := winsvc.NewMFCService()
	//_ = service.Stop()
	//_ = service.Uninstall()
	//return service.Install()
	sm := winsvc.ServiceManager{}
	sm.Install()
	return nil
}

func Start(context *cli.Context) error {
	sm := winsvc.ServiceManager{}
	sm.Start()
	return nil
	//return winsvc.NewMFCService().Start()
}

func ReStart(context *cli.Context) error {
	//return winsvc.NewMFCService().Restart()
	sm := winsvc.ServiceManager{}
	sm.Restart()
	return nil
}

func Stop(context *cli.Context) error {
	//return winsvc.NewMFCService().Stop()
	sm := winsvc.ServiceManager{}
	sm.Stop()
	return nil
}

func Daemon(context *cli.Context) error {
	zap.S().Infof("启动 daemon %v", cfg.Exe)
	cmd := exec.Command(cfg.Exe, "run")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: windows.CREATE_NEW_PROCESS_GROUP | windows.DETACHED_PROCESS,
	}
	err := cmd.Start()
	if err != nil {
		zap.S().Infof("启动 daemon 失败 %v : %v", cfg.Exe, err)
		panic(err)
	}
	err = cmd.Process.Release()
	if err != nil {
		zap.S().Infof("启动 daemon 失败 %v : %v", cfg.Exe, err)
		panic(err)
	}
	return nil
}

func RunSysTray(context *cli.Context) error {
	bytes, _ := os.ReadFile(cfg.PidFile)
	if len(bytes) > 0 {
		s := string(bytes)
		zap.S().Infof("重复进程 %v", s)
		pid, err := strconv.Atoi(s)
		if err == nil {
			process, err := os.FindProcess(pid)
			if err != nil {
				zap.S().Infof("后台进程已结束 pid: %v", pid)
			}
			if process != nil {
				err = process.Kill()
				if err != nil {
					zap.S().Errorf("无法杀死重复 pid: %v :%v", pid, err)
					return err
				}
				zap.S().Infof("重复进程 %d 已终止", pid)
			}
		}
	}
	_ = os.WriteFile(cfg.PidFile, []byte(fmt.Sprintf("%v", os.Getpid())), 0777)
	defer func() {
		_ = os.Remove(cfg.PidFile)
	}()
	sm := winsvc.ServiceManager{}
	mfcSystray := systray.MFCSystray{Service: sm}
	mfcSystray.Start()
	return nil
}

func Uninstall(context *cli.Context) error {
	//service := winsvc.NewMFCService()
	//_ = service.Stop()
	//return service.Uninstall()
	sm := winsvc.ServiceManager{}
	sm.Uninstall()
	return nil
}
