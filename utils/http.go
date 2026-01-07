package utils

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

// CreateHTTPClient 创建并返回一个配置好的HTTP客户端
// 该客户端禁用了连接复用，设置了无限超时，并使用了cookie jar
func CreateHTTPClient() (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("cookiejar.New error %s", err.Error())
	}

	return &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			DisableKeepAlives:   true,
			MaxIdleConnsPerHost: -1,
			//Proxy:               http.ProxyFromEnvironment,
		},
		Timeout: 0,
	}, nil
}
