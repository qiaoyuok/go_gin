package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go_gin/dal/biz"
	"go_gin/dal/model"
	"go_gin/dal/po"
	"go_gin/utils"
	"net/http"
	url2 "net/url"
	"sync"
	"time"
)

const Timeout = 8
const PageSize = 50
const BdUrl = "https://www.baidu.com/s?wd=%s&cl=3&rn=%d&pn=%d"
const Cookie = "BIDUPSID=C9DC0580F968AB704DA2200168AD7B99; PSTM=1671088294; ZFY=ETrVjx7VAseITBDx8AJPClav7nKeipxfR97iL2TQMlw:C; BDUSS=0std2xyTHdNSm5SelhmTFR6UGpnTnFlVFpJTGtCeVptZ1pqNXNPUWxZVC1WY0pqSUFBQUFBJCQAAAAAAAAAAAEAAABYwawpy-~Hx9Pqb2sAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP7ImmP-yJpjd; BDUSS_BFESS=0std2xyTHdNSm5SelhmTFR6UGpnTnFlVFpJTGtCeVptZ1pqNXNPUWxZVC1WY0pqSUFBQUFBJCQAAAAAAAAAAAEAAABYwawpy-~Hx9Pqb2sAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP7ImmP-yJpjd; BD_HOME=1; H_PS_PSSID=36560_37972_37646_37556_37515_38023_37906_38012_36921_37989_37937_38001_37900_26350_37958_37881; BD_UPN=123253; BA_HECTOR=25ag2ka0012h0kag24858gcv1hr7hq51i; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; delPer=0; BD_CK_SAM=1; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; ISSW=1; ISSW=1; BAIDUID=C9DC0580F968AB702864D132237EB703:SL=1:NR=50:FG=1; BAIDUID_BFESS=C9DC0580F968AB702864D132237EB703:SL=1:NR=50:FG=1; sug=3; sugstore=1; ORIGIN=2; bdime=0; H_PS_645EC=ff77gKj9%2BxqFHFGOvy9zGcJTzrmKKgpHueMb4RA3iPzDV2T5qGmluKVbqCh%2BH3j5pICS; B64_BOT=1; channel=baidusearch; baikeVisitId=4ee13502-dc02-4970-b356-9bc055f7ad95; H_PS_PSSID=36560_37972_37646_37556_37515_38023_37906_38012_36921_37989_37937_38001_37900_26350_37958_37881; PSINO=5; delPer=0; BDSVRTM=22; BD_CK_SAM=1"

// SearchBd 百度搜索
func SearchBd(wg *sync.WaitGroup, kw string, page int) error {
	defer func() {
		wg.Done()
	}()
	client := &http.Client{
		Timeout: time.Second * Timeout,
	}
	targetUrl := fmt.Sprintf(BdUrl, kw, PageSize, (page-1)*PageSize)
	fmt.Println(targetUrl, 111)
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", Cookie)
	req.Header.Add("DNT", "1")
	req.Header.Add("Referer", "https://www.baidu.com/")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	urls := make([]string, 0)
	doc.Find("[class='c-container']").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Find(".c-container h3.c-title a").Html())
		href, exists := s.Find(".c-container h3.c-title a").Attr("href")
		if exists && len(href) > 0 {
			urls = append(urls, href)
		}
	})

	//fmt.Println(urls)

	for _, url := range urls {
		wg.Add(1)
		go GetUrlTKD(wg, url, kw)
	}

	return nil
}

// GetUrlTKD 获取网站TKD
func GetUrlTKD(wg *sync.WaitGroup, url, kw string) (err error) {
	defer func() {
		wg.Done()
	}()
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

	keywords = utils.GetUtf8(keywords)
	description = utils.GetUtf8(description)
	urlParse, _ := url2.Parse(resp.Request.URL.String())

	targetSitePo := po.TargetSitePo{
		Domain:      urlParse.Scheme + "://" + urlParse.Host,
		Url:         resp.Request.URL.String(),
		Title:       title,
		Keywords:    keywords,
		Description: description,
	}

	domain := urlParse.Scheme + "://" + urlParse.Host
	site := &model.KlznSite{
		URL:         urlParse.Scheme + "://" + urlParse.Host,
		TargetURL:   resp.Request.URL.String(),
		Title:       title,
		Keywords:    keywords,
		Description: description,
		Kw:          kw,
		Icon:        domain + "/favicon.ico",
		Status:      1,
	}

	biz.SaveKlznSite(site)
	utils.ZapSugarLogger.Info(targetSitePo)
	return
}
