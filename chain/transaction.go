package chain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"main/utils"
	"math/big"
	"strings"
)

// DynamicGasPriceParams 动态gas价格参数

// GetDynamicGasPrice 获取动态gas价格（baseFee, priorityFee, gasFeeCap）
func GetDynamicGasPrice(ctx context.Context, client *ethclient.Client) (*big.Int, *big.Int, *big.Int, error) {
	// Suggest the base fee for inclusion in a block.
	baseFee, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("suggest gas price error: %v", err)
	}

	// Suggest a gas tip cap (priority fee) for miner incentive.
	priorityFee, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("suggest gas tip cap error: %v", err)
	}

	// Calculate the maximum gas fee cap, adding a 2 GWei margin to the base fee plus priority fee.
	increment := new(big.Int).Mul(big.NewInt(2), big.NewInt(params.GWei))
	gasFeeCap := new(big.Int).Add(baseFee, increment)
	gasFeeCap.Add(gasFeeCap, priorityFee)

	return baseFee, priorityFee, gasFeeCap, nil
}

// BuildAndSendTransaction 构建并发送交易
func BuildAndSendTransaction(ctx context.Context, client *ethclient.Client, sender, toAddr common.Address, data string, value string, gasLimit string, privateKey *ecdsa.PrivateKey) (string, error) {
	// Prepare data payload.
	if ok := strings.HasPrefix(data, "0x"); !ok {
		data = "0x" + data
	}

	bytesData, err := hexutil.Decode(data)
	if err != nil {
		return "", fmt.Errorf("encode calldata error: %v", err)
	}

	valueBig, err := utils.StringToBigInt(value)
	if err != nil {
		return "", err
	}

	// Get dynamic gas price
	_, priorityFee, gasFeeCap, err := GetDynamicGasPrice(ctx, client)
	if err != nil {
		return "", err
	}

	// Get chain ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("get chainId error: %v", err)
	}

	// Parse gas limit
	gasLimitBig, err := utils.StringToBigInt(gasLimit)
	if err != nil {
		return "", err
	}

	// Get nonce
	nonce, err := client.PendingNonceAt(ctx, sender)
	if err != nil {
		return "", fmt.Errorf("get nonce of user error: %v", err)
	}

	// Build transaction
	txData := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimitBig.Uint64(),
		To:        &toAddr,
		Value:     valueBig,
		Data:      bytesData,
	}

	tx := types.NewTx(&txData)

	// Sign the transaction with the private key of the sender.
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("sign transactions error: %v", err)
	}

	// Encode the signed transaction into RLP (Recursive Length Prefix) format for transmission.
	var buf bytes.Buffer
	err = signedTx.EncodeRLP(&buf)
	if err != nil {
		return "", fmt.Errorf("encode signed transactions error: %v", err)
	}

	// Return the RLP-encoded transaction as a hexadecimal string.
	//	rawTxRLPHex := hex.EncodeToString(buf.Bytes())

	// Send transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", fmt.Errorf("send transaction onchain error: %v", err)
	}

	// Wait for transaction to be processed
	//time.Sleep(2 * time.Second)

	return fmt.Sprintf("%s", signedTx.Hash().Hex()), nil
}
