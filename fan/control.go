package fan

import (
	"github.com/bavelee/mfc/fan/dll"
	"os"
	"os/signal"
	"syscall"
)

type RegisterControl interface {
	SetRegisterValue([]uint32)
	GetRegisterValue([]uint32) uintptr
}

type RegCtrl struct {
	RegisterControl
}

var regCtrl *RegCtrl

func init() {
	regCtrl = &RegCtrl{RegisterControl: dll.Control{}}
}

func GetRegCtrl() *RegCtrl {
	return regCtrl
}

func (regCtrl *RegCtrl) SetConsoleCtrlHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		regCtrl.TurnOffFanControl()
		os.Exit(0)
	}()
}
func (regCtrl *RegCtrl) SetFanControlStatus(on bool) {
	if on {
		regCtrl.SetRegisterValue(BuildForSetFanSpeed(AutoControl, ON))
	} else {
		regCtrl.SetRegisterValue(BuildForSetFanSpeed(AutoControl, OFF))
	}
}

func (regCtrl *RegCtrl) GetFanControlStatus() bool {
	value := regCtrl.GetRegisterValue(BuildForGetValue(AutoControl))
	return int(value) == ON
}

func (regCtrl *RegCtrl) TurnOnFanControl() {
	regCtrl.SetFanControlStatus(true)
}

func (regCtrl *RegCtrl) TurnOffFanControl() {
	regCtrl.SetFanControlStatus(false)
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

func (regCtrl *RegCtrl) SetCPUFanSpeed(rate int) {
	regCtrl.SetRegisterValue(BuildForSetFanSpeed(CPUFanControl, formatSpeed(rate, 127)))
}

func (regCtrl *RegCtrl) SetGPUFanSpeed(rate int) {
	regCtrl.SetRegisterValue(BuildForSetFanSpeed(GPUFanControl, formatSpeed(rate, 127)))
}
func (regCtrl *RegCtrl) GetCpuTemperature() int {
	value := regCtrl.GetRegisterValue(BuildForGetValue(GetCpuTemperature))
	return int(value)
}
func (regCtrl *RegCtrl) GetEnvTemperature() int {
	value := regCtrl.GetRegisterValue(BuildForGetValue(GetEnvTemperature))
	return int(value)
}
func (regCtrl *RegCtrl) GetDC() int {
	value := regCtrl.GetRegisterValue(BuildForGetValue(GetDc))
	return int(value)
}

func (regCtrl *RegCtrl) GetCPUFanSpeed() int {
	value := regCtrl.GetRegisterValue(BuildForGetValue(GetCPUFanSpeed))
	return int(value)
}

func (regCtrl *RegCtrl) GetGPUFanSpeed() int {
	value := regCtrl.GetRegisterValue(BuildForGetValue(GetGPUFanSpeed))
	return int(value)
}
