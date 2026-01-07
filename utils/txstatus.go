package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ReadTxStatus(txhash string, client *ethclient.Client) (bool, uint64) { //pendingstatus, status
	_, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(txhash))
	if isPending || err != nil {
		return false, 0
	}
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txhash))
	if err != nil {
		return false, 0
	}
	return true, receipt.Status
}
