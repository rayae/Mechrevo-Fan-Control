package fan

/**
 * @Author: Bave Lee
 * @Date: 2022/7/16 8:54
 * @Desc:
 */

func BuildForSetFanSpeed(fan Operation, speed uint32) []uint32 {
	if speed < fan.MinDC {
		speed = fan.MinDC
	} else if speed > fan.MaxDC {
		speed = fan.MaxDC
	}
	return []uint32{
		0x6C,
		0x68,
		fan.Offset,
		fan.Address,
		speed,
	}
}

func BuildForGetValue(fan Operation) []uint32 {
	return []uint32{
		0x6C,
		0x68,
		fan.Offset,
		fan.Address,
	}
}
