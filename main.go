package main

import (
	"main/chain"
	"main/explorer"
)

func main() {
	// 显式初始化区块链客户端
	chain.InitChainClients()

	explorer.Explorer()
	//portal.EncodeMayanDataExamples()
	//portal.SwapAndForwardExamples()
}
