package main

import (
	"fmt"
	"github.com/opesun/goquery"
)

func main() {
	var url = "http://toutiao.io"

	p, err := goquery.ParseUrl(url)

	if err != nil {
		panic(err)
	}
	// toutiao.io
	title := p.Find("title").Text()

	fmt.Println(title)

	t := p.Find(".title a")
	for i := 0; i < t.Length(); i++ {
		d := t.Eq(i).Text()
		l := t.Eq(i).Attr("href")
		c := p.Find(".summary a").Eq(i).Text()
		fmt.Println(l, d, "||", c)
	}

	// geek.csdn.net
	p, err = goquery.ParseUrl("http://geek.csdn.net/hotest")
	if err != nil {
		panic(err)
	}

	title = p.Find("title").Text()
	t = p.Find("a.title")

	fmt.Println(title)
	for i := 0; i < t.Length(); i++ {
		l := t.Eq(i).Attr("href")
		c := t.Eq(i).Text()
		fmt.Println(l, c)
	}
	// ituring.com
	p, err = goquery.ParseUrl("http://www.ituring.com.cn/")
	if err != nil {
		panic(err)
	}

	title = p.Find("title").Text()
	t = p.Find(".arc-list").Eq(0)
	x, _ := goquery.ParseString(t.Html())

	t = x.Find("dt a")

	fmt.Println(title)

	for i := 0; i < t.Length(); i++ {
		c := t.Eq(i).Text()
		l := t.Eq(i).Attr("href")
		fmt.Printf("http://www.ituring.com.cn/%s %s\n", l, c)
	}
}
