package systray

import (
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

func (m *MFCSystray) bindPresetModeMenus(handlerMap map[*systray.MenuItem]func(), mfcConfig *cfg.MFCConfig, mightDisableMenuItems []*systray.MenuItem) (func(), []*systray.MenuItem) {
	var presetModeCheckbox []*systray.MenuItem
	for i := range cfg.PresetQuickModes {
		qm := cfg.PresetQuickModes[i]
		checkbox := systray.AddMenuItemCheckbox(qm.Label, "", false)
		//checkbox.SetIcon(qm.IconData)
		presetModeCheckbox = append(presetModeCheckbox, checkbox)
		handlerMap[checkbox] = func() {
			zap.S().Info("切换到预设模式", qm)
			mfcConfig.Reset()
			mfcConfig.WorkMode = cfg.WorkModeQuickControl
			mfcConfig.SelectedQuickMode = util.ToStringPtr(qm.Id)
		}
		mightDisableMenuItems = append(mightDisableMenuItems, checkbox)
	}
	systray.AddSeparator()
	return func() {
		// 预设模式
		if mfcConfig.WorkMode == cfg.WorkModeQuickControl {
			m := *mfcConfig.SelectedQuickMode
			for i, cb := range presetModeCheckbox {
				if cfg.PresetQuickModes[i].Id == m {
					cb.Check()
				} else {
					cb.Uncheck()
				}
			}
		} else {
			// 全部反选
			for _, cb := range presetModeCheckbox {
				cb.Uncheck()
			}
		}
	}, mightDisableMenuItems
}
