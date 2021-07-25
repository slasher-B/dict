package mio

import (
	"bytes"
	"fmt"
	"net/url"
)

//获取用户输入，拼成url
func GetInput() string {
	//https://github.com/search?q=in%3Aname+[target_name]+stars%3A%3E[star_count]+[...]+&type=
	u := "https://github.com/search?q="
	uri := ""
	var buffer bytes.Buffer
	for {
		fmt.Scan(&uri)
		if uri == "y" {
			break
		}
		buffer.WriteString(uri)
		buffer.WriteString(" ")
	}
	return u + url.QueryEscape(buffer.String()) + "&type="
}
