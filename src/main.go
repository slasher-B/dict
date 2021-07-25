package main

import (
	"mods/src/gui"
	"mods/src/httpClient"
	"mods/src/mio"
)

func main() {
	gui.Banner()
	gui.Printc("请输入搜索条件:\n", 1, 33)
	gui.Printc("如: ", 1, 33)
	gui.Printc("in:name vue\n", 0, 37)
	u := mio.GetInput()
	c := httpClient.Conf()
	httpClient.RunProxy(c)
	httpClient.FindList(c, u)
	httpClient.ReadLinks(c)
}
