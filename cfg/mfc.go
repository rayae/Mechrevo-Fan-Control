package cfg

import (
	_ "embed"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/17 23:31
 * @Desc:
 */

type WorkMode string

const (
	WorkModeSystemControl WorkMode = "system"
	WorkModeQuickControl  WorkMode = "quick"
	WorkModeSpeedControl  WorkMode = "speed"
	WorkModeDynamic       WorkMode = "dynamic"
)

type MFCConfig struct {
	ConfigFile         string                     `yaml:"-" json:"-"`
	StartWithWindows   bool                       `yaml:"start-with-windows" json:"start-with-windows"`
	RefreshInterval    int                        `yaml:"refresh-interval,omitempty"`
	WorkMode           WorkMode                   `yaml:"work-mode,omitempty" json:"work-mode,omitempty"`
	CPUSpeedIndex      *int                       `yaml:"cpu-speed-index,omitempty" json:"cpu-speed-index,omitempty"`
	GPUSpeedIndex      *int                       `yaml:"gpu-speed-index,omitempty" json:"gpu-speed-index,omitempty"`
	SelectedQuickMode  *string                    `yaml:"selected-quick-mode,omitempty" json:"selected-quick-mode,omitempty"`
	SelectedSpeedTable *string                    `yaml:"selected-speed-table,omitempty" json:"selected-speed-table,omitempty"`
	CustomSpeedTables  *map[string][]DynamicSpeed `yaml:"custom-speed-tables,omitempty" json:"custom-speed-tables,omitempty"`
}

type DynamicSpeed struct {
	Temp int `yaml:"temp"`
	Cpu  int `yaml:"cpu"`
	Gpu  int `yaml:"gpu"`
}

func (m *MFCConfig) Reset() {
	m.CPUSpeedIndex = nil
	m.GPUSpeedIndex = nil
	m.SelectedQuickMode = nil
	m.SelectedSpeedTable = nil
}

func (m *MFCConfig) Save() {
	var tmp MFCConfig
	switch m.WorkMode {
	case WorkModeSystemControl:
		break
	case WorkModeDynamic:
		tmp.CustomSpeedTables = m.CustomSpeedTables
		tmp.SelectedSpeedTable = m.SelectedSpeedTable
	case WorkModeQuickControl:
		tmp.SelectedQuickMode = m.SelectedQuickMode
	case WorkModeSpeedControl:
		tmp.CPUSpeedIndex = m.CPUSpeedIndex
		tmp.GPUSpeedIndex = m.GPUSpeedIndex
	default:
		break
	}
	tmp.WorkMode = m.WorkMode
	tmp.RefreshInterval = m.RefreshInterval
	tmp.StartWithWindows = m.StartWithWindows
	bytes, err := yaml.Marshal(tmp)
	if err != nil {
		zap.S().Error("保存配置文件失败", err, tmp)
		return
	}
	err = os.WriteFile(m.ConfigFile, bytes, 0777)
	if err != nil {
		zap.S().Error("写入配置文件失败", err, tmp)
	}
}

//go:embed default-config.yaml
var defaultConfigBytes []byte

func NewMFCConfig(configFile string) *MFCConfig {
	var err error
	var tmp MFCConfig
	tmp.ConfigFile = configFile
	var bytes []byte
	if _, err := os.Stat(configFile); err != nil {
		_ = os.WriteFile(configFile, defaultConfigBytes, 0777)
		bytes = defaultConfigBytes
	} else {
		bytes, err = os.ReadFile(configFile)
		if err != nil {
			zap.S().Error("读取配置文件失败", err, configFile)
			return nil
		}
	}
	err = yaml.Unmarshal(bytes, &tmp)
	if err != nil {
		zap.S().Error("解析配置文件失败", err, configFile)
		return nil
	}
	if tmp.WorkMode == "" {
		tmp.WorkMode = WorkModeSystemControl
	}
	return &tmp
}
