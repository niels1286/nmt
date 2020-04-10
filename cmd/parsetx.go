/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/tx/txdata"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"github.com/spf13/cobra"
)

var txHex string

type TxInfo struct {
	Hash     string
	TxType   string
	TxData   map[string]string
	CoinData string
}

var TypeMap = map[int]string{
	1:                                 "共识奖励",
	2:                                 "转账交易",
	5:                                 "委托交易",
	txprotocal.TX_TYPE_ACCOUNT_ALIAS:  "设置别名",
	txprotocal.TX_TYPE_CANCEL_DEPOSIT: "退出委托",
	txprotocal.TX_TYPE_STOP_AGENT:     "停止节点",
	txprotocal.TX_TYPE_REGISTER_AGENT: "创建节点",
}

func (ti *TxInfo) String() string {
	bus := "TxExtend:\n"
	for key, val := range ti.TxData {
		bus += "\t" + key + " : " + val + "\n"
	}
	value := fmt.Sprintf("===========tx info============\nhash:%s\ntype:%s\n%s%s", ti.Hash, ti.TxType, bus, ti.CoinData)
	return value
}

// parsetxCmd represents the parsetx command
var parsetxCmd = &cobra.Command{
	Use:   "parsetx",
	Short: "Deserialize transactions to readable content",
	Long:  `Deserialize the transaction into readable content. Mainly focus on transaction type, coindata content and txdata content.`,
	Run: func(cmd *cobra.Command, args []string) {
		if "" == txHex {
			fmt.Println("txHex is valid.")
			return
		}
		txBytes, err := hex.DecodeString(txHex)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tx := txprotocal.ParseTransactionByReader(seria.NewByteBufReader(txBytes, 0))
		fmt.Println(tx)

		info := getTxInfo(tx)
		fmt.Println(info.String())
	},
}

func getTxInfo(tx *txprotocal.Transaction) *TxInfo {
	typeStr := TypeMap[int(tx.TxType)]
	txData := map[string]string{}
	switch tx.TxType {
	case txprotocal.TX_TYPE_DEPOSIT:
		deposit := &txdata.Deposit{}
		deposit.Parse(seria.NewByteBufReader(tx.Extend, 0))
		txData["address"] = account.GetStringAddress(deposit.Address, account.NULSPrefix)
		txData["agentHash"] = deposit.AgentHash.String()
		txData["amount"] = fmt.Sprintf("%d", deposit.Amount.Uint64()/100000000)
	case txprotocal.TX_TYPE_REGISTER_AGENT:
		agent := &txdata.Agent{}
		agent.Parse(seria.NewByteBufReader(tx.Extend, 0))
		txData["agentAddress"] = account.GetStringAddress(agent.AgentAddress, account.NULSPrefix)
		txData["packingAddress"] = account.GetStringAddress(agent.PackingAddress, account.NULSPrefix)
		txData["rewardAddress"] = account.GetStringAddress(agent.RewardAddress, account.NULSPrefix)
		txData["amount"] = fmt.Sprintf("%d", agent.Amount.Uint64()/100000000)
		txData["commissionRate"] = fmt.Sprintf("%d%", agent.CommissionRate)
	case txprotocal.TX_TYPE_STOP_AGENT:
		info := &txdata.StopAgent{}
		info.Parse(seria.NewByteBufReader(tx.Extend, 0))
		txData["agentHash"] = info.AgentHash.String()
	case txprotocal.TX_TYPE_CANCEL_DEPOSIT:
		info := txdata.Withdraw{}
		info.Parse(seria.NewByteBufReader(tx.Extend, 0))
		txData["depositTxHash"] = info.DepositTxHash.String()
	case txprotocal.TX_TYPE_ACCOUNT_ALIAS:
		info := &txdata.Alias{}
		info.Parse(seria.NewByteBufReader(tx.Extend, 0))
		txData["address"] = account.GetStringAddress(info.Address, account.NULSPrefix)
		txData["alias"] = info.Alias
	default:
		if tx.Extend != nil {
			txData["hex"] = hex.EncodeToString(tx.Extend)
		}
	}
	cd := &txprotocal.CoinData{}
	cd.Parse(seria.NewByteBufReader(tx.CoinData, 0))
	var message = "From:\n"
	for _, from := range cd.Froms {
		message += "\t" + account.GetStringAddress(from.Address, account.NULSPrefix) + "(" + fmt.Sprintf("%d", from.AssetsChainId) + "-" + fmt.Sprintf("%d", from.AssetsChainId) + ") :: " + mathutils.GetStringAmount(from.Amount, 8) + "\n"
	}
	message += "To:\n"
	for _, to := range cd.Tos {
		lock := fmt.Sprintf("%d", to.LockValue)
		if to.LockValue == uint64(18446744073709551615) {
			lock = "-1"
		}
		message += "\t" + account.GetStringAddress(to.Address, account.NULSPrefix) + "(" + fmt.Sprintf("%d", to.AssetsChainId) + "-" + fmt.Sprintf("%d", to.AssetsChainId) + ") :: " + mathutils.GetStringAmount(to.Amount, 8) + " (lock:" + lock + ")\n"
	}

	return &TxInfo{
		Hash:     tx.GetHash().String(),
		TxType:   typeStr,
		TxData:   txData,
		CoinData: message,
	}
}

func init() {
	rootCmd.AddCommand(parsetxCmd)
	parsetxCmd.Flags().StringVarP(&txHex, "txhex", "t", "", "Transaction serialization data in hexadecimal string format")
	parsetxCmd.MarkFlagRequired("txhex")
}
