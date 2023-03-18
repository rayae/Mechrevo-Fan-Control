package systray

import (
	"fmt"
	"github.com/bavelee/mfc/cfg"
	"github.com/bavelee/mfc/systray/systray"
	"github.com/bavelee/mfc/util"
	"go.uber.org/zap"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/18 11:09
 * @Desc:
 */

func (m *MFCSystray) bindCPUSpeedMenus(mightDisableMenuItems []*systray.MenuItem, handlerMap map[*systray.MenuItem]func(), mfcConfig *cfg.MFCConfig) (func(), []*systray.MenuItem) {
	cpuSpeed := systray.AddMenuItem("CPU 转速", "单独调节 CPU 转速")
	mightDisableMenuItems = append(mightDisableMenuItems, cpuSpeed)
	var cpuSpeedMenuItems []*systray.MenuItem
	for i := range cfg.PresetSpeedValues {
		v := cfg.PresetSpeedValues[i]
		index := i
		item := cpuSpeed.AddSubMenuItem(fmt.Sprintf("%v%%", v), "")
		handlerMap[item] = func() {
			pi := util.ToIntPtr(index)
			gi := mfcConfig.GPUSpeedIndex
			if gi == nil {
				gi = util.ToIntPtr(index)
			}
			zap.S().Info("切换到CPU速度为", index, v, *gi)
			mfcConfig.Reset()
			mfcConfig.WorkMode = cfg.WorkModeSpeedControl
			mfcConfig.CPUSpeedIndex = pi
			mfcConfig.GPUSpeedIndex = gi
		}
		cpuSpeedMenuItems = append(cpuSpeedMenuItems, item)
	}
	return func() {
		if mfcConfig.WorkMode == cfg.WorkModeSpeedControl {
			index := *mfcConfig.CPUSpeedIndex
			for i, cb := range cpuSpeedMenuItems {
				if i == index {
					cb.Check()
					cpuSpeed.SetTitle(fmt.Sprintf("%v(%v%%)", "CPU 转速", cfg.PresetSpeedValues[index]))
				} else {
					cb.Uncheck()
				}
			}
		} else {
			// 全部反选
			for _, cb := range cpuSpeedMenuItems {
				cb.Uncheck()
			}
			cpuSpeed.SetTitle("CPU 转速")
		}
	}, mightDisableMenuItems
}

func (m *MFCSystray) bindGpuSpeedMenus(mightDisableMenuItems []*systray.MenuItem, handlerMap map[*systray.MenuItem]func(), mfcConfig *cfg.MFCConfig) (func(), []*systray.MenuItem) {
	gpuSpeed := systray.AddMenuItem("GPU 转速", "单独调节 GPU 转速")
	mightDisableMenuItems = append(mightDisableMenuItems, gpuSpeed)
	var gpuSpeedMenuItems []*systray.MenuItem
	for i := range cfg.PresetSpeedValues {
		v := cfg.PresetSpeedValues[i]
		index := i
		item := gpuSpeed.AddSubMenuItem(fmt.Sprintf("%v%%", v), "")
		handlerMap[item] = func() {
			pi := util.ToIntPtr(index)
			ci := mfcConfig.CPUSpeedIndex
			if ci == nil {
				ci = util.ToIntPtr(index)
			}
			zap.S().Info("切换到GPU速度为", index, v, *ci)
			mfcConfig.Reset()
			mfcConfig.WorkMode = cfg.WorkModeSpeedControl
			mfcConfig.CPUSpeedIndex = ci
			mfcConfig.GPUSpeedIndex = pi
		}
		gpuSpeedMenuItems = append(gpuSpeedMenuItems, item)
	}
	systray.AddSeparator()
	return func() {
		if mfcConfig.WorkMode == cfg.WorkModeSpeedControl {
			index := *mfcConfig.GPUSpeedIndex
			for i, cb := range gpuSpeedMenuItems {
				if i == index {
					cb.Check()
					gpuSpeed.SetTitle(fmt.Sprintf("%v(%v%%)", "GPU 转速", cfg.PresetSpeedValues[index]))
				} else {
					cb.Uncheck()
				}
			}
		} else {
			// 全部反选
			for _, cb := range gpuSpeedMenuItems {
				cb.Uncheck()
			}
			gpuSpeed.SetTitle("GPU 转速")
		}
	}, mightDisableMenuItems
}
