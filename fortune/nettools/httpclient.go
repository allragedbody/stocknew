package nettools

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"stocknew/fortune/db"
	"time"
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
	fmt.Println("打开一个客户端。")
	httpClient := &HttpClient{}

	tr := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(network, addr, time.Duration(1000)*time.Millisecond)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TLSHandshakeTimeout: time.Duration(1000) * time.Millisecond,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConnsPerHost:   100,
		ResponseHeaderTimeout: time.Duration(3000) * time.Millisecond}

	httpClient.httpClient = &http.Client{Transport: tr, Timeout: time.Duration(8000) * time.Millisecond}

	return httpClient
}

func (client *HttpClient) HttpDoGet(url string, sm *StockMeter) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("get_type", sm.GetType)
	q.Add("prod_code", sm.ProdCode)
	q.Add("candle_period", sm.CandlePeriod)
	q.Add("start_date", sm.StartDate)
	q.Add("end_date", sm.EndDate)

	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "APPCODE 27319841797a486cb7b634b2dfef7ecb")
	//	fmt.Printf("%v", q.Encode())
	//	req, err := http.NewRequest("GET", url, nil)
	//	if err != nil {
	//		// handle error
	//	}

	//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	req.Header.Set("Cookie", "name=anny")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		fmt.Println("请求失败，错误：%v", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取接口数据失败，错误：%v", err)
		return nil, err
	}

	return body, nil
}
