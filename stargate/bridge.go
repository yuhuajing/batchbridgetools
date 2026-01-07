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

func StargateFastBridge(fromChain string, fromAddress common.Address, amount *big.Int, data Quotes, pk *ecdsa.PrivateKey) (string, error) {
	var ctx = context.Background()
	client := chain.ClientFromChain[fromChain]

	//amountIn, _ := utils.StringToBigInt(amount)
	quotes := data.Steps[0].Transaction
	t := chain.TokenContracts[fromChain+"_"+chain.STARGATE]
	token := common.HexToAddress(t)
	// ERC20
	if token != common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
		allowance, err := tokens.CheckAllowance(token, common.HexToAddress(quotes.To), fromAddress, client)
		if err != nil {
			return "", fmt.Errorf("get allowance error: %v", err)
		}

		if allowance.Cmp(amount) == -1 {
			hash, err := tokens.TokenApprove(ctx, token, common.HexToAddress(quotes.To), fromAddress, amount, pk, client)
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
	}
	return chain.BuildAndSendTransaction(
		ctx,
		client,
		fromAddress,
		common.HexToAddress(quotes.To),
		quotes.Data,
		quotes.Value,
		"400000",
		pk)
}
