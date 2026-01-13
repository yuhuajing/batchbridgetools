package spoke

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/chain"
	"main/tokens"
	"main/utils"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func SquidDynamicTransaction(fromChain, fromToken, amount string, sender common.Address, to, data, value, gasLimit string, privateKey *ecdsa.PrivateKey) (string, error) {
	var ctx = context.Background()
	client := chain.ClientFromChain[fromChain]

	token := common.HexToAddress(fromToken)
	// ERC20
	if token != common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee") {
		allowance, err := tokens.CheckAllowance(token, common.HexToAddress(to), sender, client)
		if err != nil {
			return "", fmt.Errorf("get allowance error: %v", err)
		}

		amountInt := utils.ParseApprovalFloat()
		if allowance.Cmp(amountInt) == -1 {
			hash, err := tokens.TokenApprove(ctx, token, common.HexToAddress(to), sender, amountInt, privateKey, client)
			if err == nil && hash != "" {
				for {
					txOnchain, txStatus := utils.ReadTxStatus(hash, client)
					if txOnchain && txStatus == 0 {
						log.Fatalf("squid Approve 授权失败")
					}
					if txOnchain && txStatus == 1 {
						break
					}
					time.Sleep(1 * time.Second)
				}
			}
		}
	}

	// 使用通用的交易构建和发送函数
	return chain.BuildAndSendTransaction(ctx, client, sender, common.HexToAddress(to), data, value, gasLimit, privateKey)
}
