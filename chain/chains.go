package chain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// 链名称常量
const (
	BNB  = "bsc"
	BASE = "base"
	ETH  = "ethereum"
	ARB  = "arbitrum"
	POLY = "polygon"
)

const (
	SQUID    = "squid"
	SYNA     = "syna"
	STARGATE = "stargate"
)

const (
	ETHToken = "eth"
	USDC     = "usdc"
)

var BridgeTools = []string{SQUID, SYNA, STARGATE}

// ClientConfig - 连接配置结构体
type ClientConfig struct {
	MaxRetries    int           // 最大重试次数
	InitialDelay  time.Duration // 初始重试延迟
	MaxDelay      time.Duration // 最大重试延迟
	ConnectionTTL time.Duration // 连接最大生存时间
}

type ChainConfig struct {
	RPC string // RPC端点URL
	Id  string
}

var ChainConfigs = map[string]ChainConfig{
	BNB: {
		RPC: "https://bsc.drpc.org",
		Id:  "56",
	},
	//ETH: {
	//	RPC: "https://eth.llamarpc.com",
	//	Id:  "1",
	//},
	ARB: {
		RPC: "https://arbitrum.rpc.subquery.network/public",
		Id:  "42161",
	},
	//POLY: {
	//	RPC: "https://polygon.drpc.org",
	//	Id:  "137",
	//},
	BASE: {
		RPC: "https://base.drpc.org",
		Id:  "8453",
	},
}

// 默认连接配置
var DefaultClientConfig = ClientConfig{
	MaxRetries:    3,
	InitialDelay:  2 * time.Second,
	MaxDelay:      30 * time.Second,
	ConnectionTTL: 24 * time.Hour,
}

// ClientInfo - 客户端信息结构体，包含客户端实例和元数据
type ClientInfo struct {
	Client     *ethclient.Client
	CreatedAt  time.Time
	LastUsedAt time.Time
	Config     ClientConfig
}

var (
	ClientFromChain map[string]*ethclient.Client // 普通客户端映射

	clientsMutex      sync.RWMutex           // 保护客户端映射表的互斥锁
	clientsInfo       map[string]*ClientInfo // 存储普通客户端详细信息
	filterClientsInfo map[string]*ClientInfo // 存储过滤客户端详细信息
)

// TokenContractConfig - 代币合约配置结构体

// 代币合约地址映射
var TokenContracts = map[string]string{
	//ETH + "_" + SQUID:  "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	ARB + "_" + ETHToken + "_" + SQUID:    "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	ARB + "_" + ETHToken + "_" + SYNA:     "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	ARB + "_" + ETHToken + "_" + STARGATE: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	ARB + "_" + USDC + "_" + SQUID:        "0xaf88d065e77c8cc2239327c5edb3a432268e5831",
	//ARB + "_" + USDC + "_" + SYNA:         "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
	ARB + "_" + USDC + "_" + STARGATE: "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",

	BASE + "_" + ETHToken + "_" + SQUID:    "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	BASE + "_" + ETHToken + "_" + SYNA:     "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	BASE + "_" + ETHToken + "_" + STARGATE: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
	BASE + "_" + USDC + "_" + SQUID:        "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913",
	//BASE + "_" + USDC + "_" + SYNA:         "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
	BASE + "_" + USDC + "_" + STARGATE: "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",

	BNB + "_" + ETHToken + "_" + SQUID: "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
}

var TokenDecimals = map[string]float64{
	ETHToken: params.Ether,
	USDC:     params.GWei / 1000,
}

func init() {
	ClientFromChain = make(map[string]*ethclient.Client)
	clientsInfo = make(map[string]*ClientInfo)
	filterClientsInfo = make(map[string]*ClientInfo)
}

// SafeBuildClient - 安全构建指定链的以太坊客户端
// 参数:
//
//	fromChain - 链名称（如 bsc, ethereum, arbitrum 等）
//	ctx - 可选的上下文，用于控制超时和取消
//
// 返回:
//
//	*ethclient.Client - 以太坊客户端实例
//	error - 构建过程中的错误
func SafeBuildClient(fromChain string, ctx ...context.Context) (*ethclient.Client, error) {
	// 使用默认上下文或传入的上下文
	context := context.Background()
	if len(ctx) > 0 {
		context = ctx[0]
	}

	// 检查连接是否已存在且有效
	if client := getCachedClient(fromChain); client != nil {
		return client, nil
	}

	// 从配置中获取RPC URL
	chainConfig, exists := ChainConfigs[fromChain]
	if !exists {
		return nil, fmt.Errorf("链 %s 的配置未找到", fromChain)
	}
	rpcURL := chainConfig.RPC
	if rpcURL == "" {
		return nil, fmt.Errorf("链 %s 的RPC URL未配置", fromChain)
	}

	// 安全连接区块链
	client, err := SafeConnBlockchainWithContext(context, rpcURL, DefaultClientConfig)
	if err != nil {
		return nil, fmt.Errorf("构建客户端失败: %w", err)
	}

	// 存储客户端实例到映射表
	saveClientToCache(fromChain, client, DefaultClientConfig)
	log.Printf("成功为链 %s 构建客户端", fromChain)
	return client, nil
}

// BuildClient - 兼容旧代码的客户端构建函数
func BuildClient(fromChain string) error {
	_, err := SafeBuildClient(fromChain)
	return err
}

// SafeConnBlockchain - 安全连接区块链，支持自动重试机制
func SafeConnBlockchain(rpcURL string) (*ethclient.Client, error) {
	return SafeConnBlockchainWithContext(context.Background(), rpcURL, DefaultClientConfig)
}

// GetChainID - 获取指定链的链ID
// 参数:
//
//	chainName - 链名称（如 bsc, ethereum, arbitrum 等）
//
// 返回:
//
//	string - 链ID字符串
//	error - 获取过程中的错误
func GetChainID(chainName string) (string, error) {
	chainConfig, exists := ChainConfigs[chainName]
	if !exists {
		return "", fmt.Errorf("链 %s 的配置未找到", chainName)
	}
	if chainConfig.Id == "" {
		return "", fmt.Errorf("链 %s 的ID未配置", chainName)
	}
	return chainConfig.Id, nil
}

// SafeConnBlockchainWithContext - 带上下文的安全连接区块链函数
func SafeConnBlockchainWithContext(ctx context.Context, rpcURL string, config ClientConfig) (*ethclient.Client, error) {
	if config.MaxRetries <= 0 {
		config.MaxRetries = DefaultClientConfig.MaxRetries
	}
	if config.InitialDelay <= 0 {
		config.InitialDelay = DefaultClientConfig.InitialDelay
	}
	if config.MaxDelay <= 0 {
		config.MaxDelay = DefaultClientConfig.MaxDelay
	}

	retryDelay := config.InitialDelay
	var lastErr error

	for i := 0; i < config.MaxRetries; i++ {
		// 创建带有超时的子上下文
		connCtx, cancel := context.WithTimeout(ctx, 30*time.Second)

		// 尝试连接到区块链节点
		client, err := ethclient.DialContext(connCtx, rpcURL)
		cancel() // 确保及时释放资源

		if err == nil {
			// 验证连接是否正常工作
			if ctx.Err() == nil { // 检查外部上下文是否已取消
				if _, err := client.ChainID(ctx); err == nil {
					log.Printf("成功连接到区块链: %s (尝试 %d/%d)", rpcURL, i+1, config.MaxRetries)
					return client, nil
				}
				// 连接已建立但功能异常
				client.Close()
				lastErr = fmt.Errorf("连接已建立但功能异常: %w", err)
			} else {
				client.Close()
				return nil, ctx.Err() // 返回外部上下文的错误
			}
		} else {
			lastErr = fmt.Errorf("连接失败: %w", err)
		}

		log.Printf("连接区块链失败 (尝试 %d/%d): %v", i+1, config.MaxRetries, lastErr)

		// 如果不是最后一次尝试，则等待后重试
		if i < config.MaxRetries-1 {
			// 检查外部上下文是否已取消
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}

			// 等待并使用指数退避策略，但不超过最大延迟
			time.Sleep(retryDelay)
			retryDelay *= 2
			if retryDelay > config.MaxDelay {
				retryDelay = config.MaxDelay
			}
		}
	}

	return nil, fmt.Errorf("超过最大重试次数 %d，连接 %s 失败: %w", config.MaxRetries, rpcURL, lastErr)
}

// CloseClient - 关闭指定链的客户端连接
func CloseClient(fromChain string) error {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	client, exists := ClientFromChain[fromChain]
	if !exists {
		return fmt.Errorf("链 %s 的客户端不存在", fromChain)
	}

	// 关闭客户端连接
	client.Close()

	// 从映射表中删除
	delete(ClientFromChain, fromChain)
	delete(clientsInfo, fromChain)
	log.Printf("成功关闭并移除链 %s 的客户端", fromChain)
	return nil
}

// CloseAllClients - 关闭所有客户端连接
func CloseAllClients() map[string]error {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	errors := make(map[string]error)
	// 关闭普通客户端
	for chain, client := range ClientFromChain {
		client.Close()
		log.Printf("成功关闭链 %s 的客户端", chain)
		delete(ClientFromChain, chain)
		delete(clientsInfo, chain)
	}

	return errors
}

// GetClientStats - 获取客户端统计信息（调试用）
func GetClientStats() map[string]map[string]interface{} {
	clientsMutex.RLock()
	defer clientsMutex.RUnlock()

	stats := make(map[string]map[string]interface{})
	// 统计普通客户端
	for chain, info := range clientsInfo {
		stats[chain+"_client"] = map[string]interface{}{
			"type":         "client",
			"created_at":   info.CreatedAt,
			"last_used_at": info.LastUsedAt,
			"uptime":       time.Since(info.CreatedAt),
			"max_retries":  info.Config.MaxRetries,
		}
	}
	// 统计过滤客户端
	for chain, info := range filterClientsInfo {
		stats[chain+"_filter"] = map[string]interface{}{
			"type":         "filter",
			"created_at":   info.CreatedAt,
			"last_used_at": info.LastUsedAt,
			"uptime":       time.Since(info.CreatedAt),
			"max_retries":  info.Config.MaxRetries,
		}
	}
	return stats
}

// 内部辅助函数

// getCachedClient - 从缓存中获取有效的客户端实例
func getCachedClient(fromChain string) *ethclient.Client {
	clientsMutex.RLock()
	defer clientsMutex.RUnlock()

	client, exists := ClientFromChain[fromChain]
	if !exists {
		return nil
	}

	// 检查连接是否过期
	if info, infoExists := clientsInfo[fromChain]; infoExists {
		if time.Since(info.CreatedAt) > info.Config.ConnectionTTL {
			return nil
		}
		// 更新最后使用时间
		info.LastUsedAt = time.Now()
	}
	return client
}

// saveClientToCache - 将客户端实例保存到缓存中
func saveClientToCache(fromChain string, client *ethclient.Client, config ClientConfig) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	// 如果已存在客户端，先关闭旧连接
	if existing, exists := ClientFromChain[fromChain]; exists {
		existing.Close()
	}

	// 保存新客户端
	ClientFromChain[fromChain] = client
	clientsInfo[fromChain] = &ClientInfo{
		Client:     client,
		CreatedAt:  time.Now(),
		LastUsedAt: time.Now(),
		Config:     config,
	}
}
