package cfg

/**
 * @Author: bavelee
 * @Date: 2023/3/17 20:44
 * @Desc:
 */

type QuickMode struct {
	Cpu   int    `json:"cpu" yaml:"cpu"`
	Gpu   int    `json:"gpu" yaml:"gpu"`
	Label string `json:"label" yaml:"label"`
	Id    string `json:"id" yaml:"id"`
	//IconData []byte `json:"-" yaml:"-"`
}

var PresetQuickModes = []QuickMode{
	{
		Id:    "silent",
		Label: "安静(30%)",
		Cpu:   30,
		Gpu:   30,
		//IconData: icon.IconSilent,
	},
	{
		Id:    "work",
		Label: "工作(50%)",
		Cpu:   50,
		Gpu:   50,
		//IconData: icon.IconWork,
	},
	{
		Id:    "perf",
		Label: "性能(70%)",
		Cpu:   70,
		Gpu:   70,
		//IconData: icon.IconPerf,
	},
	{
		Id:    "gaming",
		Label: "游戏(85%)",
		Cpu:   85,
		Gpu:   85,
		//IconData: icon.IconGaming,
	},
	{
		Id:    "turbo",
		Label: "狂飙(100%)",
		Cpu:   100,
		Gpu:   100,
		//IconData: icon.IconTurbo,
	},
}

var PresetSpeedValues []int

func init() {
	for i := 5; i <= 100; i += 5 {
		PresetSpeedValues = append(PresetSpeedValues, i)
	}
}

type RefreshInterval struct {
	Name  string
	Value int
}

var PresetRefreshIntervals = []RefreshInterval{
	{
		Name:  "1 秒",
		Value: 1000,
	},
	{
		Name:  "100 毫秒",
		Value: 100,
	},
	{
		Name:  "500 毫秒",
		Value: 500,
	},
	{
		Name:  "3 秒",
		Value: 3000,
	},
	{
		Name:  "5 秒",
		Value: 5000,
	},
	{
		Name:  "10 秒",
		Value: 10000,
	},
}
