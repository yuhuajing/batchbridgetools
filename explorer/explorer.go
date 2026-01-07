package explorer

import (
	"crypto/ecdsa"
	"main/stargate"
	"math/big"
	"mime/multipart"
	"time"

	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"main/chain"
	"main/config"
	"main/syna"
	"math/rand"
	"net/http"

	"main/squid"
	"main/squid/spoke"

	"main/utils"

	fiber "github.com/gofiber/fiber/v2"
	html "github.com/gofiber/template/html/v2"
	"log"
	"path/filepath"
	"strings"
)

// Explorer 启动区块链跨链交易探索器服务
// 提供 Web 界面用于上传 CSV 文件并执行跨链交易
func Explorer() {
	app := fiber.New(fiber.Config{
		Views: html.New(filepath.Join(config.GetLogConfig().ExecuePath, "explorer", "views"), ".html"),
	})
	app.Static("/", filepath.Join(config.GetLogConfig().ExecuePath, "explorer", "public"))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	// 添加专门的路由来提供 running.txt 文件
	app.Get("/running.txt", func(c *fiber.Ctx) error {
		// 返回根目录下的 running.txt 文件
		return c.SendFile(filepath.Join(config.GetLogConfig().ExecuePath, "running.txt"))
	})

	app.Post("/bridge", handleBridge)
	app.Post("/squidbridge", func(c *fiber.Ctx) error { return handleBridgeWithTool(c, chain.SQUID) })
	app.Post("/stargatebridge", func(c *fiber.Ctx) error { return handleBridgeWithTool(c, chain.STARGATE) })
	app.Post("/synabridge", func(c *fiber.Ctx) error { return handleBridgeWithTool(c, chain.SYNA) })

	log.Fatal(app.Listen(":3005"))
}

type DataResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

// 统一的桥接处理函数，支持所有跨链工具
func handleBridge(c *fiber.Ctx) error {
	return handleBridgeWithTools(c, []string{chain.SQUID, chain.SYNA, chain.STARGATE})
}

// 针对特定跨链工具的桥接处理函数
func handleBridgeWithTool(c *fiber.Ctx, tool string) error {
	return handleBridgeWithTools(c, []string{tool})
}

func handleBridgeWithTools(c *fiber.Ctx, tools []string) error {
	fromChain := c.FormValue("fromChain")
	toChain := c.FormValue("toChain")

	// 从表单数据中获取 CSV 文件
	handler, err := c.FormFile("file")
	if err != nil {
		msg := fmt.Sprintf("CSV文件上传失败: %s", err.Error())
		utils.Intolog(msg)
		return c.Status(400).JSON(DataResponse{
			Error:   msg,
			Success: false,
		})
	}
	err = bridge(fromChain, toChain, tools, handler)
	if err != nil {
		utils.Intolog(err.Error())
		return c.Status(400).JSON(DataResponse{
			Error:   err.Error(),
			Success: false,
		})
	}
	return c.JSON(DataResponse{
		Error:   "",
		Success: true,
	})
}

func bridge(fromChain, toChain string, tools []string, handler *multipart.FileHeader) error {
	if fromChain == "" || toChain == "" || fromChain == toChain {
		return fmt.Errorf("源链： %s 和目标链: %s 不能为空或者相同", fromChain, toChain)
	}

	// 打开 CSV 文件
	file, err := handler.Open()
	if err != nil {
		return fmt.Errorf("无法打开CSV文件: %s", err.Error())

	}
	defer file.Close()

	// 读取并解析 CSV 文件
	lines, err := utils.ReadCSVFile(file)
	if err != nil {
		return fmt.Errorf("读取CSV文件数据错误: %s", err.Error())

	}

	// 创建 HTTP 客户端
	httpclient, err := utils.CreateHTTPClient()
	if err != nil {
		return fmt.Errorf("创建HTTP客户端失败: %s", err.Error())
	}

	// 遍历 CSV 文件中的每一行，执行跨链交易
	for index, line := range lines {
		msg := fmt.Sprintf("处理第 %d 行数据", index+1)
		utils.Intolog(msg)

		// 解析 CSV 行数据
		privateKey := line[0]
		receiver := common.HexToAddress(line[1])
		amount := strings.Trim(line[2], " ") //utils.ParseFloat(strings.Trim(line[2], " "))

		// 解析私钥
		ecdsaPrivateKey, sender, err := utils.ParsePk(privateKey)
		if err != nil {
			msg = fmt.Sprintf("---第 %d 行数据，解析私钥失败: %v\n", index+1, err)
			utils.Intolog(msg)
			continue
		}

		tool := tools[rand.Intn(len(tools))]

		switch tool {
		case chain.SQUID:
			err = squids(fromChain, toChain, amount, sender, receiver, httpclient, index, ecdsaPrivateKey)
		case chain.STARGATE:
			err = stargates(fromChain, toChain, amount, sender, receiver, httpclient, index, ecdsaPrivateKey)
		case chain.SYNA:
			err = synas(fromChain, toChain, amount, sender, receiver, httpclient, index, ecdsaPrivateKey)
		default:
			err = fmt.Errorf("不支持的跨链工具")
		}
		if err != nil {
			utils.Intolog(msg)
			continue
		} else {
			// 随机延迟，避免请求频率过高
			sleepDuration := time.Duration(rand.Intn(5)+config.GetLogConfig().BaseTime) * time.Minute
			msg = fmt.Sprintf("处理第 %d 行数据, 等待 %d 分钟", index+1, sleepDuration/time.Minute)
			utils.Intolog(msg)
			time.Sleep(sleepDuration)
		}
	}
	return nil
}

func squids(fromChain, toChain, amount string, sender, receiver common.Address, httpclient *http.Client, index int, pk *ecdsa.PrivateKey) error {
	// 获取 Squid 路由
	amount = utils.ParseFloat(amount)
	utils.Intolog(fmt.Sprintf("---获取Squid-Route中,from :%s , to: %s， value: %s", fromChain, toChain, amount))
	fromToken, toToken, err := getToken(fromChain, toChain, chain.SQUID)
	if err != nil {
		return err
	}
	target, data, value, limit, err := squid.SendSquidRouteRequest(httpclient, squid.QuoteRequest{
		FromAddress: sender.Hex(),
		FromChain:   chain.ChainConfigs[fromChain].Id,
		FromToken:   fromToken,
		FromAmount:  amount,
		ToChain:     chain.ChainConfigs[toChain].Id,
		ToToken:     toToken,
		ToAddress:   receiver.Hex(),
	})
	if err != nil {
		return fmt.Errorf("###第 %d 行 获取Squid-Route请求失败: %v\n", index+1, err)

	}
	utils.Intolog("---获取Squid-Route成功")

	// 发送 Squid 跨链交易
	utils.Intolog(fmt.Sprintf("---发送Squid跨链交易中,from：%s, to: %s, amount: %s", fromChain, toChain, amount))
	hash, err := spoke.SquidDynamicTransaction(fromChain, amount, sender, target, data, value, limit, pk)
	if err != nil {
		return fmt.Errorf("###第 %d 行 发送Squid跨链交易失败: %v \n", index+1, err.Error())
	}
	utils.Intolog(fmt.Sprintf("!!!第 %d 行 Squid跨链交易已发送上链,,fromChain: %s,toChain: %s, Amount: %s hash: %v \n", index+1, fromChain, toChain, amount, hash))
	return nil
}

func synas(fromChain, toChain, amount string, sender, receiver common.Address, httpclient *http.Client, index int, pk *ecdsa.PrivateKey) error {
	// 获取 Syna 路由
	amount = utils.ParseFloat(amount)
	utils.Intolog(fmt.Sprintf("---获取Syna-Route中,from :%s , to: %s， value: %s", fromChain, toChain, amount))
	fromToken, toToken, err := getToken(fromChain, toChain, chain.SYNA)
	if err != nil {
		return err
	}
	amountInt := utils.HexStringToBigInt(utils.ParseFloat(amount))
	payload := map[string]string{
		"fromChainId": chain.ChainConfigs[fromChain].Id,
		"fromToken":   fromToken,
		"toChainId":   chain.ChainConfigs[toChain].Id,
		"toToken":     toToken,
		"fromAmount":  amount,
		"fromSender":  sender.Hex(),
		"toRecipient": receiver.Hex(),
	}

	quote, err := syna.GetSynaBridgeRoute(httpclient, payload)
	if err != nil {
		return fmt.Errorf("###第 %d 行 获取Syna-Route请求失败: %v\n", index+1, err)
	}
	utils.Intolog("---获取Syna-Route成功")

	// 发送 Syna 跨链交易
	utils.Intolog("---发送Syna跨链交易中")
	//amount := utils.ParseFloat(amountStr)
	hash, err := syna.SynaFastBridge(fromChain, sender, amountInt, quote, pk)
	if err != nil {
		return fmt.Errorf("###第 %d 行 发送Syna跨链交易失败: %v \n", index+1, err.Error())
	}
	utils.Intolog(fmt.Sprintf("!!!第 %d 行 Syna跨链交易已发送上链,,fromChain: %s,toChain: %s, Amount: %s hash: %v \n", index+1, fromChain, toChain, amount, hash))
	return nil
}

func stargates(fromChain, toChain, amount string, sender, receiver common.Address, httpclient *http.Client, index int, pk *ecdsa.PrivateKey) error {
	// 获取 Stargate 路由
	utils.Intolog(fmt.Sprintf("---获取Stargate-Route中,from :%s , to: %s， value: %s", fromChain, toChain, amount))
	fromToken, toToken, err := getToken(fromChain, toChain, chain.STARGATE)
	if err != nil {
		return err
	}

	amountInt := utils.HexStringToBigInt(utils.ParseFloat(amount))
	amountIntMin := new(big.Int).Mul(new(big.Int).Div(amountInt, big.NewInt(100)), big.NewInt(95))
	payload := map[string]string{
		"srcToken":     fromToken,
		"dstToken":     toToken,
		"srcChainKey":  fromChain,
		"dstChainKey":  toChain,
		"srcAddress":   sender.Hex(),
		"dstAddress":   receiver.Hex(),
		"srcAmount":    amountInt.String(),
		"dstAmountMin": amountIntMin.String(),
	}

	quote, err := stargate.GetStargateQuotes(httpclient, payload)
	if err != nil {
		return fmt.Errorf("###第 %d 行 获取Stargate-Route请求失败: %v\n", index+1, err)
	}
	utils.Intolog("---获取Stargate-Route成功")

	// 发送 Stargate 跨链交易
	utils.Intolog("---发送Stargate跨链交易中")
	//amount := utils.ParseFloat(amountStr)
	hash, err := stargate.StargateFastBridge(fromChain, sender, amountInt, quote, pk)
	if err != nil {
		return fmt.Errorf("###第 %d 行 发送 stargate 跨链交易失败: %v \n", index+1, err.Error())
	}
	utils.Intolog(fmt.Sprintf("!!!第 %d 行 stargate 跨链交易已发送上链,,fromChain: %s,toChain: %s, Amount: %s hash: %v \n", index+1, fromChain, toChain, amount, hash))
	return nil
}

func getToken(fromChain, toChain, chains string) (string, string, error) {
	fromToken := chain.TokenContracts[fromChain+"_"+chains]
	toToken := chain.TokenContracts[toChain+"_"+chains]
	if fromToken == "" || toToken == "" {
		return "", "", fmt.Errorf("%s 不支持的链，源链 :%s , 目标链: %s", chains, fromChain, toChain)
	}
	return fromToken, toToken, nil
}
