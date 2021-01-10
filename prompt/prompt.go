package prompt

import "github.com/gookit/color"

// 基础信息展示
func message(color color.Color, v interface{}) {
	color.Light().Printf("%s\n>", v)
}

func Info(v interface{}) {
	message(color.FgLightCyan, v)
}

func Success(v interface{}) {
	message(color.Green, v)
}

func Warning(v ...interface{}) {
	message(color.Yellow, v)
}

func Err(v ...interface{}) {
	message(color.FgLightRed, v)
}
