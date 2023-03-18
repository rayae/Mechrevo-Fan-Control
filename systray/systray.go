package systray

import (
	"github.com/bavelee/mfc/cfg"
	"github.com/bavelee/mfc/fan"
	"github.com/bavelee/mfc/systray/icon"
	"github.com/bavelee/mfc/systray/systray"
	"github.com/bavelee/mfc/util"
	"github.com/bavelee/mfc/winsvc"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"time"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/17 9:43
 * @Desc:
 */

type MFCSystray struct {
	Service winsvc.ServiceManager
}

func (m *MFCSystray) Start() {
	mfcConfig := cfg.NewMFCConfig(cfg.ConfigFile)
	var mightDisableMenuItems []*systray.MenuItem
	var handlerMap = map[*systray.MenuItem]func(){}

	systray.Run(func() {
		zap.S().Info("开始注入通知区域")
		systray.SetIcon(icon.IconSysTray)
		systray.SetTitle("无界 16 风扇调节")

		// 判断是否开启系统控制
		m.bindControlMenus(mfcConfig, handlerMap)
		m.bindMonitorMenus(mfcConfig)
		updatePresetModeFn, mightDisableMenuItems := m.bindPresetModeMenus(handlerMap, mfcConfig, mightDisableMenuItems)
		updateCpuSpeedFn, mightDisableMenuItems := m.bindCPUSpeedMenus(mightDisableMenuItems, handlerMap, mfcConfig)
		updateGpuSpeedFn, mightDisableMenuItems := m.bindGpuSpeedMenus(mightDisableMenuItems, handlerMap, mfcConfig)
		updateRefreshIntervalFn, mightDisableMenuItems := m.bindRefreshIntervalMenus(mightDisableMenuItems, handlerMap, mfcConfig)
		m.bindMiscMenus(handlerMap)

		var updateHooks = []func(){
			updatePresetModeFn,
			updateCpuSpeedFn,
			updateGpuSpeedFn,
			updateRefreshIntervalFn,
		}

		var updateTrayFn = func() {
			isSystemControl := mfcConfig.WorkMode == cfg.WorkModeSystemControl
			zap.S().Info("更新 SysTray 信息", isSystemControl, util.ToJson(mfcConfig))
			mfcConfig.Save()
			for _, item := range mightDisableMenuItems {
				if isSystemControl {
					item.Disable()
				} else {
					item.Enable()
				}
			}
			if !isSystemControl {
				for _, fn := range updateHooks {
					fn()
				}
			}
		}
		for item, fn := range handlerMap {
			go func(item *systray.MenuItem, fn func()) {
				for {
					select {
					case <-item.ClickedCh:
						fn()
						updateTrayFn()
					}
				}
			}(item, fn)
		}
		updateTrayFn()
		zap.S().Info("通知区域图标已注入")

		go func() {
			for {
				fan.Update(mfcConfig)
				time.Sleep(time.Duration(mfcConfig.RefreshInterval) * time.Millisecond)
			}
		}()

	}, onExit)
}

func (m *MFCSystray) bindMiscMenus(handlerMap map[*systray.MenuItem]func()) {
	logButton := systray.AddMenuItem("日志", "")
	handlerMap[logButton] = func() {
		cmd := exec.Command("explorer", cfg.LogFile)
		_ = cmd.Start()
	}
	exitButton := systray.AddMenuItem("退出", "")
	handlerMap[exitButton] = func() {
		os.Exit(0)
	}
}

func onExit() {
	zap.S().Infof("Systray 已退出")
}
