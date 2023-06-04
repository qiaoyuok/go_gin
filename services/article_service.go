package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go_gin/dal/biz"
	"go_gin/dal/model"
	"go_gin/utils"
	"net/http"
	url2 "net/url"
	"time"
)

func GetTKD(url string) {
	client := &http.Client{
		Timeout: time.Second * Timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	title := utils.GetUtf8(doc.Find("title").Text())
	description, _ := doc.Find("meta[name='description']").Attr("content")
	keywords, _ := doc.Find("meta[name='keywords']").Attr("content")
	content, _ := doc.Find(".main .entry-content").Html()

	keywords = utils.GetUtf8(keywords)
	description = utils.GetUtf8(description)
	urlParse, _ := url2.Parse(resp.Request.URL.String())

	movie := &model.KlznMovie{
		URL:         urlParse.Scheme + "://" + urlParse.Host,
		Title:       title,
		Keyword:     keywords,
		Description: description,
		Content:     content,
		Status:      true,
	}

	//fmt.Println(title, urlParse)
	biz.SaveKlznMovie(movie)
}
