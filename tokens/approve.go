package tokens

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TokenApprove 执行代币授权操作，允许目标地址花费指定数量的代币
//
// 参数:
//
//	ctx - 上下文，用于控制超时和取消
//	token - 代币合约地址
//	Target - 被授权的目标地址
//	fromAddress - 发送者地址
//	amount - 授权数量
//	pk - 发送者的私钥
//	client - Ethereum 客户端
//
// 返回:
//
//	string - 交易哈希
//	error - 错误信息
//
// 示例:
//
//	hash, err := TokenApprove(ctx, tokenAddr, targetAddr, senderAddr, amount, privateKey, client)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Approve transaction hash:", hash)
func TokenApprove(ctx context.Context, token, Target, fromAddress common.Address, amount *big.Int, pk *ecdsa.PrivateKey, client *ethclient.Client) (string, error) {
	// 创建代币合约实例
	instance, err := NewOik(token, client)
	if err != nil {
		return "", fmt.Errorf("构建 Approve 交易Instance失败: %v", err)
	}

	// 获取链ID并创建交易选项
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("获取链ID失败: %v", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return "", fmt.Errorf("绑定链数据失败: %v", err)
	}

	// 获取发送者的当前nonce
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", fmt.Errorf("获取地址 %s nonce失败: %v", fromAddress, err)
	}
	opts.Nonce = big.NewInt(int64(nonce))

	// 设置Gas参数
	opts.GasLimit = uint64(150000) // 固定Gas限制

	// 获取建议的Gas价格并适当调整
	suggestPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("获取GasPrice失败: %v", err)
	}
	// Gas价格调整为建议价格的1.2倍，提高交易被打包的速度
	opts.GasPrice = new(big.Int).Mul(suggestPrice, big.NewInt(12))
	opts.GasPrice = opts.GasPrice.Div(opts.GasPrice, big.NewInt(10))

	// 执行授权交易
	trans, err := instance.Approve(opts, Target, amount)
	if err != nil {
		return "", fmt.Errorf("发送Approve 交易失败 [sender: %s, target: %s, amount: %s]: %v",
			fromAddress, Target, amount, err)
	}

	return trans.Hash().Hex(), nil
}
