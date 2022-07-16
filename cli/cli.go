package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type App struct {
	Args     *[]*Arg
	Title    string
	Examples []string
	usage    string
}

type Arg struct {
	Name    string
	Desc    string
	Alias   string
	More    bool
	Value   string
	Present bool
}

func (app App) Parse(args []string) error {
	l := len(args)
	for i := 1; i < l; i++ {
		val := args[i]
		var arg *Arg = nil
		for _, v := range *app.Args {
			if val == v.Name || val == v.Alias {
				arg = v
				break
			}
		}
		if arg == nil {

			return fmt.Errorf("未知参数 : %s", val)
		}
		arg.Present = true
		if arg.More {
			i++
			if i < l {
				arg.Value = args[i]
			} else {
				return fmt.Errorf("%s 缺少必要VALUE参数", val)
			}
		} else {
			arg.Value = "<cli.arg.used>"
		}
	}
	return nil
}

func (app App) Build() App {
	maxLen := 0
	for _, arg := range *app.Args {
		l := len(arg.Name)
		if arg.More {
			l = l + len(" VALUE")
		}
		if arg.Alias != "" {
			l = l + len(arg.Alias) + 3
		}
		if l > maxLen {
			maxLen = l
		}
	}
	sb := strings.Builder{}
	sb.WriteString(app.Title)
	format := "\t%-" + strconv.Itoa(maxLen) + "v" + " : %v\n"
	for _, arg := range *app.Args {
		s := arg.Name
		if arg.Alias != "" {
			s = arg.Alias + ", " + s
		}
		if arg.More {
			s = s + " VALUE"
		}
		sb.WriteString(fmt.Sprintf(format, s, arg.Desc))
	}
	sb.WriteString(fmt.Sprintf("示例 : \n"))
	for _, exp := range app.Examples {
		sb.WriteString(exp)
	}
	app.usage = sb.String()
	return app
}
func (app App) ShowUsage() {
	fmt.Println(app.usage)
}
