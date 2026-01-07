package syna

import (
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"net/http"
	"net/url"
	"strings"
)

type SynaQuote struct {
	Id        string `json:"id"`
	FeeAmount struct {
		Type string `json:"type"`
		Hex  string `json:"hex"`
	} `json:"feeAmount"`
	FeeConfig struct {
		BridgeFee int `json:"bridgeFee"`
		MinFee    struct {
			Type string `json:"type"`
			Hex  string `json:"hex"`
		} `json:"minFee"`
		MaxFee struct {
			Type string `json:"type"`
			Hex  string `json:"hex"`
		} `json:"maxFee"`
	} `json:"feeConfig"`
	RouterAddress string `json:"routerAddress"`
	MaxAmountOut  struct {
		Type string `json:"type"`
		Hex  string `json:"hex"`
	} `json:"maxAmountOut"`
	OriginQuery struct {
		RouterAdapter string `json:"routerAdapter"`
		TokenOut      string `json:"tokenOut"`
		MinAmountOut  struct {
			Type string `json:"type"`
			Hex  string `json:"hex"`
		} `json:"minAmountOut"`
		Deadline struct {
			Type string `json:"type"`
			Hex  string `json:"hex"`
		} `json:"deadline"`
		RawParams string `json:"rawParams"`
	} `json:"originQuery"`
	DestQuery struct {
		RouterAdapter string `json:"routerAdapter"`
		TokenOut      string `json:"tokenOut"`
		MinAmountOut  struct {
			Type string `json:"type"`
			Hex  string `json:"hex"`
		} `json:"minAmountOut"`
		Deadline struct {
			Type string `json:"type"`
			Hex  string `json:"hex"`
		} `json:"deadline"`
		RawParams string `json:"rawParams"`
	} `json:"destQuery"`
	EstimatedTime    int    `json:"estimatedTime"`
	BridgeModuleName string `json:"bridgeModuleName"`
	GasDropAmount    struct {
		Type string `json:"type"`
		Hex  string `json:"hex"`
	} `json:"gasDropAmount"`
	OriginChainId      int         `json:"originChainId"`
	DestChainId        int         `json:"destChainId"`
	MaxAmountOutStr    string      `json:"maxAmountOutStr"`
	BridgeFeeFormatted string      `json:"bridgeFeeFormatted"`
	CallData           interface{} `json:"callData"`
}

type Swap struct {
	FromChain string `json:"fromChain"`
	FromToken string `json:"fromToken"`
	ToChain   string `json:"toChain"`
	ToToken   string `json:"toToken"`
	Amount    string `json:"amount"`
}

type BridgeRoute struct {
	Id               string   `json:"id"`
	FromChainId      int      `json:"fromChainId"`
	FromToken        string   `json:"fromToken"`
	FromAmount       string   `json:"fromAmount"`
	ToChainId        int      `json:"toChainId"`
	ToToken          string   `json:"toToken"`
	ExpectedToAmount string   `json:"expectedToAmount"`
	MinToAmount      string   `json:"minToAmount"`
	RouterAddress    string   `json:"routerAddress"`
	EstimatedTime    int      `json:"estimatedTime"`
	ModuleNames      []string `json:"moduleNames"`
	GasDropAmount    string   `json:"gasDropAmount"`
	NativeFee        string   `json:"nativeFee"`
	CallData         struct {
		To    string `json:"to"`
		Data  string `json:"data"`
		Value string `json:"value"`
	} `json:"callData"`
}

// GetSynaBridgeRoute 发送 Syna 报价请求并返回最佳报价
// 参数:
//
//	params - 请求参数映射
//
// 返回:
//
//	SynaQuote - 最佳报价信息
//	error - 错误信息
func GetSynaBridgeRoute(httpClient *http.Client, params map[string]string) (BridgeRoute, error) {
	// 构造完整 URL
	synaHost, synaQuote := config.GetSynaConfig()
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	fullURL := fmt.Sprintf("%s?%s", synaHost+synaQuote, query.Encode())

	// 发起 HTTP GET 请求
	resp, err := httpClient.Get(fullURL)
	if err != nil {
		return BridgeRoute{}, fmt.Errorf("new Syna swap request error %s", err.Error())
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return BridgeRoute{}, fmt.Errorf("syna swap api response status %d, body: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result []BridgeRoute
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return BridgeRoute{}, fmt.Errorf("json decode Syna swap api response error %s", err.Error())
	}

	if len(result) == 0 {
		return BridgeRoute{}, fmt.Errorf("no route")
	}

	// 选择手续费最低的报价
	var bestQuote BridgeRoute
	//var minBridgeFee = 1

	for _, quote := range result {
		if strings.ToLower(quote.RouterAddress) == strings.ToLower("0x512000a034E154908Efb1eC48579F4ffDb000512") {
			bestQuote = quote
			break
		}
		//if quote.FeeConfig.BridgeFee < minBridgeFee {
		//	minBridgeFee = quote.FeeConfig.BridgeFee
		//	bestQuote = quote
		//}
	}
	if bestQuote.Id == "" {
		return BridgeRoute{}, fmt.Errorf("no route")
	}

	return bestQuote, nil
}
