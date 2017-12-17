package nettools

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	httpClient *http.Client
}

func CreateClient() *HttpClient {
	fmt.Println("打开一个客户端。")
	client := &HttpClient{}
	client.httpClient = &http.Client{}
	return client
}

func (client *HttpClient) HttpDo(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle error
	}

	//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	req.Header.Set("Cookie", "name=anny")

	resp, err := client.httpClient.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取接口数据失败，错误：%v", err)
	}

	return body, nil
}
