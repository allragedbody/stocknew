package nettools

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"time"
)

type HttpClient struct {
	httpClient *http.Client
}

type StockMeter struct {
	CandlePeriod string //时间周期
	DataCount    string //数据条数
	ProdCode     string //股票代码
	AppCode      string //Auth
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
	q.Add("candle_period", sm.CandlePeriod)
	q.Add("data_count", sm.DataCount)
	q.Add("prod_code", sm.ProdCode)
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

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取接口数据失败，错误：%v", err)
		return nil, err
	}

	return body, nil
}
