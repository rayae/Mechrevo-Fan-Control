# Mechrevo-Fan-Control
适用于机械革命14/16的风扇控制小程序

```bash
fan-control : 适用于机械革命无界14/16 风扇控制(v0.0.1-20220716)

参数 :
        -t, --turbo        : 左右风扇速度拉满
        --hung             : 涡轮加速并在程序退出时关闭(默认)
        -s, --stop         : 关闭手动风扇控制
        -l, --left VALUE   : 设置左侧风扇转速
        -r, --right VALUE  : 设置右侧风扇转速
        -h, --help         : 显示帮助菜单
示例 :
        左右转速拉满 : fan-control --turbo
        关闭风扇控制 : fan-control --stop
        设置左侧风扇转速30%,右风扇45% : fan-control --left 30 --right 45
```