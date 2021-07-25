package httpClient

import (
	"bufio"
	"crypto/tls"
	"github.com/gocolly/colly"
	"io"
	"math/rand"
	"mods/src/gui"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	ProxyURL = []string{
		"https://www.kuaidaili.com/free/inha/",
	}
	UA = []string{
		// 31个UA
		//Windows IE
		"Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (compatible; MSIE 10.0; AOL 9.7; AOLBuild 4346.13; Windows NT 6.2; WOW64; Trident/7.0; LCJB)",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
		//Edge
		"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.19 Safari/537.36 Edg/91.0.864.11",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4482.0 Safari/537.36 Edg/92.0.874.0",
		//Chrome
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.7113.93 Safari/537.36",
		//Firefox
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0",
		//Opera
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36 OPR/76.0.4017.94",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36 OPR/76.0.4017.94 (Edition utorrent)",
		//Linux Chrome
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.11 (KHTML, like Gecko) Ubuntu/14.04.6 Chrome/81.0.3990.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4464.5 Safari/537.36",
		//Firefox
		"Mozilla/5.0 (X11; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0",
		//Opera
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36 OPR/76.0.4017.94",
		"Opera/9.80 (X11; Linux i686; Ubuntu/14.10) Presto/2.12.388 Version/12.16",
		//Mac Safari
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.2 Safari/605.1.15",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.3538.77 Safari/537.36",
		//chrome
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.3538.77 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4450.0 Safari/537.36",
		//Firefox
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:89.0) Gecko/20100101 Firefox/89.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:88.0) Gecko/20100101 Firefox/88.0",
		//Opera
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36 OPR/76.0.4017.94",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36 OPR/75.0.3969.250",
		//Edge
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36 Edg/90.0.818.51",
		//IOS chrome
		"Mozilla/5.0 (iPod; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.163 Mobile/15E148 Safari/604.1",
		//Safari
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.5 Mobile/15E148 Snapchat/10.77.5.59 (like Safari/604.1)",
		//WindowsPhone Edge
		"Mozilla/5.0 (Windows Phone 10.0; Android 6.0.1; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Mobile Safari/537.36 Mobile Edge/42.0.0.2028",
		//chrome
		"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 650) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.85 Safari/537.36",
	}
	k       int64
	proxies []*url.URL
)

//全局配置
func Conf() *colly.Collector {
	rand.Seed(time.Now().Unix())
	c := colly.NewCollector(
		colly.Async(true),                       // 开启异步请求
		colly.AllowURLRevisit(),                 // 允许重定向
		colly.UserAgent(UA[rand.Intn(len(UA))]), // 随机User-Agent
	)
	_ = c.Limit(&colly.LimitRule{
		RandomDelay: 5 * time.Second, // 发送两个请求之间的随机延时
		Parallelism: 10,              // 请求并发数
	})
	c.WithTransport(&http.Transport{
		DisableKeepAlives:     true,                                  // 关闭keep-alive
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, // 跳过证书认证
		Proxy:                 http.ProxyFromEnvironment,             // 默认跟随系统代理
		ResponseHeaderTimeout: 5 * time.Second,                       // 响应头接收超时=5s
	})
	c.OnError(func(r *colly.Response, err error) {
		gui.LogErr("请求出错", err)
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		r.Headers.Set("Accept-Encoding", "gzip")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
		r.Headers.Set("Host", "github.com")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})
	c.Wait()
	return c
}

//从repoLinks.log中获取项目链接,进行批量访问
//从文件读取链接
func ReadLinks(c *colly.Collector) {
	f, err := os.Open("repoLinks.log")
	if err != nil {
		gui.LogErr("项目链接文件打开出错", err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, e := rd.ReadString('\n')
		if e != nil || err == io.EOF {
			break
		} else if strings.HasPrefix(line, "#") { // 跳过注释
			continue
		}
		getPath(c, line) // 真正的爬取部分
		time.Sleep(3 * time.Second)
	}
}
