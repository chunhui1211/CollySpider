package main

import "fmt"
import (
	"github.com/gocolly/colly"
)
func main() {
	c := colly.NewCollector() 

	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		 fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnHTML(".qa-list__title--ironman", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnHTML("#read_more", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155") // Visit 要放最後
}