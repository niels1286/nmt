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
	"encoding/json"
	"fmt"
	"github.com/niels1286/nmt/utils"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/tx/txdata"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

var depositTxHash string

// withdrawCmd represents the withdraw command
var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "退出委托",
	Long:  `退出指定一笔委托，立即解锁对应的资产`,
	Run: func(cmd *cobra.Command, args []string) {
		hashBytes, err := hex.DecodeString(depositTxHash)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		ahash := txprotocal.NewNulsHash(hashBytes)

		txJson, err := utils.GetOfficalSdk().GetTxJson(ahash)
		if err != nil {
			fmt.Println("Can't find the deposit transaction.")
			return
		}
		//fmt.Println(txJson)
		txmap := map[string]interface{}{}
		json.Unmarshal([]byte(txJson), &txmap)
		txDataHex := txmap["txDataHex"].(string)
		if txDataHex == "" {
			fmt.Println("Failed to parse the deposit transaction.")
			return
		}
		txDataBytes, err := hex.DecodeString(txDataHex)
		if err != nil {
			fmt.Println("Failed to parse the deposit transaction.")
			return
		}
		depositData := txdata.Deposit{}
		depositData.Parse(seria.NewByteBufReader(txDataBytes, 0))
		value := depositData.Amount.Div(depositData.Amount, big.NewInt(100000000))
		amount = float64(value.Uint64()) - 0.001
		pkArray := strings.Split(pks, ",")
		if len(pkArray) < m {
			fmt.Println("Incorrect public keys")
			return
		}
		address := utils.CreateAddress(m, pkArray)

		tx := utils.AssembleTransferTx(m, pkArray, amount, "", address, 255, 0, hashBytes[24:])
		if tx == nil {
			return
		}
		tx.TxType = txprotocal.TX_TYPE_CANCEL_DEPOSIT

		withdrawData := txdata.Withdraw{DepositTxHash: ahash}

		tx.Extend, err = withdrawData.Serialize()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		txBytes, err := tx.Serialize()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		txHex := hex.EncodeToString(txBytes)

		fmt.Println("Successed:\ntxHex : " + txHex)
		fmt.Println("txHash : " + tx.GetHash().String())
	},
}

func init() {
	rootCmd.AddCommand(withdrawCmd)
	withdrawCmd.Flags().IntVarP(&m, "m", "m", 0, "发起交易的最小签名个数")
	withdrawCmd.MarkFlagRequired("m")
	withdrawCmd.Flags().StringVarP(&pks, "publickeys", "p", "", "多签地址的成员公钥，以','分隔不同的公钥")
	withdrawCmd.MarkFlagRequired("publickeys")
	withdrawCmd.Flags().StringVarP(&depositTxHash, "depositTxHash", "h", "", "委托交易的交易hash")
	withdrawCmd.MarkFlagRequired("depositTxHash")
}
