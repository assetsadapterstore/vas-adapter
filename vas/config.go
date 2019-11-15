/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package vas

import (
	"github.com/assetsadapterstore/vas-adapter/vasTransaction"
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/common/file"
	"github.com/shopspring/decimal"
	"path/filepath"
	"strings"
	"time"
)

const (
	//币种
	Symbol    = "VAS"
	CurveType = owcrypt.ECC_CURVE_SECP256K1
	Decimals  = int32(8)
)

var (
	MainNetAddressPrefix = vasTransaction.AddressPrefix{P2PKHPrefix: []byte{0x46}, P2WPKHPrefix: []byte{0x05}, P2SHPrefix: nil, Bech32Prefix: "vas"}
	TestNetAddressPrefix = vasTransaction.AddressPrefix{P2PKHPrefix: []byte{0x46}, P2WPKHPrefix: []byte{0x05}, P2SHPrefix: nil, Bech32Prefix: "vas"}
)

type WalletConfig struct {

	//币种
	Symbol    string
	MasterKey string

	//RPC认证账户名
	RpcUser string
	//RPC认证账户密码
	RpcPassword string
	//证书目录
	CertsDir string
	//钥匙备份路径
	keyDir string
	//地址导出路径
	addressDir string
	//配置文件路径
	configFilePath string
	//配置文件名
	configFileName string
	//rpc证书
	CertFileName string
	//区块链数据文件
	BlockchainFile string
	//是否测试网络
	IsTestNet bool
	// 核心钱包是否只做监听
	CoreWalletWatchOnly bool
	//最大的输入数量
	MaxTxInputs int
	//本地数据库文件路径
	DBPath string
	//备份路径
	backupDir string
	//钱包服务API
	ServerAPI string
	//钱包安装的路径
	NodeInstallPath string
	//钱包数据文件目录
	WalletDataPath string
	//汇总阀值
	Threshold decimal.Decimal
	//汇总地址
	SumAddress string
	//汇总执行间隔时间
	CycleSeconds time.Duration
	//默认配置内容
	DefaultConfig string
	//曲线类型
	CurveType uint32
	//核心钱包密码，配置有值用于自动解锁钱包
	WalletPassword string
	//后台数据源类型
	RPCServerType int
	//s是否支持隔离验证
	SupportSegWit bool
	//Omni代币转账最低成本
	OmniTransferCost string
	//OmniCore API
	OmniCoreAPI string
	//Omni rpc user
	OmniRPCUser string
	//Omni rpc password
	OmniRPCPassword string
	//是否支持omni
	OmniSupport bool
	//主网地址前缀
	MainNetAddressPrefix vasTransaction.AddressPrefix
	//测试网地址前缀
	TestNetAddressPrefix vasTransaction.AddressPrefix
	//小数位精度
	Decimals int32
	//最低手续费
	MinFees decimal.Decimal
	//数据目录
	DataDir string
}

func NewConfig(symbol string, curveType uint32, decimals int32) *WalletConfig {

	c := WalletConfig{}

	//币种
	c.Symbol = symbol
	c.CurveType = curveType

	//RPC认证账户名
	c.RpcUser = ""
	//RPC认证账户密码
	c.RpcPassword = ""
	//证书目录
	c.CertsDir = filepath.Join("data", strings.ToLower(c.Symbol), "certs")
	//钥匙备份路径
	c.keyDir = filepath.Join("data", strings.ToLower(c.Symbol), "key")
	//地址导出路径
	c.addressDir = filepath.Join("data", strings.ToLower(c.Symbol), "address")
	//区块链数据
	//blockchainDir = filepath.Join("data", strings.ToLower(Symbol), "blockchain")
	//配置文件路径
	c.configFilePath = filepath.Join("conf")
	//配置文件名
	c.configFileName = c.Symbol + ".ini"
	//rpc证书
	c.CertFileName = "rpc.cert"
	//区块链数据文件
	c.BlockchainFile = "blockchain.db"
	//是否测试网络
	c.IsTestNet = true
	// 核心钱包是否只做监听
	c.CoreWalletWatchOnly = true
	//最大的输入数量
	c.MaxTxInputs = 50
	//本地数据库文件路径
	c.DBPath = filepath.Join("data", strings.ToLower(c.Symbol), "db")
	//备份路径
	c.backupDir = filepath.Join("data", strings.ToLower(c.Symbol), "backup")
	//钱包服务API
	c.ServerAPI = "http://127.0.0.1:10000"
	//钱包安装的路径
	c.NodeInstallPath = ""
	//钱包数据文件目录
	c.WalletDataPath = ""
	//汇总阀值
	c.Threshold = decimal.NewFromFloat(5)
	//汇总地址
	c.SumAddress = ""
	//汇总执行间隔时间
	c.CycleSeconds = time.Second * 10
	//核心钱包密码，配置有值用于自动解锁钱包
	c.WalletPassword = ""
	//后台数据源类型
	c.RPCServerType = 0
	//支持隔离见证
	c.SupportSegWit = true
	//是否支持omni
	c.OmniSupport = false
	//小数位精度
	c.Decimals = decimals
	//最低手续费
	c.MinFees = decimal.Zero
	c.MainNetAddressPrefix = MainNetAddressPrefix
	c.TestNetAddressPrefix = TestNetAddressPrefix

	return &c
}

//创建文件夹
func (wc *WalletConfig) makeDataDir() {

	if len(wc.DataDir) == 0 {
		//默认路径当前文件夹./data
		wc.DataDir = "data"
	}

	//本地数据库文件路径
	wc.DBPath = filepath.Join(wc.DataDir, strings.ToLower(wc.Symbol), "db")

	//创建目录
	file.MkdirAll(wc.DBPath)
}
