package nettools

import (
	//	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	//	"net/url"
	"stocknew/fortune/db"
	//	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

type HttpClient struct {
	httpClient *http.Client
}

type StockMeter struct {
	GetType      string //查找类别
	ProdCode     string //股票代码
	CandlePeriod string //时间周期
	StartDate    string //开始日期
	EndDate      string //截止日期
	AppCode      string //Auth
}

func (st *StockMeter) GetDateRange(code string) (string, string, error) {
	dc := db.GetDB()
	start, stop, err := dc.GetDateRange(code)
	if err != nil {
		return "error", "error", err
	}
	if stop == "notupdate" {
		return "没有新数据", "没有新数据", errors.New("没有新数据")
	}

	return start, stop, nil
}

func CreateClient() *HttpClient {
	logs.Info("打开一个客户端。")
	httpClient := &HttpClient{}

	tr := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(network, addr, time.Duration(1000)*time.Millisecond)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		//		TLSHandshakeTimeout: time.Duration(1000) * time.Millisecond,
		//		TLSClientConfig: &tls.Config{
		//			InsecureSkipVerify: true,
		//		},
		MaxIdleConnsPerHost:   20,
		ResponseHeaderTimeout: time.Duration(2000) * time.Millisecond}

	httpClient.httpClient = &http.Client{Transport: tr, Timeout: time.Duration(4000) * time.Millisecond}

	return httpClient
}

func (client *HttpClient) HttpDoGet(requrl string) ([]byte, error) {

	req, err := http.NewRequest("GET", requrl, nil)
	if err != nil {
		logs.Error("请求失败1，错误：%v", err)
		return nil, err
	}

	//	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36")
	resp, err := client.httpClient.Do(req)
	if err != nil {
		logs.Error("请求失败2，错误：%v", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("获取接口数据失败，错误：%v", err)
		return nil, err
	}
	logs.Info("body：%v", string(body))
	return body, nil
}
