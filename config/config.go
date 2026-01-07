package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

// config包负责管理应用程序的配置信息
// 提供对非区块链网络配置的结构化管理，包括API端点、日志配置等
// 主要特性：单例模式初始化、安全的配置访问、配置验证

// 全局配置存储
var (
	// once 用于确保配置初始化只执行一次的同步原语
	once sync.Once
	// Config 日志配置
	LogConfig Config
	// apiConfig API相关配置
	apiConfig APIConfig
	// 是否初始化完成
	initialized bool
)

// Config 日志配置结构体
type Config struct {
	LogFile    string // 日志文件路径
	ExecuePath string // 执行路径
	BaseTime   int    // 基础时间配置
}

// APIConfig API配置结构体
type APIConfig struct {
	// Squid相关配置
	SquidHost  string
	SquidRoute string
	// Stargate相关配置
	StargateHost  string
	StargateQuote string
	// Syna相关配置
	SynaHost  string
	SynaQuote string
}

// 配置验证错误
var (
	errInvalidLogFile = fmt.Errorf("日志文件路径不能为空")
	errInvalidAPIHost = fmt.Errorf("API主机地址不能为空")
)

// InitConfig 初始化所有配置
// 使用sync.Once确保配置只被初始化一次
func InitConfig() error {

	errChan := make(chan error, 1)

	err := godotenv.Load()
	if err != nil {
		errChan <- fmt.Errorf("error loading .env file: %v", err)
	}

	once.Do(func() {
		// 初始化默认配置
		LogConfig = Config{
			LogFile:    "running.txt",
			ExecuePath: "",
			BaseTime:   5,
		}

		apiConfig = APIConfig{
			// Squid相关配置
			SquidHost:  "https://v2.api.squidrouter.com/v2",
			SquidRoute: "/route",
			// Stargate相关配置
			StargateHost:  "https://stargate.finance/api/v1",
			StargateQuote: "/quotes",
			// Syna相关配置
			SynaHost:  "https://api.synapseprotocol.com",
			SynaQuote: "/bridge/v2",
		}

		// 验证配置
		if err := validateConfig(); err != nil {
			errChan <- fmt.Errorf("配置验证失败: %v", err)
			return
		}

		initialized = true
		log.Printf("配置初始化成功")
		errChan <- nil
	})

	return <-errChan
}

// validateConfig 验证配置的有效性
func validateConfig() error {
	// 验证日志配置
	if LogConfig.LogFile == "" {
		return errInvalidLogFile
	}

	// 验证API配置
	if apiConfig.SquidHost == "" || apiConfig.StargateHost == "" || apiConfig.SynaHost == "" {
		return errInvalidAPIHost
	}

	return nil
}

// init 函数在包被导入时自动执行
// 调用InitConfig初始化配置
func init() {
	if err := InitConfig(); err != nil {
		log.Printf("警告: 配置初始化失败: %v，继续使用默认配置", err)
	}
}

// GetLogConfig 安全获取日志配置
func GetLogConfig() Config {
	return LogConfig
}

// GetAPIConfig 安全获取API配置
func GetAPIConfig() APIConfig {
	return apiConfig
}

// GetSquidConfig 获取Squid API配置
func GetSquidConfig() (string, string) {
	return apiConfig.SquidHost, apiConfig.SquidRoute
}

// GetPortalConfig 获取Portal API配置
func GetStargateConfig() (string, string) {
	return apiConfig.StargateHost, apiConfig.StargateQuote
}

// GetSynaConfig 获取Syna API配置
func GetSynaConfig() (string, string) {
	return apiConfig.SynaHost, apiConfig.SynaQuote
}

// IsInitialized 检查配置是否已初始化
func IsInitialized() bool {
	return initialized
}
