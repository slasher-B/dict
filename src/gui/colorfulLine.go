package gui

import "fmt"

func Printc(s interface{}, f int, d int) {
	// 前景f 背景b 颜色     |   代码d  意义
	// ---------------------------------------
	// 30    40   黑色    |    0    终端默认设置
	// 31    41   红色    |    1    高亮显示
	// 32    42   绿色    |    4    使用下划线
	// 33    43   黄色    |    5    闪烁
	// 34    44   蓝色    |    7    反白显示
	// 35    45   紫红色  |    8    不可见
	// 36    46   青蓝色  |    9    划掉
	// 37    47   白色    |
	fmt.Printf("%c[%d;;%dm%v%c[0m", 0x1B, f, d, s, 0x1B)
}

func LogInfo(depiction string, s interface{}) {
	fmt.Printf("%c[%d;;%dm[INFO]%s -> %c[0m%c[%d;;%dm%v%c[0m\n", 0x1B, 0, 36, depiction, 0x1B, 0x1B, 1, 37, s, 0x1B)
}
func LogWarn(depiction string, s interface{}) {
	fmt.Printf("%c[%d;;%dm[WARN]%s -> %c[0m%c[%d;;%dm%v%c[0m\n", 0x1B, 0, 33, depiction, 0x1B, 0x1B, 1, 37, s, 0x1B)
}
func LogErr(depiction string, s interface{}) {
	fmt.Printf("%c[%d;;%dm[INFO]%s -> %c[0m%c[%d;;%dm%v%c[0m\n", 0x1B, 1, 31, depiction, 0x1B, 0x1B, 0, 37, s, 0x1B)
}
func Banner() {
	b := " __       __  __       __  _______   _______   ________  ______   __    __ \n" +
		"/  \\     /  |/  |  _  /  |/       \\ /       \\ /        |/      \\ /  |  /  |\n" +
		"$$  \\   /$$ |$$ | / \\ $$ |$$$$$$$  |$$$$$$$  |$$$$$$$$//$$$$$$  |$$ |  $$ |\n" +
		"$$$  \\ /$$$ |$$ |/$  \\$$ |$$ |__$$ |$$ |  $$ |$$ |__   $$ | _$$/ $$ |__$$ |\n" +
		"$$$$  /$$$$ |$$ /$$$  $$ |$$    $$/ $$ |  $$ |$$    |  $$ |/    |$$    $$ |\n" +
		"$$ $$ $$/$$ |$$ $$/$$ $$ |$$$$$$$/  $$ |  $$ |$$$$$/   $$ |$$$$ |$$$$$$$$ |\n" +
		"$$ |$$$/ $$ |$$$$/  $$$$ |$$ |      $$ |__$$ |$$ |     $$ \\__$$ |$$ |  $$ |\n" +
		"$$ | $/  $$ |$$$/    $$$ |$$ |      $$    $$/ $$ |     $$    $$/ $$ |  $$ |\n" +
		"$$/      $$/ $$/      $$/ $$/       $$$$$$$/  $$/       $$$$$$/  $$/   $$/\n" +
		"                                                                   (v1.0)"
	s := "Make Web Path Dictionary From GitHub"
	l := "-------------------------------------------------------------------------"
	fmt.Printf("%c[%d;;%dm%v%c[0m\n", 0x1B, 0, 35, b, 0x1B)
	fmt.Printf("%c[%d;;%dm%v%c[0m\n", 0x1B, 1, 33, s, 0x1B)
	fmt.Printf("%c[%d;;%dm%v%c[0m\n", 0x1B, 0, 35, l, 0x1B)
}
