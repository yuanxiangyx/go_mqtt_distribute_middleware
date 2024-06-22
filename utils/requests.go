package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpConfig struct {
	TimeOut int               `json:"TimeOut"` // 超时时间  处理
	Proxies map[string]string `json:"Proxies"` // 代理
}

// Args 参数结构体
type Args struct {
	Url     string                 `json:"Url"`     // 请求URL     处理
	Data    map[string]interface{} `json:"Data"`    // 请求体数据   处理
	Params  map[string]string      `json:"Params"`  // 请求参数     处理
	Headers map[string]string      `json:"Headers"` // 请求头数据   处理
	Json    map[string]string      `json:"Json"`    // Json类型请求体   处理

	Cookies map[string]string `json:"Cookies"` // Cookies  处理
	TimeOut int               `json:"TimeOut"` // 超时时间  处理
	Proxies map[string]string `json:"Proxies"` // 代理
	//Verify  bool                   `json:"Verify"`  // 安全认证
	//Cert    bool                   `json:"Cert"`    // 证书
}

func RequestProcess(method string, args Args) (response []byte, err error) {

	var data []byte
	var client *http.Client
	var request *http.Request
	var httpConfig HttpConfig

	// 处理请求参数params
	if args.Params != nil {
		args.Url = args.Url + "?"
		for key, val := range args.Params {
			args.Url = fmt.Sprintf("%s%s=%s&", args.Url, key, val)
		}
		// 处理尾缀多一个&
		args.Url = args.Url[:len(args.Url)-1]
	}

	if args.Data != nil || args.Json != nil {
		if args.Data != nil {
			data, _ = json.Marshal(args.Data)
		} else {
			data, _ = json.Marshal(args.Json)
		}
	} else {
		data = nil
	}

	if method == "GET" {
		request, err = http.NewRequest("GET", args.Url, bytes.NewBuffer(data))
	} else if method == "POST" {
		request, err = http.NewRequest("POST", args.Url, bytes.NewBuffer(data))
	} else {
		return []byte(""), errors.New(fmt.Sprintf("%s请求类型不正确", method))
	}

	if err != nil {
		return []byte(""), err
	}

	if args.Headers != nil {
		for headerTitle, headerVal := range args.Headers {
			request.Header.Set(headerTitle, headerVal)
		}
	} else {
		request.Header.Set("Content-Type", "application/json")
	}

	if args.Cookies != nil {
		for key, val := range args.Cookies {
			request.AddCookie(&http.Cookie{Name: key, Value: val})
		}
	}

	if args.TimeOut != 0 {
		httpConfig.TimeOut = args.TimeOut
	} else {
		httpConfig.TimeOut = 3
	}

	// set http proxy
	//if args.Proxies != nil {
	//	httpConfig.Proxies["http"] = args.Proxies["http"]
	//	_, _ = url.Parse(args.Proxies["http"])
	//}

	client = &http.Client{
		Timeout: time.Duration(httpConfig.TimeOut) * time.Second,
	}
	resp, _ := client.Do(request)

	responseContent, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return responseContent, err
}

func Get(args Args) ([]byte, error) {
	response, err := RequestProcess("GET", args)
	if err != nil {
		panic(err.Error())
	}
	return response, nil
}

func Post(args Args) ([]byte, error) {
	response, err := RequestProcess("POST", args)
	if err != nil {
		panic(err.Error())
	}
	return response, nil
}

func Put(args Args) ([]byte, error) {
	response, err := RequestProcess("PUT", args)
	if err != nil {
		panic(err.Error())
	}
	return response, nil
}
