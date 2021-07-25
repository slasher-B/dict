package mio

import (
	"mods/src/gui"
	"os"
)

//检查文件是否存在,不存在则新建
func CheckFile(fname string) (f *os.File) {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		f, _ = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	} else {
		f, _ = os.OpenFile(fname, os.O_WRONLY|os.O_APPEND, 0644)
	}
	return f
}

//写入链接
func WriteLinks(s string) {
	f := CheckFile("repoLinks.log")
	defer f.Close()
	if f != nil {
		_, err := f.WriteString(s + "\n")
		if err != nil {
			gui.LogErr("文件写入出错", err)
		}
	}
}

//写入路径
func WriteDict(fname string, s string) {
	f := CheckFile("./Dicts/" + fname + ".log")
	defer f.Close()
	if f != nil {
		_, err := f.WriteString(s)
		if err != nil {
			gui.LogErr("文件写入出错", err)
		}
	}
}
