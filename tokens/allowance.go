package tokens

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// CheckAllowance 检查代币授权额度，返回目标地址可以花费的代币数量
//
// 参数:
//   token - 代币合约地址
//   target - 被授权的目标地址
//   fromAddress - 发送者地址
//   client - Ethereum 客户端
//
// 返回:
//   *big.Int - 授权额度
//   error - 错误信息
//
// 示例:
//   allowance, err := CheckAllowance(tokenAddr, targetAddr, senderAddr, client)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Println("Current allowance:", allowance.String())
func CheckAllowance(token, target, fromAddress common.Address, client *ethclient.Client) (*big.Int, error) {
	// 创建代币合约调用实例
	instance, err := NewOikCaller(token, client)
	if err != nil {
		return nil, fmt.Errorf("构建CheckAllowance交易失败: %v", err)
	}

	// 调用合约的Allowance方法检查授权额度
	allowance, err := instance.Allowance(nil, fromAddress, target)
	if err != nil {
		return nil, fmt.Errorf("发起CheckAllowance查询失败: %v", err)
	}

	return allowance, nil
}
