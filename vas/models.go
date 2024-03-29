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
	"encoding/hex"
	"fmt"
	"github.com/blocktree/openwallet/crypto"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/btcsuite/btcd/txscript"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
	"strings"
)

// ParseHeader 区块链头
func ParseHeader(b *Block) *openwallet.BlockHeader {
	obj := openwallet.BlockHeader{}

	//解析josn
	obj.Merkleroot = b.Merkleroot
	obj.Hash = b.Hash
	obj.Previousblockhash = b.Previousblockhash
	obj.Height = b.Height
	obj.Time = b.Time
	obj.Symbol = Symbol
	return &obj
}

//BlockchainInfo 本地节点区块链信息
type BlockchainInfo struct {
	Chain                string `json:"chain"`
	Blocks               uint64 `json:"blocks"`
	Headers              uint64 `json:"headers"`
	Bestblockhash        string `json:"bestblockhash"`
	Difficulty           string `json:"difficulty"`
	Mediantime           uint64 `json:"mediantime"`
	Verificationprogress string `json:"verificationprogress"`
	Chainwork            string `json:"chainwork"`
	Pruned               bool   `json:"pruned"`
}

func NewBlockchainInfo(json *gjson.Result) *BlockchainInfo {
	b := &BlockchainInfo{}
	//解析json
	b.Chain = gjson.Get(json.Raw, "chain").String()
	b.Blocks = gjson.Get(json.Raw, "blocks").Uint()
	b.Headers = gjson.Get(json.Raw, "headers").Uint()
	b.Bestblockhash = gjson.Get(json.Raw, "bestblockhash").String()
	b.Difficulty = gjson.Get(json.Raw, "difficulty").String()
	b.Mediantime = gjson.Get(json.Raw, "mediantime").Uint()
	b.Verificationprogress = gjson.Get(json.Raw, "verificationprogress").String()
	b.Chainwork = gjson.Get(json.Raw, "chainwork").String()
	b.Pruned = gjson.Get(json.Raw, "pruned").Bool()
	return b
}

//Unspent 未花记录
type Unspent struct {

	/*
			{
		        "txid" : "d54994ece1d11b19785c7248868696250ab195605b469632b7bd68130e880c9a",
		        "vout" : 1,
		        "address" : "mgnucj8nYqdrPFh2JfZSB1NmUThUGnmsqe",
		        "account" : "test label",
		        "scriptPubKey" : "76a9140dfc8bafc8419853b34d5e072ad37d1a5159f58488ac",
		        "amount" : 0.00010000,
		        "confirmations" : 6210,
		        "spendable" : true,
		        "solvable" : true
		    }
	*/
	Key           string `storm:"id"`
	TxID          string `json:"txid"`
	Vout          uint64 `json:"vout"`
	Address       string `json:"address"`
	AccountID     string `json:"account" storm:"index"`
	ScriptPubKey  string `json:"scriptPubKey"`
	Amount        string `json:"amount"`
	Confirmations uint64 `json:"confirmations"`
	Spendable     bool   `json:"spendable"`
	Solvable      bool   `json:"solvable"`
	HDAddress     openwallet.Address
}

func NewUnspent(json *gjson.Result) *Unspent {
	obj := &Unspent{}
	//解析json
	obj.TxID = gjson.Get(json.Raw, "txid").String()
	obj.Vout = gjson.Get(json.Raw, "vout").Uint()
	obj.Address = gjson.Get(json.Raw, "address").String()
	obj.AccountID = gjson.Get(json.Raw, "account").String()
	obj.ScriptPubKey = gjson.Get(json.Raw, "scriptPubKey").String()
	obj.Amount = gjson.Get(json.Raw, "amount").String()
	obj.Confirmations = gjson.Get(json.Raw, "confirmations").Uint()
	//obj.Spendable = gjson.Get(json.Raw, "spendable").Bool()
	obj.Spendable = true
	obj.Solvable = gjson.Get(json.Raw, "solvable").Bool()

	return obj
}

type UnspentSort struct {
	Values     []*Unspent
	Comparator func(a, b *Unspent) int
}

func (s UnspentSort) Len() int {
	return len(s.Values)
}
func (s UnspentSort) Swap(i, j int) {
	s.Values[i], s.Values[j] = s.Values[j], s.Values[i]
}
func (s UnspentSort) Less(i, j int) bool {
	return s.Comparator(s.Values[i], s.Values[j]) < 0
}

//type Address struct {
//	Address   string `json:"address" storm:"id"`
//	Account   string `json:"account" storm:"index"`
//	HDPath    string `json:"hdpath"`
//	CreatedAt time.Time
//}

type User struct {
	UserKey string `storm:"id"`     // primary key
	Group   string `storm:"index"`  // this field will be indexed
	Email   string `storm:"unique"` // this field will be indexed with a unique constraint
	Name    string // this field will not be indexed
	Age     int    `storm:"index"`
}

type Block struct {

	/*

			"hash": "0056d7f09c9e6e3aa437c63f6c5dfc36b71254b450771f50f275c9414720b9d3",
		        "miner": "VUYMX6z5XwBhqfW1mM9vyJEAzJXtfqmBK3",
		        "confirmations": 103,
		        "size": 1352,
		        "height": 42491,
		        "version": 3,
		        "merkleroot": "5a08f932a991dd01e9077b112f070ce02e27384140a1bedd6d5c05e6fe90a866",
		        "tx": [
		            "6ef2a0402e658a5535bf5f92085bfb44de3bb8f062b7b3ac27ed6096fa8577ee",
		            "a8e77558ad5ac44b2913fc1a53094394249ddc41a4d5bb55e1ceba348d7ba5b4",
		            "13f7a6f00c7ca6357f3f26e4bc201004206ace588fb93691463d0571cae7b698",
		            "7c6acfc0cd4400ccfea3a22028bb8f40f1f307a638bd4ad51bec796c21a718fa",
		            "1b8a9835e78465571b726addec38378a0ffab2b43fcb84af0d86c6f64a02566a"
		        ],
		        "time": 1573722270,
		        "nonce": 258,
		        "bits": "2000ffff",
		        "difficulty": 5.96046447753906e-8,
		        "chainwork": "0000000000000000000000000000000000000000000000000000000000a5fc00",
		        "prevblockhash": "0077f8bb8270d2be89979813ed9f673bbc6c84587605a90e00fcba75c3cb8587",
		        "nextblockhash": "003b43eca0e46c7a823bd3b73d90368a47820bc93ecd52cefc30146b742f1042"

	*/

	Hash              string
	Confirmations     uint64
	Merkleroot        string
	tx                []string
	Previousblockhash string
	Height            uint64 `storm:"id"`
	Version           uint64
	Time              uint64
	Fork              bool
	txDetails         []*Transaction
	isVerbose         bool
}

func (wm *WalletManager) NewBlock(json *gjson.Result) *Block {
	obj := &Block{}
	//解析json
	obj.Height = gjson.Get(json.Raw, "height").Uint()
	obj.Hash = gjson.Get(json.Raw, "hash").String()
	obj.Confirmations = gjson.Get(json.Raw, "confirmations").Uint()
	obj.Merkleroot = gjson.Get(json.Raw, "merkleroot").String()
	obj.Previousblockhash = gjson.Get(json.Raw, "prevblockhash").String()
	obj.Version = gjson.Get(json.Raw, "version").Uint()
	obj.Time = gjson.Get(json.Raw, "time").Uint()

	txs := make([]string, 0)
	txDetails := make([]*Transaction, 0)
	for _, tx := range gjson.Get(json.Raw, "tx").Array() {
		if tx.IsObject() {
			obj.isVerbose = true
			txObj := wm.newTxByCore(&tx)
			txObj.BlockHeight = obj.Height
			txObj.BlockHash = obj.Hash
			txObj.Blocktime = int64(obj.Time)
			txDetails = append(txDetails, txObj)
		} else {
			obj.isVerbose = false
			txs = append(txs, tx.String())
		}

	}

	obj.tx = txs
	obj.txDetails = txDetails

	return obj
}

//BlockHeader 区块链头
func (b *Block) BlockHeader(symbol string) *openwallet.BlockHeader {

	obj := openwallet.BlockHeader{}
	//解析json
	obj.Hash = b.Hash
	obj.Confirmations = b.Confirmations
	obj.Merkleroot = b.Merkleroot
	obj.Previousblockhash = b.Previousblockhash
	obj.Height = b.Height
	obj.Version = b.Version
	obj.Time = b.Time
	obj.Symbol = symbol

	return &obj
}

//UnscanRecords 扫描失败的区块及交易
type UnscanRecord struct {
	ID          string `storm:"id"` // primary key
	BlockHeight uint64
	TxID        string
	Reason      string
}

func NewUnscanRecord(height uint64, txID, reason string) *UnscanRecord {
	obj := UnscanRecord{}
	obj.BlockHeight = height
	obj.TxID = txID
	obj.Reason = reason
	obj.ID = common.Bytes2Hex(crypto.SHA256([]byte(fmt.Sprintf("%d_%s", height, txID))))
	return &obj
}

func (wm *WalletManager) newTxByCore(json *gjson.Result) *Transaction {

	/*
		{
			"txid": "6595e0d9f21800849360837b85a7933aeec344a89f5c54cf5db97b79c803c462",
			"hash": "f758cb5181d51f8bee1512b4a862faad5b51c7c85a1a11cd6092ffc1c6649bc5",
			"version": 2,
			"size": 249,
			"vsize": 168,
			"locktime": 1414190,
			"vin": [],
			"vout": [],
			"hex": "02000000000101cc8a3077023c08040e677647ad0e528564764f456b01d8519828df165ab3c4550100000017160014aa59f94152351c79b57b14a53e538a923e332468feffffff02a716167c6f00000017a914a0fe07f130a36d9c7581ccd2886895c049b0cc8287ece29c00000000001976a9148c0bceb59d452b3e077f73a420b8bfe09e0550a788ac0247304402205e667171c1798cde426282bb8bff45901866ad6bf0d209e856c1765eda65ba4802203aaa319ea3de00eccef0006e6ee2089aed4b91ada7953f420a47c9c258d424ca0121033cfda2f93d13b01d46ecc406b03ebaba3e1bd526d2148a0a5d579d52f8c7cf022e941500",
			"blockhash": "0000000040730ea7935cce346ce68bf4c07c10b137ba31960bf8a47c4f7da4ec",
			"confirmations": 20076,
			"time": 1537841342,
			"blocktime": 1537841342
		}
	*/

	obj := Transaction{}
	//解析json
	obj.TxID = gjson.Get(json.Raw, "txid").String()
	obj.Version = gjson.Get(json.Raw, "version").Uint()
	obj.LockTime = gjson.Get(json.Raw, "locktime").Int()
	obj.BlockHash = gjson.Get(json.Raw, "blockhash").String()
	//obj.BlockHeight = gjson.Get(json.Raw, "blockheight").Uint()
	obj.Confirmations = gjson.Get(json.Raw, "confirmations").Uint()
	obj.Blocktime = gjson.Get(json.Raw, "blocktime").Int()
	obj.Size = gjson.Get(json.Raw, "size").Uint()
	//obj.Fees = gjson.Get(json.Raw, "fees").String()
	obj.Decimals = wm.Decimal()
	obj.Vins = make([]*Vin, 0)
	if vins := gjson.Get(json.Raw, "vin"); vins.IsArray() {
		for i, vin := range vins.Array() {
			input := newTxVinByCore(&vin)
			input.N = uint64(i)
			obj.Vins = append(obj.Vins, input)
		}
	}

	obj.Vouts = make([]*Vout, 0)
	if vouts := gjson.Get(json.Raw, "vout"); vouts.IsArray() {
		for _, vout := range vouts.Array() {
			output := newTxVoutByCore(&vout)
			obj.Vouts = append(obj.Vouts, output)
		}
	}

	return &obj
}

func newTxVinByCore(json *gjson.Result) *Vin {

	/*
		{
			"txid": "55c4b35a16df289851d8016b454f766485520ead4776670e04083c0277308acc",
			"vout": 1,
			"scriptSig": {
				"asm": "0014aa59f94152351c79b57b14a53e538a923e332468",
				"hex": "160014aa59f94152351c79b57b14a53e538a923e332468"
			},
			"txinwitness": ["304402205e667171c1798cde426282bb8bff45901866ad6bf0d209e856c1765eda65ba4802203aaa319ea3de00eccef0006e6ee2089aed4b91ada7953f420a47c9c258d424ca01", "033cfda2f93d13b01d46ecc406b03ebaba3e1bd526d2148a0a5d579d52f8c7cf02"],
			"sequence": 4294967294
		}
	*/
	obj := Vin{}
	//解析json
	obj.TxID = gjson.Get(json.Raw, "txid").String()
	obj.Vout = gjson.Get(json.Raw, "vout").Uint()
	obj.Coinbase = gjson.Get(json.Raw, "coinbase").String()
	//obj.Addr = gjson.Get(json.Raw, "addr").String()
	//obj.Value = gjson.Get(json.Raw, "value").String()

	return &obj
}

func newTxVoutByCore(json *gjson.Result) *Vout {

	/*
		{
			"value": 4788.23192231,
			"n": 0,
			"scriptPubKey": {
				"asm": "OP_HASH160 a0fe07f130a36d9c7581ccd2886895c049b0cc82 OP_EQUAL",
				"hex": "a914a0fe07f130a36d9c7581ccd2886895c049b0cc8287",
				"reqSigs": 1,
				"type": "scripthash",
				"addresses": ["2N7vURMwMDjqgijLNFsErFLAWtAg58S6qNv"]
			}
		}
	*/
	obj := Vout{}
	//解析json
	obj.Value = gjson.Get(json.Raw, "value").String()
	obj.N = gjson.Get(json.Raw, "n").Uint()
	obj.ScriptPubKey = gjson.Get(json.Raw, "scriptPubKey.hex").String()

	//提取地址
	if addresses := gjson.Get(json.Raw, "scriptPubKey.addresses"); addresses.IsArray() {
		obj.Addr = addresses.Array()[0].String()
	}

	obj.Type = gjson.Get(json.Raw, "scriptPubKey.type").String()

	//if len(obj.Addr) == 0 {
	//	scriptBytes, _ := hex.DecodeString(obj.ScriptPubKey)
	//	obj.Addr, _ = wm.Decoder.ScriptPubKeyToBech32Address(scriptBytes)
	//}

	return &obj
}

func DecodeScript(script string) ([]byte, error) {
	opcodes := strings.Split(script, " ")
	scriptBuilder := txscript.NewScriptBuilder()
	for _, codeName := range opcodes {
		code, ok := txscript.OpcodeByName[codeName]
		if ok {
			scriptBuilder.AddOp(code)
		} else {
			if len(codeName)%2 != 0 {
				codeName = "0" + codeName
			}
			data, err := hex.DecodeString(codeName)
			if err != nil {
				return nil, err
			}
			scriptBuilder.AddData(data)
		}
	}
	return scriptBuilder.Script()
}
