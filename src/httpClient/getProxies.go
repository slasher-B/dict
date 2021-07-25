package httpClient

import (
	"github.com/gocolly/colly"
	"math/rand"
	"mods/src/gui"
	"mods/src/util"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//抓取免费高匿代理
type p struct {
	ip   string
	port string
}

func GetProxy(c *colly.Collector, u string) {
	info := strings.Split(u, "/")
	getProxy := c.Clone()
	getProxy.Async = false
	getProxy.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", info[2])
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})
	getProxy.OnHTML(".table > tbody:nth-child(2)", func(e *colly.HTMLElement) {
		pro := p{}
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			k++
			pro.ip = el.ChildText("td[data-title=IP]")
			pro.port = el.ChildText("td[data-title=PORT]")
			util.MapSet(k, "http://"+pro.ip+":"+pro.port)
		})
	})
	getProxy.Visit(u)
}

//检查代理可用性,周期=10s
func CheckProxy() {
	for {
		i := 0
		for key, u := range util.ProxyAddr {
			time.Sleep(10 * time.Second)
			if key == 0 {
				continue
			}
			pu, e := url.Parse(u)
			if e != nil {
				break
			}
			client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(pu)}}
			if _, err := client.Get("https://www.baidu.com"); err != nil {
				gui.LogWarn("代理失效", u)
				util.MapDel(key)
				proxies[i], _ = url.Parse(u)
				i++
				continue
			}
		}
	}
}

//抓取代理启动函数,爬取页数=y
func RunProxy(c *colly.Collector) {
	gui.Printc("获取代理中...\n", 0, 36)
	for x := 0; x < len(ProxyURL); x++ {
		for y := 1; y <= 2; y++ {
			GetProxy(c, ProxyURL[x]+strconv.Itoa(y))
			time.Sleep(3 * time.Second)
		}
	}
	go CheckProxy()
}

//colly自带的代理轮巡
func randProxySwitcher(_ *http.Request) (*url.URL, error) {
	return proxies[rand.Intn(len(proxies))], nil
}
