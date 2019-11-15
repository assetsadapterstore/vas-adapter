package openwtester

import (
	"github.com/assetsadapterstore/vas-adapter/vas"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(vas.Symbol, vas.NewWalletManager())
}
