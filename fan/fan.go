package fan

type Fan struct {
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

var AutoControl = Fan{
	Address:     0x30,
	Name:        "Auto Control",
	FanName:     "Fan Control",
	ChineseName: "自动控制开关",
	MaxDC:       ON,
	MinDC:       OFF,
}

var Left = Fan{
	Address:     0x33,
	Name:        "Left Fan",
	FanName:     "CPU Fan #1",
	ChineseName: "左侧风扇",
	MaxDC:       MAX,
	MinDC:       MIN,
}

var Right = Fan{
	Address:     0x6F,
	Name:        "Right Fan",
	FanName:     "CPU Fan #1",
	ChineseName: "右侧风扇",
	MaxDC:       MAX,
	MinDC:       MIN,
}
