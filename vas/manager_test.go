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
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/log"
	"github.com/codeskyblue/go-sh"
	"github.com/shopspring/decimal"
	"path/filepath"
	"testing"
)

var (
	tw *WalletManager
)

func init() {

	tw = testNewWalletManager()
}

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", "VAS.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return nil
	}
	wm.LoadAssetsConfig(c)
	//wm.ExplorerClient.Debug = false
	wm.WalletClient.Debug = true
	return wm
}

func TestWalletManager(t *testing.T) {

	t.Log("Symbol:", tw.Config.Symbol)
	t.Log("ServerAPI:", tw.Config.ServerAPI)
}

//func TestImportPrivKey(t *testing.T) {
//
//	tests := []struct {
//		seed []byte
//		name string
//		tag  string
//	}{
//		{
//			seed: tw.GenerateSeed(),
//			name: "Chance",
//			tag:  "first",
//		},
//		{
//			seed: tw.GenerateSeed(),
//			name: "Chance",
//			tag:  "second",
//		},
//	}
//
//	for i, test := range tests {
//		key, err := keystore.NewHDKey(test.seed, test.name, "m/44'/88'")
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		privateKey, err := key.MasterKey.ECPrivKey()
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		publicKey, err := key.MasterKey.ECPubKey()
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		wif, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true)
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		t.Logf("Privatekey wif[%d] = %s\n", i, wif.String())
//
//		address, err := btcutil.NewAddressPubKey(publicKey.SerializeCompressed(), &chaincfg.MainNetParams)
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//			continue
//		}
//
//		t.Logf("Privatekey address[%d] = %s\n", i, address.EncodeAddress())
//
//		//解锁钱包
//		err = tw.UnlockWallet("1234qwer", 120)
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//		}
//
//		//导入私钥
//		err = tw.ImportPrivKey(wif.String(), test.name)
//		if err != nil {
//			t.Errorf("ImportPrivKey[%d] failed unexpected error: %v\n", i, err)
//		} else {
//			t.Logf("ImportPrivKey[%d] success \n", i)
//		}
//	}
//
//}

func TestListAddress(t *testing.T) {
	addresses, err := tw.ListAddressesRPC()
	if err != nil {
		t.Errorf("ListAddress failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addresses {
		t.Logf("ListAddress address[%d] = %s\n", i, a)
	}
}

//func TestCreateBatchPrivateKey(t *testing.T) {
//
//	w, err := tw.GetWalletInfo("Zhiquan Test")
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	key, err := w.HDKey("1234qwer")
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	wifs, err := tw.CreateBatchPrivateKey(key, 10000)
//	if err != nil {
//		t.Errorf("CreateBatchPrivateKey failed unexpected error: %v\n", err)
//		return
//	}
//
//	for i, wif := range wifs {
//		t.Logf("CreateBatchPrivateKey[%d] wif = %v \n", i, wif)
//	}
//
//}

//func TestImportMulti(t *testing.T) {
//
//	addresses := []string{
//		"1CoRcQGjPEyWmB1ZyG6CEDN3SaMsaD3ERa",
//		"1ESGCsXkNr3h5wvWScdCpVHu2GP3KJtCdV",
//	}
//
//	keys := []string{
//		"L5k8VYSvuZxC5FCczGVC8MmnKKix3Mcs6t185eUJVKTzZb1f6bsX",
//		"L3RVDjPVBSc7DD4WtmzbHkAHJW4kDbyXbw4vBppZ4DRtPt5u8Naf",
//	}
//
//	UnlockWallet("1234qwer", 120)
//	failed, err := ImportMulti(addresses, keys, "Zhiquan Test")
//	if err != nil {
//		t.Errorf("ImportMulti failed unexpected error: %v\n", err)
//	} else {
//		t.Errorf("ImportMulti result: %v\n", failed)
//	}
//}

func TestGOSH(t *testing.T) {
	//text, err := sh.Command("go", "env").Output()
	//text, err := sh.Command("wmd", "version").Output()
	text, err := sh.Command("wmd", "Config", "see", "-s", "btm").Output()
	if err != nil {
		t.Errorf("GOSH failed unexpected error: %v\n", err)
	} else {
		t.Errorf("GOSH output: %v\n", string(text))
	}
}

func TestListUnspent(t *testing.T) {
	utxos, err := tw.ListUnspent(0, "VSaJg2ARstrpqh6GdwfMZF1xBY25xnPEBV")
	if err != nil {
		t.Errorf("ListUnspent failed unexpected error: %v\n", err)
		return
	}
	totalBalance := decimal.Zero
	for _, u := range utxos {
		t.Logf("ListUnspent %s: %s = %s\n", u.Address, u.AccountID, u.Amount)
		amount, _ := decimal.NewFromString(u.Amount)
		totalBalance = totalBalance.Add(amount)
	}

	t.Logf("totalBalance: %s \n", totalBalance.String())
}

func TestEstimateFee(t *testing.T) {
	feeRate, _ := tw.EstimateFeeRate()
	t.Logf("EstimateFee feeRate = %s\n", feeRate.StringFixed(8))
	fees, _ := tw.EstimateFee(10, 2, feeRate)
	t.Logf("EstimateFee fees = %s\n", fees.StringFixed(8))
}

func TestWalletManager_ImportAddress(t *testing.T) {
	addr := "Ga2thK76EF4Y1q14RtmCfBZepC2GYBvaCy"
	err := tw.ImportAddress(addr, "")
	if err != nil {
		t.Errorf("RestoreWallet failed unexpected error: %v\n", err)
		return
	}
	log.Info("imported success")
}

func TestWalletManager_ListAddresses(t *testing.T) {
	addresses, err := tw.ListAddresses()
	if err != nil {
		t.Errorf("GetAddressesByAccount failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addresses {
		t.Logf("ListAddresses address[%d] = %s\n", i, a)
	}
}

func TestWalletManager_GetInfo(t *testing.T) {
	tw.GetInfo()
}