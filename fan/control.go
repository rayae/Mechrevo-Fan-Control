package fan

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type RegisterControl interface {
	SetRegisterValue(code OperationCode)
	GetRegisterValue(code OperationCode) interface{}
}

type RegCtrl struct {
	RegisterControl
}

func (regCtrl RegCtrl) SetConsoleCtrlHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		regCtrl.TurnOffFanControl()
		os.Exit(0)
	}()
}
func (regCtrl RegCtrl) SetFanControlStatus(on bool) {
	if on {
		regCtrl.SetRegisterValue(CreateOperationCode(AutoControl, ON))
	} else {
		regCtrl.SetRegisterValue(CreateOperationCode(AutoControl, OFF))
	}
}

func (regCtrl RegCtrl) TurnOnFanControl() {
	regCtrl.SetFanControlStatus(true)
}

func (regCtrl RegCtrl) TurnOffFanControl() {
	regCtrl.SetFanControlStatus(false)
	fmt.Printf("手动风扇控制关闭\n")
}

func formatSpeed(rate, max int) uint32 {
	if rate < 0 {
		rate = 1
	}
	if rate > 100 {
		rate = 100
	}
	return uint32(max * rate / 100)
}

func (regCtrl RegCtrl) SetLeftFan(rate int) {
	regCtrl.SetRegisterValue(CreateOperationCode(Left, formatSpeed(rate, 127)))
	fmt.Printf("左侧风扇转速设置为 : %d%%\n", rate)
}

func (regCtrl RegCtrl) SetRightFan(rate int) {
	regCtrl.SetRegisterValue(CreateOperationCode(Right, formatSpeed(rate, 127)))
	fmt.Printf("右侧风扇转速设置为 : %d%%\n", rate)
}
