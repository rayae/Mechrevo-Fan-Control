package systray

import (
	"fmt"
	"github.com/bavelee/mfc/cfg"
	"github.com/bavelee/mfc/systray/systray"
	"go.uber.org/zap"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/18 11:09
 * @Desc:
 */

func (m *MFCSystray) bindRefreshIntervalMenus(mightDisableMenuItems []*systray.MenuItem, handlerMap map[*systray.MenuItem]func(), mfcConfig *cfg.MFCConfig) (func(), []*systray.MenuItem) {
	refreshInterval := systray.AddMenuItem("刷新频率", "频率过快控制的响应速度更快")
	mightDisableMenuItems = append(mightDisableMenuItems, refreshInterval)
	var menuItems []*systray.MenuItem
	for i := range cfg.PresetRefreshIntervals {
		pri := cfg.PresetRefreshIntervals[i]
		index := i
		item := refreshInterval.AddSubMenuItem(pri.Name, "")
		handlerMap[item] = func() {
			zap.S().Info("切换到刷新频率为", index, cfg.PresetRefreshIntervals[index])
			mfcConfig.RefreshInterval = cfg.PresetRefreshIntervals[index].Value
		}
		menuItems = append(menuItems, item)
	}
	return func() {
		interval := mfcConfig.RefreshInterval
		for i, cb := range menuItems {
			if interval == cfg.PresetRefreshIntervals[i].Value {
				cb.Check()
				refreshInterval.SetTitle(fmt.Sprintf("%v(%v)", "刷新频率", cfg.PresetRefreshIntervals[i].Name))
			} else {
				cb.Uncheck()
			}
		}
	}, mightDisableMenuItems
}
