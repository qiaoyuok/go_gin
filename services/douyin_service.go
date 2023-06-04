package services

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

func GetDYVideoList(userID string, maxCursor, count int64) {
	wg := sync.WaitGroup{}
	//url := "https://www.douyin.com/aweme/v1/web/aweme/post/?sec_user_id=MS4wLjABAAAAUoce16bhp1iv971RZnOk0xcRNoxZ1gYAethlFJZyJhY&count=35&max_cursor=0&aid=1128&version_name=23.5.0&device_platform=android&os_version=2333&X-Bogus="
	url1 := fmt.Sprintf("https://www.douyin.com/aweme/v1/web/aweme/post/?sec_user_id=%s&count=%d&max_cursor=%d&aid=1128&version_name=23.5.0&device_platform=android&os_version=2333&X-Bogus=", userID, count, maxCursor)
	//fmt.Println(url, 11111)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url1, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "www.douyin.com")
	req.Header.Add("dnt", "1")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("Cookie", "passport_csrf_token=d23418ef9cc6f5c3f547ed0aa9c809a2; passport_csrf_token_default=d23418ef9cc6f5c3f547ed0aa9c809a2; ttwid=1%7CazMiUmQ-bBPwS3jjIx4WWy62XSPAKbi192XG_APjhKk%7C1682259320%7C586d9fa787a30db2a3639fb2f5817d2cf09146bbdcc16b89aae03f8467531354; douyin.com; strategyABtestKey=%221685875303.706%22; s_v_web_id=verify_lihamlsv_5te9IScY_hK7C_45lL_BZ7U_j8RO3lkadXAY; __bd_ticket_guard_local_probe=1685875304236; xgplayer_user_id=6912483545; n_mh=2VrOTksOlZexWfq6a-j6xvTeY1GqrwYTDlMjqD2IfR8; sso_uid_tt=fcabca11f39208423ac55ae18a899bca; sso_uid_tt_ss=fcabca11f39208423ac55ae18a899bca; toutiao_sso_user=9e1aa3f4e16d4f75fbe71e8b54d3c2a9; toutiao_sso_user_ss=9e1aa3f4e16d4f75fbe71e8b54d3c2a9; passport_auth_status=9dbeb289e22992dafac474941567adcc%2C; passport_auth_status_ss=9dbeb289e22992dafac474941567adcc%2C; uid_tt=8e37d12497f988f3c4757b9fea99d84e; uid_tt_ss=8e37d12497f988f3c4757b9fea99d84e; sid_tt=4cfc9d7e708beeaeed4bff4f29afe06d; sessionid=4cfc9d7e708beeaeed4bff4f29afe06d; sessionid_ss=4cfc9d7e708beeaeed4bff4f29afe06d; odin_tt=598cff685f279cdd354f1b83f71ee35cc2bf671516427dc0675b4dfbb75d7836e6dad4cff1d9a54b483da28a00601c18; passport_assist_user=CjwZWAro8oX4Jy47UZBBQ9sBFfzp8B2WV_Q3we7mWtBqL-0J2NbCSfY85QAeazb0Y6-du7BG2yr6ibvCrlsaSAo8KKMTt2iKD9dLdprwXIv_FWTz5v0yCyoLcPuEO5dPhBCFHIHjQfhD2Vrpkavg1GNtridDd_yv0IuImRz6EIn6sg0Yia_WVCIBA-vL0Oc%3D; sid_ucp_sso_v1=1.0.0-KDg0OWFhMjE3MjI5ZTdmZGQ0ZjEyOTI0NGE1YWJhNDYyN2UzMTExOTMKHQiz68uc7AEQu9XxowYY7zEgDDDagJjLBTgGQPQHGgJscSIgOWUxYWEzZjRlMTZkNGY3NWZiZTcxZThiNTRkM2MyYTk; ssid_ucp_sso_v1=1.0.0-KDg0OWFhMjE3MjI5ZTdmZGQ0ZjEyOTI0NGE1YWJhNDYyN2UzMTExOTMKHQiz68uc7AEQu9XxowYY7zEgDDDagJjLBTgGQPQHGgJscSIgOWUxYWEzZjRlMTZkNGY3NWZiZTcxZThiNTRkM2MyYTk; VIDEO_FILTER_MEMO_SELECT=%7B%22expireTime%22%3A1686480188626%2C%22type%22%3A1%7D; LOGIN_STATUS=1; store-region=cn-sh; store-region-src=uid; sid_guard=4cfc9d7e708beeaeed4bff4f29afe06d%7C1685875390%7C5183998%7CThu%2C+03-Aug-2023+10%3A43%3A08+GMT; sid_ucp_v1=1.0.0-KDcxNDFlNWY1MGJlNzcxZWY0MTVkOWJiMzhjMDk5ZTY2NDAyNWRlODcKGQiz68uc7AEQvtXxowYY7zEgDDgGQPQHSAQaAmxmIiA0Y2ZjOWQ3ZTcwOGJlZWFlZWQ0YmZmNGYyOWFmZTA2ZA; ssid_ucp_v1=1.0.0-KDcxNDFlNWY1MGJlNzcxZWY0MTVkOWJiMzhjMDk5ZTY2NDAyNWRlODcKGQiz68uc7AEQvtXxowYY7zEgDDgGQPQHSAQaAmxmIiA0Y2ZjOWQ3ZTcwOGJlZWFlZWQ0YmZmNGYyOWFmZTA2ZA; csrf_session_id=126350b0a66b57556b418e2f8b501dbd; bd_ticket_guard_server_data=; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtY2xpZW50LWNlcnQiOiItLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS1cbk1JSUNFekNDQWJxZ0F3SUJBZ0lVWTEzNWlMRStCUGVRZGdqV0YyVnV1NUVsM0ZFd0NnWUlLb1pJemowRUF3SXdcbk1URUxNQWtHQTFVRUJoTUNRMDR4SWpBZ0JnTlZCQU1NR1hScFkydGxkRjluZFdGeVpGOWpZVjlsWTJSellWOHlcbk5UWXdIaGNOTWpNd05qQTBNVEEwTXpBM1doY05Nek13TmpBME1UZzBNekEzV2pBbk1Rc3dDUVlEVlFRR0V3SkRcblRqRVlNQllHQTFVRUF3d1BZbVJmZEdsamEyVjBYMmQxWVhKa01Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERcbkFRY0RRZ0FFRTQyK29uQldVaTEreGVZMS9BdHZVcUZBa0VuSTIzTTNWZE1UUnRsMFlta20yaDBhUFBhamp5R1pcbkg1ZWNvQklkVUYyZVZQQTRCeUN2cko2VjBObGZmYU9CdVRDQnRqQU9CZ05WSFE4QkFmOEVCQU1DQmFBd01RWURcblZSMGxCQ293S0FZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ0JnZ3JCZ0VGQlFjREF3WUlLd1lCQlFVSEF3UXdcbktRWURWUjBPQkNJRUlMSjVMSGMzdTJkQmZQdFl3bVRSTVZobk0rNFF5VWxDMVVoTG9EZ1pnWlB4TUNzR0ExVWRcbkl3UWtNQ0tBSURLbForcU9aRWdTamN4T1RVQjdjeFNiUjIxVGVxVFJnTmQ1bEpkN0lrZURNQmtHQTFVZEVRUVNcbk1CQ0NEbmQzZHk1a2IzVjVhVzR1WTI5dE1Bb0dDQ3FHU000OUJBTUNBMGNBTUVRQ0lEbzR5bXNtYStSOXlnclNcbng2UHFPTjZnaW5xSy9xdlF3cjhhRUl1SmZjdGFBaUFrVXJ2VlRrbkxtZnRMTTlrK01FdDc4UGNVcXI1aGExOEFcbm5NQkN6ZVplVHc9PVxuLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLVxuIn0; d_ticket=ba28431ce7ef26e48d430688d3023f0966722; download_guide=%223%2F20230604%2F0%22; pwa2=%223%7C1%22; msToken=8HmxVbf6DPb9e0lVGlCKn-br4owZX3AaZP0co22p0yXRK0LZtOK_X-7gsK9rG7zaFamPLH3Q8InykfcDq3ZlXElLNpS34kRJGKXlyIUpcz4krHhR6pJ2; FOLLOW_NUMBER_YELLOW_POINT_INFO=%22MS4wLjABAAAAHBL4iEUvIFFjp472f7gxnxdcxZAza69F2KuLFV_SYwo%2F1685894400000%2F0%2F0%2F1685877295849%22; __ac_nonce=0647c758500776612d09; __ac_signature=_02B4Z6wo00f0158sySgAAIDAz1uN0OJifK-fDM2AAIO.wPEVqM5sLcxlHJfzd6yQ0LfWlVSJPXr4E9NWBMvjluFuVKBz0jJSipK73xfIuO2kucelewLR4TUev2kGjf-sw0GrJRlmBcVIz6mp44; FOLLOW_LIVE_POINT_INFO=%22MS4wLjABAAAAHBL4iEUvIFFjp472f7gxnxdcxZAza69F2KuLFV_SYwo%2F1685894400000%2F0%2F1685878152723%2F0%22; home_can_add_dy_2_desktop=%221%22; passport_fe_beating_status=true; tt_scid=d615A3b5UJTzMklW5g0jCgqZfjfV9.VCQsMrx2EQdeju-eJnj5ntv6mblX7BLctM162e; publish_badge_show_info=%221%2C0%2C0%2C1685878204843%22")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))
	num := int(gjson.Get(string(body), "aweme_list.#").Int())
	for i := 0; i < num; i++ {
		desc := gjson.Get(string(body), fmt.Sprintf("aweme_list.%d.desc", i)).String()
		targetUrl := gjson.Get(string(body), fmt.Sprintf("aweme_list.%d.video.play_addr.url_list.1", i)).String()
		//url := gjson.Get(string(body), fmt.Sprintf("aweme_list.%d.video.play_addr.url_list", i)).String()

		go func() {
			wg.Add(1)
			//fmt.Println(videoInfo, url)
			err = DownloadFile(targetUrl, "./downloads/"+desc+".mp4")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(desc, strings.Replace(targetUrl, "&", "\\u0026", -1))
			wg.Done()
		}()
	}

	wg.Wait()
}

// DownloadFile 下载文件并保存到指定路径
func DownloadFile(fileUrl string, outputPath string) error {
	req, err := http.NewRequest(http.MethodGet, fileUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	tr := &http.Transport{TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	}}

	proxyUrl, err := url.Parse("http://127.0.0.1:7890")
	if err == nil { // 使用传入代理
		tr.Proxy = http.ProxyURL(proxyUrl)
	}

	r, err := (&http.Client{Transport: tr}).Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if r != nil {
		defer r.Body.Close()
	}

	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(r.Body, 32*1024)

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
	return nil
}
