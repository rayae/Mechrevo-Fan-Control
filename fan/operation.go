package fan

type Operation struct {
	Offset      uint32
	Address     uint32
	Name        string
	FanName     string
	ChineseName string
	MaxDC       uint32
	MinDC       uint32
}

const (
	OFF = 0
	ON  = 1
	MAX = 127
	MIN = 1
)

var AutoControl = Operation{
	Offset:      0x91,
	Address:     0x30,
	Name:        "Auto Control",
	FanName:     "Fan Control",
	ChineseName: "自动控制开关",
	MaxDC:       ON,
	MinDC:       OFF,
}

var CPUFanControl = Operation{
	Offset:      0x91,
	Address:     0x33,
	Name:        "CPU Fan",
	FanName:     "CPU Fan #1",
	ChineseName: "左侧风扇",
	MaxDC:       MAX,
	MinDC:       MIN,
}

var GPUFanControl = Operation{
	Offset:      0x91,
	Address:     0x6F,
	Name:        "GPU Fan",
	FanName:     "GPU Fan #1",
	ChineseName: "右侧风扇",
	MaxDC:       MAX,
	MinDC:       MIN,
}

var GetCpuTemperature = Operation{
	Offset:      0x91,
	Address:     0x18,
	Name:        "Get CPU Temperature",
	ChineseName: "CPU 温度",
}

var GetEnvTemperature = Operation{
	Offset:      0x91,
	Address:     0x15,
	Name:        "Get Env Temperature",
	ChineseName: "环境 温度",
}
var GetDc = Operation{
	Offset:      0x18,
	Address:     0x1803,
	Name:        "Get DC",
	ChineseName: "DC",
}
var GetCPUFanSpeed = Operation{
	Offset:      0x18,
	Address:     0x1821,
	Name:        "Get CPU Fan Speed",
	ChineseName: "左侧风扇速度",
}
var GetGPUFanSpeed = Operation{
	Offset:      0x18,
	Address:     0x1820,
	Name:        "Get GPU Fan Speed",
	ChineseName: "右侧风扇速度",
}

/*
		//CPU temp
	ret = t('l', 'h', 0x91, 0x18, 6, 7, 8, 1, 1);
	printf_s(" %d ", ret);
	//env temp
	ret = t('l', 'h', 0x91, 0x15, 6, 7, 8, 1, 1);
	printf_s(" %d ", ret);
	//DC
	ret = t('l', 'h', 0x18, 0x1803, 6, 7, 8, 1, 1);
	printf_s(" %d\n", ret);
	//不知道怎么用但应该和转速有关的两个
	ret = t('l', 'h', 0x18, 0x1821, 6, 7, 8, 1, 1);
	ret = t('l', 'h', 0x18, 0x1820, 6, 7, 8, 1, 1);
	Sleep(300);
*/
