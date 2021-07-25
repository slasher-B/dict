package util

import "mods/src/gui"

//希望持久化的可以根据这里的func实现redis操作
//目前这里的持久化(雾)只针对代理url,项目链接和目录直接保存进.log文件
var ProxyAddr = make(map[int64]string, 60)

func newClient() {
	//return *redis.Client
}

func MapSet(key int64, str string) {
	ProxyAddr[key] = str
}

func mapGet(key int64) string {
	result, ok := ProxyAddr[key]
	if ok {
		return result
	} else {
		gui.LogErr("代理出错", "找不到这个代理")
		return ""
	}
}

func MapDel(key int64) {
	delete(ProxyAddr, key)
}
