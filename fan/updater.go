package fan

import (
	"github.com/bavelee/mfc/cfg"
	"go.uber.org/zap"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/18 1:44
 * @Desc:
 */

var lastCpuSpeed = 0
var lastGpuSpeed = 0

func Update(mfc *cfg.MFCConfig) {
	if mfc.WorkMode == cfg.WorkModeSystemControl {
		regCtrl.TurnOffFanControl()
		lastGpuSpeed = 0
		lastCpuSpeed = 0
		return
	}

	cpuSpeed := 0
	gpuSpeed := 0
	switch mfc.WorkMode {
	case cfg.WorkModeQuickControl:
		s := *mfc.SelectedQuickMode
		for _, m := range cfg.PresetQuickModes {
			if m.Id == s {
				cpuSpeed = m.Cpu
				gpuSpeed = m.Gpu
				break
			}
		}
	case cfg.WorkModeSpeedControl:
		ci := *mfc.CPUSpeedIndex
		gi := *mfc.GPUSpeedIndex
		cpuSpeed = cfg.PresetSpeedValues[ci]
		gpuSpeed = cfg.PresetSpeedValues[gi]
	}
	regCtrl.TurnOnFanControl()
	if cpuSpeed > 0 {
		if lastCpuSpeed != cpuSpeed {
			zap.S().Info("更新CPU风扇转速", cpuSpeed)
		}
		lastCpuSpeed = cpuSpeed
		regCtrl.SetCPUFanSpeed(cpuSpeed)
	}
	if gpuSpeed > 0 {
		if lastGpuSpeed != gpuSpeed {
			zap.S().Info("更新GPU风扇转速", gpuSpeed)
		}
		lastGpuSpeed = gpuSpeed
		regCtrl.SetGPUFanSpeed(gpuSpeed)
	}
}
