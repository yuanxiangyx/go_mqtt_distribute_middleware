package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
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
	Url     string                 `json:"Url"`
	Data    map[string]interface{} `json:"Data"`
	Params  map[string]string      `json:"Params"`
	Headers map[string]string      `json:"Headers"`
	Json    map[string]string      `json:"Json"`

	Cookies map[string]string `json:"Cookies"`
	TimeOut int               `json:"TimeOut"`
	Proxies map[string]string `json:"Proxies"`
	//Verify  bool                   `json:"Verify"`
	//Cert    bool                   `json:"Cert"`
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
	resp, err := client.Do(request)
	if err != nil {
		return []byte(""), err
	}
	responseContent, err := io.ReadAll(resp.Body)

	_ = resp.Body.Close()
	return responseContent, err
}

func Get(args Args) ([]byte, error) {
	response, err := RequestProcess("GET", args)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
	}
	return response, nil
}

func Post(args Args) ([]byte, error) {
	response, err := RequestProcess("POST", args)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
	}
	return response, nil
}
