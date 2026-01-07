package stargate

import (
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"net/http"
	"net/url"
)

type Swap struct {
	srcToken     string `json:"srcToken"`
	dstToken     string `json:"dstToken"`
	srcChainKey  string `json:"srcChainKey"`
	dstChainKey  string `json:"dstChainKey"`
	srcAddress   string `json:"srcAddress"`
	dstAddress   string `json:"dstAddress"`
	srcAmount    string `json:"srcAmount"`
	dstAmountMin string `json:"dstAmountMin"`
}

type StargateQuotes struct {
	Quotes []Quotes `json:"quotes"`
}

type Quotes struct {
	Route           string      `json:"route"`
	Error           interface{} `json:"error"`
	SrcAmount       string      `json:"srcAmount"`
	DstAmount       string      `json:"dstAmount"`
	SrcAmountMax    string      `json:"srcAmountMax"`
	DstAmountMin    string      `json:"dstAmountMin"`
	SrcToken        string      `json:"srcToken"`
	DstToken        string      `json:"dstToken"`
	SrcAddress      string      `json:"srcAddress"`
	DstAddress      string      `json:"dstAddress"`
	SrcChainKey     string      `json:"srcChainKey"`
	DstChainKey     string      `json:"dstChainKey"`
	DstNativeAmount string      `json:"dstNativeAmount"`
	Duration        struct {
		Estimated float64 `json:"estimated"`
	} `json:"duration"`
	Fees  []interface{} `json:"fees"`
	Steps []struct {
		Type        string `json:"type"`
		Sender      string `json:"sender"`
		ChainKey    string `json:"chainKey"`
		Transaction struct {
			Data  string `json:"data"`
			To    string `json:"to"`
			Value string `json:"value"`
			From  string `json:"from"`
		} `json:"transaction"`
	} `json:"steps"`
}

// GetStargateQuotes 发送 Stargate 报价请求并返回最佳报价
// 参数:
//
//	params - 请求参数映射
//
// 返回:
//
//	StargateQuote - 最佳报价信息
//	error - 错误信息
func GetStargateQuotes(httpClient *http.Client, params map[string]string) (Quotes, error) {
	// 构造完整 URL
	stargateHost, stargateQuote := config.GetStargateConfig()
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	fullURL := fmt.Sprintf("%s?%s", stargateHost+stargateQuote, query.Encode())

	// 发起 HTTP GET 请求
	resp, err := httpClient.Get(fullURL)
	if err != nil {
		return Quotes{}, fmt.Errorf("new Stargate swap request error %s", err.Error())
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return Quotes{}, fmt.Errorf("stargate swap api response status %d, body: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result StargateQuotes
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Quotes{}, fmt.Errorf("json decode Stargate swap api response error %s", err.Error())
	}

	if len(result.Quotes) == 0 {
		return Quotes{}, fmt.Errorf("no route")
	}

	quotes := result.Quotes[0]
	if len(quotes.Steps) == 0 {
		return Quotes{}, fmt.Errorf("no route")
	}

	return quotes, nil
}
