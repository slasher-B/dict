package httpClient

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"mods/src/gui"
	"mods/src/mio"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//获取仓库链接列表
func FindList(c *colly.Collector, u string) {
	findList := c.Clone()
	reg := true
	register := 1
	pages := 100
	flag := "n"
	findList.SetProxyFunc(randProxySwitcher)
	findList.Async = false
	findList.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", u)
		gui.LogWarn("正在发送请求", u)
	})
	findList.OnHTML("div.col-12:nth-child(3) > div:nth-child(1)", func(e *colly.HTMLElement) {
		e.DOM.Each(func(_ int, s *goquery.Selection) {
			count := strings.Split(s.Find("div.flex-column > h3:nth-child(1)").Text(), " ")[4]
			if reg { // 为避免翻一次页就打印一遍
				gui.LogInfo("存在仓库数量", count)
				reg = false
			}
			e.ForEach("a[class=v-align-middle]", func(i int, f *colly.HTMLElement) {
				repo := f.Attr("href")
				repoLink := "https://github.com" + repo
				gui.LogInfo("获取项目链接", repoLink)
				mio.WriteLinks(repoLink) // 写入文件
			})
			l, _ := s.Find("a[class=next_page]").Attr("href")
			u = e.Request.AbsoluteURL(l)
			if l != "" {
				if register == pages {
					gui.LogWarn("警告", "已经爬取"+strconv.Itoa(pages)+"页链接，是否继续?(y/n)")
					for {
						fmt.Scan(&flag)
						if flag == "n" {
							return
						} else if flag == "y" {
							pages = pages + 400
							break
						} else {
							gui.LogErr("没有这个选项", "请选择: (y继续/n停止)")
							continue
						}
					}
				}
				time.Sleep(5 * time.Second)
				findList.Visit(u)
				register++
			}
		})
	})
	findList.Visit(u)
}

//抓取目录
//目标：resource/*   WEB-INF/*   WWW/*
func getPath(c *colly.Collector, u string) {
	repoName := strings.Split(u, "/")[4]
	re := regexp.MustCompile(`^[^.]*$`) // 文件夹=o  文件=x
	pathName := regexp.MustCompile(`/resource?|/www?|/web-inf?`)
	findPath := c.Clone()
	findPath.SetProxyFunc(randProxySwitcher)
	findPath.OnRequest(func(r *colly.Request) {
		//TODO:If-None-Match反爬暂时不知道怎么破解,好像能用Referer代替?
		//r.Headers.Set("If-None-Match","/W[crypto string]")
		r.Headers.Set("Referer", u)
		gui.LogWarn("正在请求项目链接", u)
	})
	findPath.OnHTML("a[class=js-navigation-open Link--primary]", func(e *colly.HTMLElement) {
		tarLink := e.Attr("href")
		if re.FindString(tarLink) != "" {
			if pathName.FindString(strings.ToLower(tarLink)) != "" {
				if re.FindString(tarLink) == "" {
					mio.WriteDict(repoName, "\n")
				}
				path := "/" + e.Text
				gui.LogInfo("爬取目标目录", path)
				mio.WriteDict(repoName, path)
				findPath.Visit(e.Request.AbsoluteURL(tarLink))
			}
			findPath.Visit(e.Request.AbsoluteURL(tarLink))
		}
	})
	findPath.Visit(u)
	findPath.Wait()
}
