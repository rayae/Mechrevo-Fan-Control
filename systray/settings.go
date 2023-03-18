package systray

import (
	"github.com/bavelee/mfc/cfg"
	"github.com/bavelee/mfc/systray/systray"
	"github.com/bavelee/mfc/util"
	"go.uber.org/zap"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/18 11:10
 * @Desc:
 */

func (m *MFCSystray) bindControlMenus(mfcConfig *cfg.MFCConfig, handlerMap map[*systray.MenuItem]func()) {
	manualControlCheckbox := systray.AddMenuItemCheckbox("手动控制", "", mfcConfig.WorkMode != cfg.WorkModeSystemControl)
	handlerMap[manualControlCheckbox] = func() {
		mfcConfig.Reset()
		if manualControlCheckbox.Checked() {
			zap.S().Info("切换到自动控制模式")
			mfcConfig.WorkMode = cfg.WorkModeSystemControl
			manualControlCheckbox.Uncheck()
			return
		}
		manualControlCheckbox.Check()
		mfcConfig.WorkMode = cfg.WorkModeQuickControl
		id := cfg.PresetQuickModes[1].Id
		mfcConfig.SelectedQuickMode = util.ToStringPtr(id)
		zap.S().Info("切换到手动控制模式", id)
	}

	startupAtBootCheckbox := systray.AddMenuItemCheckbox("开机自启", "", mfcConfig.StartWithWindows)
	handlerMap[startupAtBootCheckbox] = func() {
		checked := startupAtBootCheckbox.Checked()
		if !checked {
			m.Service.Install()
			mfcConfig.StartWithWindows = true
			startupAtBootCheckbox.Check()
		} else {
			m.Service.Uninstall()
			mfcConfig.StartWithWindows = false
			startupAtBootCheckbox.Uncheck()
		}
	}

	if mfcConfig.StartWithWindows {
		m.Service.Install()
	}

	systray.AddSeparator()
}
