package syna

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

func SynaFastBridge(fromChain string, fromAddress common.Address, amount *big.Int, data BridgeRoute, pk *ecdsa.PrivateKey) (string, error) {
	var ctx = context.Background()
	client := chain.ClientFromChain[fromChain]

	//amountIn, _ := utils.StringToBigInt(amount)

	t := chain.TokenContracts[fromChain+"_"+chain.SYNA]
	token := common.HexToAddress(t)
	// ERC20
	if token != common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
		allowance, err := tokens.CheckAllowance(token, common.HexToAddress(data.RouterAddress), fromAddress, client)
		if err != nil {
			return "", fmt.Errorf("get allowance error: %v", err)
		}

		if allowance.Cmp(amount) == -1 {
			hash, err := tokens.TokenApprove(ctx, token, common.HexToAddress(data.RouterAddress), fromAddress, amount, pk, client)
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
		common.HexToAddress(data.CallData.To),
		data.CallData.Data,
		data.CallData.Value,
		"200000",
		pk)
}
