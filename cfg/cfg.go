package cfg

import (
	"os"
	"path/filepath"
)

var Home string
var LogFile string
var ConfigFile string
var PidFile string
var Exe string

const TaskName = "mfcsvc"
const TaskDesc = "机械革命无界 16 风扇控制服务"

func init() {
	//Home = filepath.Join(os.Getenv("APPDATA"), "mfc")
	Exe, _ = filepath.Abs(os.Args[0])
	Home, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	LogFile = filepath.Join(Home, "mfc.log")
	PidFile = filepath.Join(Home, "pid")
	ConfigFile = filepath.Join(Home, "config.yaml")
	_ = os.MkdirAll(Home, 0777)
}
