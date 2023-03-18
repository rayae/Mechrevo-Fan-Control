package winsvc

import (
	"fmt"
	"github.com/bavelee/mfc/cfg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/18 13:15
 * @Desc:
 */

type ServiceManager struct {
}

func (s *ServiceManager) Install() {
	userDomain := os.Getenv("USERDOMAIN")
	userName := os.Getenv("USERNAME")

	// 拼接USERDOMAIN和USERNAME，中间用反斜杠分隔
	user := strings.Join([]string{userDomain, userName}, "\\")

	trigger := fmt.Sprintf("%v daemon", cfg.Exe)

	args := []string{
		"/Create", "/TN", cfg.TaskName, "/SC", "ONLOGON", "/RL", "HIGHEST", "/RU", user, "/F", "/TR", trigger,
	}
	err := s.runSchTasks(args)
	if err != nil {
		zap.S().Errorf("创建任务 %v 状态 : %v", cfg.TaskName, err)
		return
	}
	zap.S().Infof("创建任务 %v 成功", cfg.TaskName)
}

func (s *ServiceManager) Uninstall() {
	args := []string{
		"/Delete", "/TN", cfg.TaskName, "/F",
	}
	err := s.runSchTasks(args)
	if err != nil {
		zap.S().Errorf("删除任务 %v 失败 : %v", cfg.TaskName, err)
		return
	}
	zap.S().Infof("删除任务 %v 成功", cfg.TaskName)
}

func (s *ServiceManager) InRunning() bool {
	args := []string{
		"/Query", "/TN", cfg.TaskName,
	}
	err := s.runSchTasks(args)
	if err != nil {
		zap.S().Errorf("检测任务 %v 状态 : %v", cfg.TaskName, err)
	}
	return err == nil
}

func (s *ServiceManager) Start() {
	args := []string{
		"/Run", "/TN", cfg.TaskName,
	}
	err := s.runSchTasks(args)
	if err != nil {
		zap.S().Errorf("执行任务 %v 失败 : %v", cfg.TaskName, err)
		return
	}
	zap.S().Infof("执行任务 %v 成功", cfg.TaskName)
}

func (s *ServiceManager) Stop() {
	args := []string{
		"/End", "/TN", cfg.TaskName,
	}
	err := s.runSchTasks(args)
	if err != nil {
		zap.S().Errorf("停止任务 %v 失败 : %v", cfg.TaskName, err)
		return
	}
	zap.S().Infof("停止任务 %v 成功", cfg.TaskName)
}

func (s *ServiceManager) Restart() {
	s.Stop()
	s.Start()
}

func (s *ServiceManager) runSchTasks(args []string) error {
	cmd := exec.Command("schtasks.exe", args...)
	cmd.Dir = cfg.Home
	zap.S().Infof("执行命令 %v %v", cmd.Args, cmd.Dir)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		zapcore.InfoLevel,
	)
	plainLogger := zap.New(core)
	defer plainLogger.Sync()
	{
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			zap.S().Warnf("无法重定向子进程 STDOUT : %v", err)
			return nil
		}
		go func() {
			stdoutBuf := make([]byte, 4096)
			for {
				n, err := stdoutPipe.Read(stdoutBuf)
				if err != nil {
					if err != io.EOF {
						zap.S().Error("Error reading stdout:", zap.Error(err))
					}
					return
				}
				plainLogger.Info(string(stdoutBuf[:n]))
			}
		}()
	}
	{
		stderrPipe, err := cmd.StderrPipe()
		if err != nil {
			zap.S().Warnf("无法重定向子进程 STDERR: %v", err)
			return nil
		}
		go func() {
			stdoutBuf := make([]byte, 4096)
			for {
				n, err := stderrPipe.Read(stdoutBuf)
				if err != nil {
					if err != io.EOF {
						zap.S().Error("Error reading stdout:", zap.Error(err))
					}
					return
				}
				plainLogger.Error(string(stdoutBuf[:n]))
			}
		}()
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
