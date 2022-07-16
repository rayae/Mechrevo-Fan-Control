package fan

/**
 * @Author: Bave Lee
 * @Date: 2022/7/16 8:54
 * @Desc:
 */
// 0x6C, 0x68, 0x91, 0x30, 1, 6, 7, 8, 1, 2
type OperationCode struct {
	Addr   uint32
	Offset uint32
	Opcode uint32
	Fan    uint32
	Speed  uint32
}

func CreateOperationCode(fan Fan, speed uint32) OperationCode {
	if speed < fan.MinDC {
		speed = fan.MinDC
	} else if speed > fan.MaxDC {
		speed = fan.MaxDC
	}
	return OperationCode{
		Addr:   0x6C,
		Offset: 0x68,
		Opcode: 0x91,
		Fan:    fan.Address,
		Speed:  speed,
	}
}
