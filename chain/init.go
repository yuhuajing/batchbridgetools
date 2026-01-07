package chain

import "log"

// InitChainClients 初始化所有链的客户端连接
func InitChainClients() {
	for chainName, _ := range ChainConfigs {
		// 构建区块链客户端
		if err := BuildClient(chainName); err != nil {
			log.Printf("警告: 无法为链 %s 构建客户端: %v", chainName, err)
		}
	}
}
