package squid

type QuoteRequest struct {
	FromAddress string `json:"fromAddress"`
	FromChain   string `json:"fromChain"`
	FromToken   string `json:"fromToken"`
	FromAmount  string `json:"fromAmount"`
	ToChain     string `json:"toChain"`
	ToToken     string `json:"toToken"`
	ToAddress   string `json:"toAddress"`
}
type QuoteResponse struct {
	Route struct {
		Estimate struct {
			Actions []struct {
				Type      string `json:"type"`
				ChainType string `json:"chainType"`
				Data      struct {
					Address           string   `json:"address,omitempty"`
					ChainId           string   `json:"chainId,omitempty"`
					CoinAddresses     []string `json:"coinAddresses,omitempty"`
					Dex               string   `json:"dex,omitempty"`
					Enabled           bool     `json:"enabled"`
					Path              []string `json:"path,omitempty"`
					Slippage          float64  `json:"slippage,omitempty"`
					AggregateSlippage float64  `json:"aggregateSlippage,omitempty"`
					Target            string   `json:"target,omitempty"`
					Name              string   `json:"name,omitempty"`
					Provider          string   `json:"provider,omitempty"`
					Type              string   `json:"type,omitempty"`
					LogoURI           string   `json:"logoURI,omitempty"`
				} `json:"data"`
				FromChain string `json:"fromChain"`
				ToChain   string `json:"toChain"`
				FromToken struct {
					Id                  string      `json:"id"`
					Symbol              string      `json:"symbol"`
					Address             string      `json:"address"`
					ChainId             string      `json:"chainId"`
					Name                string      `json:"name"`
					Decimals            int         `json:"decimals"`
					CoingeckoId         string      `json:"coingeckoId"`
					Type                string      `json:"type"`
					LogoURI             string      `json:"logoURI"`
					AxelarNetworkSymbol string      `json:"axelarNetworkSymbol"`
					SubGraphOnly        bool        `json:"subGraphOnly"`
					SubGraphIds         []string    `json:"subGraphIds"`
					Enabled             bool        `json:"enabled"`
					Active              bool        `json:"active"`
					Visible             bool        `json:"visible"`
					UsdPrice            float64     `json:"usdPrice"`
					InterchainTokenId   interface{} `json:"interchainTokenId"`
				} `json:"fromToken"`
				ToToken struct {
					Id                  string      `json:"id"`
					Symbol              string      `json:"symbol"`
					Address             string      `json:"address"`
					ChainId             string      `json:"chainId"`
					Name                string      `json:"name"`
					Decimals            int         `json:"decimals"`
					InterchainTokenId   interface{} `json:"interchainTokenId"`
					CoingeckoId         string      `json:"coingeckoId"`
					Type                string      `json:"type"`
					LogoURI             string      `json:"logoURI"`
					AxelarNetworkSymbol string      `json:"axelarNetworkSymbol"`
					SubGraphOnly        bool        `json:"subGraphOnly"`
					SubGraphIds         []string    `json:"subGraphIds"`
					Enabled             bool        `json:"enabled"`
					Active              bool        `json:"active"`
					Visible             bool        `json:"visible"`
					UsdPrice            float64     `json:"usdPrice"`
				} `json:"toToken"`
				AggregatedVolatility int    `json:"aggregatedVolatility"`
				FromAmount           string `json:"fromAmount"`
				ToAmount             string `json:"toAmount"`
				ToAmountMin          string `json:"toAmountMin"`
				ExchangeRate         string `json:"exchangeRate"`
				PriceImpact          string `json:"priceImpact"`
				Stage                int    `json:"stage"`
				Provider             string `json:"provider"`
				LogoURI              string `json:"logoURI"`
				Description          string `json:"description"`
				EstimatedDuration    int    `json:"estimatedDuration,omitempty"`
			} `json:"actions"`
			FromAmount           string  `json:"fromAmount"`
			ToAmount             string  `json:"toAmount"`
			ToAmountMin          string  `json:"toAmountMin"`
			ExchangeRate         string  `json:"exchangeRate"`
			AggregatePriceImpact string  `json:"aggregatePriceImpact"`
			FromAmountUSD        string  `json:"fromAmountUSD"`
			ToAmountUSD          string  `json:"toAmountUSD"`
			ToAmountMinUSD       string  `json:"toAmountMinUSD"`
			AggregateSlippage    float64 `json:"aggregateSlippage"`
			FromToken            struct {
				Id                  string        `json:"id"`
				Symbol              string        `json:"symbol"`
				Address             string        `json:"address"`
				ChainId             string        `json:"chainId"`
				Name                string        `json:"name"`
				Decimals            int           `json:"decimals"`
				CoingeckoId         string        `json:"coingeckoId"`
				Type                string        `json:"type"`
				LogoURI             string        `json:"logoURI"`
				Volatility          int           `json:"volatility"`
				AxelarNetworkSymbol string        `json:"axelarNetworkSymbol"`
				SubGraphOnly        bool          `json:"subGraphOnly"`
				SubGraphIds         []interface{} `json:"subGraphIds"`
				Enabled             bool          `json:"enabled"`
				Active              bool          `json:"active"`
				Visible             bool          `json:"visible"`
				UsdPrice            float64       `json:"usdPrice"`
			} `json:"fromToken"`
			ToToken struct {
				Id                  string   `json:"id"`
				Symbol              string   `json:"symbol"`
				Address             string   `json:"address"`
				ChainId             string   `json:"chainId"`
				Name                string   `json:"name"`
				Decimals            int      `json:"decimals"`
				CoingeckoId         string   `json:"coingeckoId"`
				Type                string   `json:"type"`
				LogoURI             string   `json:"logoURI"`
				Volatility          int      `json:"volatility"`
				AxelarNetworkSymbol string   `json:"axelarNetworkSymbol"`
				SubGraphOnly        bool     `json:"subGraphOnly"`
				SubGraphIds         []string `json:"subGraphIds"`
				Enabled             bool     `json:"enabled"`
				Active              bool     `json:"active"`
				Visible             bool     `json:"visible"`
				UsdPrice            float64  `json:"usdPrice"`
			} `json:"toToken"`
			IsBoostSupported bool `json:"isBoostSupported"`
			FeeCosts         []struct {
				Amount        string  `json:"amount"`
				AmountUsd     string  `json:"amountUsd"`
				Description   string  `json:"description"`
				GasLimit      string  `json:"gasLimit,omitempty"`
				GasMultiplier float64 `json:"gasMultiplier,omitempty"`
				Name          string  `json:"name"`
				Token         struct {
					Id                  string   `json:"id"`
					Symbol              string   `json:"symbol"`
					Address             string   `json:"address"`
					ChainId             string   `json:"chainId"`
					Name                string   `json:"name"`
					Decimals            int      `json:"decimals"`
					CoingeckoId         string   `json:"coingeckoId"`
					Type                string   `json:"type"`
					LogoURI             string   `json:"logoURI"`
					Volatility          int      `json:"volatility"`
					AxelarNetworkSymbol string   `json:"axelarNetworkSymbol"`
					SubGraphOnly        bool     `json:"subGraphOnly"`
					SubGraphIds         []string `json:"subGraphIds"`
					Enabled             bool     `json:"enabled"`
					Active              bool     `json:"active"`
					Visible             bool     `json:"visible"`
					UsdPrice            float64  `json:"usdPrice"`
				} `json:"token"`
				LogoURI string `json:"logoURI"`
				Data    struct {
					AxelarFeeData struct {
						BaseFee           float64 `json:"baseFee"`
						ExpressFee        float64 `json:"expressFee"`
						ExecuteMultiplier float64 `json:"executeMultiplier"`
						ExpressMultiplier float64 `json:"expressMultiplier"`
						ExpressSupported  bool    `json:"expressSupported"`
					} `json:"axelarFeeData"`
					ToChainFeeData struct {
						LastBaseFeePerGas    string `json:"lastBaseFeePerGas"`
						MaxFeePerGas         string `json:"maxFeePerGas"`
						MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
						GasPrice             string `json:"gasPrice"`
					} `json:"toChainFeeData"`
				} `json:"data,omitempty"`
			} `json:"feeCosts"`
			GasCosts []struct {
				Type  string `json:"type"`
				Token struct {
					Id                  string   `json:"id"`
					Symbol              string   `json:"symbol"`
					Address             string   `json:"address"`
					ChainId             string   `json:"chainId"`
					Name                string   `json:"name"`
					Decimals            int      `json:"decimals"`
					CoingeckoId         string   `json:"coingeckoId"`
					Type                string   `json:"type"`
					LogoURI             string   `json:"logoURI"`
					Volatility          int      `json:"volatility"`
					AxelarNetworkSymbol string   `json:"axelarNetworkSymbol"`
					SubGraphOnly        bool     `json:"subGraphOnly"`
					SubGraphIds         []string `json:"subGraphIds"`
					Enabled             bool     `json:"enabled"`
					Active              bool     `json:"active"`
					Visible             bool     `json:"visible"`
					UsdPrice            float64  `json:"usdPrice"`
				} `json:"token"`
				Amount    string `json:"amount"`
				GasLimit  string `json:"gasLimit"`
				AmountUsd string `json:"amountUsd"`
			} `json:"gasCosts"`
			EstimatedRouteDuration int `json:"estimatedRouteDuration"`
		} `json:"estimate"`
		TransactionRequest struct {
			Type                 string `json:"type"`
			Target               string `json:"target"`
			Data                 string `json:"data"`
			Value                string `json:"value"`
			GasLimit             string `json:"gasLimit"`
			LastBaseFeePerGas    string `json:"lastBaseFeePerGas"`
			MaxFeePerGas         string `json:"maxFeePerGas"`
			MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
			GasPrice             string `json:"gasPrice"`
			RequestId            string `json:"requestId"`
			Expiry               string `json:"expiry"`
			ExpiryOffset         string `json:"expiryOffset"`
		} `json:"transactionRequest"`
		Params struct {
			FromAddress string `json:"fromAddress"`
			FromChain   string `json:"fromChain"`
			FromToken   string `json:"fromToken"`
			FromAmount  string `json:"fromAmount"`
			ToChain     string `json:"toChain"`
			ToToken     string `json:"toToken"`
			ToAddress   string `json:"toAddress"`
		} `json:"params"`
		QuoteId string `json:"quoteId"`
	} `json:"route"`
}
