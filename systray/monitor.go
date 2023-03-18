package systray

import (
	"fmt"
	"github.com/bavelee/mfc/cfg"
	"github.com/bavelee/mfc/fan"
	"github.com/bavelee/mfc/systray/systray"
	"time"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/18 11:10
 * @Desc:
 */

func (m *MFCSystray) bindMonitorMenus(mfcConfig *cfg.MFCConfig) {
	tCpu := "CPU 温度 : %v℃"
	tGpu := "GPU 温度 : %v℃"
	tEnv := "环境温度 : %v℃"
	cpuTemp := systray.AddMenuItem("CPU 温度 : -1℃", "")
	cpuTemp.Disable()
	gpuTemp := systray.AddMenuItem("GPU 温度 : -1℃", "")
	gpuTemp.Disable()
	envTemp := systray.AddMenuItem("环境 温度 : -1℃", "")
	envTemp.Disable()
	go func() {
		rc := fan.GetRegCtrl()
		for {
			ct := rc.GetCpuTemperature()
			et := rc.GetEnvTemperature()
			cpuTemp.SetTitle(fmt.Sprintf(tCpu, ct))
			gpuTemp.SetTitle(fmt.Sprintf(tGpu, -1))
			envTemp.SetTitle(fmt.Sprintf(tEnv, et))
			time.Sleep(time.Duration(mfcConfig.RefreshInterval) * time.Millisecond)
		}
	}()
	systray.AddSeparator()
}
