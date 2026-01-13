package stargate

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/chain"
	"main/tokens"
	"main/utils"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func StargateFastBridge(fromChain, fromToken string, fromAddress common.Address, amount *big.Int, data Quotes, pk *ecdsa.PrivateKey) (string, error) {
	var ctx = context.Background()
	client := chain.ClientFromChain[fromChain]
	token := common.HexToAddress(fromToken)

	target := ""
	for _, quotes := range data.Steps {
		if quotes.Type == "bridge" {
			target = quotes.Transaction.To
		}
	}

	for _, quotes := range data.Steps {
		// ERC20
		if quotes.Type == "approve" {
			allowance, err := tokens.CheckAllowance(token, common.HexToAddress(target), fromAddress, client)
			if err != nil {
				return "", fmt.Errorf("get allowance error: %v", err)
			}

			if allowance.Cmp(amount) == -1 {
				//hash, err := chain.BuildAndSendTransaction(
				//	ctx,
				//	client,
				//	fromAddress,
				//	common.HexToAddress(quotes.Transaction.To),
				//	quotes.Transaction.Data,
				//	quotes.Transaction.Value,
				//	"100000",
				//	pk)
				amountInt := utils.ParseApprovalFloat()
				hash, err := tokens.TokenApprove(ctx, common.HexToAddress(quotes.Transaction.To), common.HexToAddress(target), fromAddress, amountInt, pk, client)
				if err == nil && hash != "" {
					for {
						//time.Sleep(500 * time.Millisecond)
						txOnchain, txStatus := utils.ReadTxStatus(hash, client)
						if txOnchain && txStatus == 0 {
							log.Fatalf("syna Approve 授权失败")
						}
						if txOnchain && txStatus == 1 {
							break
						}
						time.Sleep(1 * time.Second)
					}
				}
			}
		} else if quotes.Type == "bridge" {
			return chain.BuildAndSendTransaction(
				ctx,
				client,
				fromAddress,
				common.HexToAddress(quotes.Transaction.To),
				quotes.Transaction.Data,
				quotes.Transaction.Value,
				"400000",
				pk)
		}
	}

	return "", nil

}
