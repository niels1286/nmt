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
	"github.com/niels1286/nuls-go-sdk"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/spf13/cobra"
	"math/big"
	"strings"
	"time"
)

var to string
var amount float64
var remark string

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Assemble a transfer transaction",
	Long:  `根据参数组装一个转账交易，并返回交易hex`,
	Run: func(cmd *cobra.Command, args []string) {
		tx := txprotocal.Transaction{
			TxType:   txprotocal.TX_TYPE_TRANSFER,
			Time:     uint32(time.Now().Unix()),
			Remark:   []byte(remark),
			Extend:   nil,
			CoinData: nil,
			SignData: nil,
		}
		value := big.NewFloat(amount)
		value = value.Mul(value, big.NewFloat(100000000))
		x, _ := value.Uint64()
		val := new(big.Int)
		val.SetUint64(x)
		fromVal := big.NewInt(100000)
		fromVal.Add(fromVal, val)
		if m < 1 || m > 15 {
			fmt.Println("m value valid")
			return
		}
		pkArray := strings.Split(pks, ",")
		if len(pkArray) < m {
			fmt.Println("Incorrect public keys")
			return
		}
		from := CreateAddress(m, pkArray)
		if "" == from {
			fmt.Println("")
			return
		}
		nonce := GetNonce(from)
		from1 := txprotocal.CoinFrom{
			Coin: txprotocal.Coin{
				Address:       account.AddressStrToBytes(from),
				AssetsChainId: account.NULSChainId,
				AssetsId:      1,
				Amount:        fromVal,
			},
			Nonce:  nonce,
			Locked: 0,
		}
		to1 := txprotocal.CoinTo{
			Coin: txprotocal.Coin{
				Address:       account.AddressStrToBytes(to),
				AssetsChainId: account.NULSChainId,
				AssetsId:      1,
				Amount:        val,
			},
			LockValue: 0,
		}
		coinData := txprotocal.CoinData{
			Froms: []txprotocal.CoinFrom{from1},
			Tos:   []txprotocal.CoinTo{to1},
		}
		var err error
		tx.CoinData, err = coinData.Serialize()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		publicKeys := [][]byte{}
		for _, pk := range pkArray {
			bytes, err := hex.DecodeString(pk)
			if err != nil {
				fmt.Println("public key not right.")
				return
			}
			publicKeys = append(publicKeys, bytes)
		}
		txSign := txprotocal.MultiAddressesSignData{
			M:              uint8(m),
			PubkeyList:     publicKeys,
			CommonSignData: txprotocal.CommonSignData{},
		}
		tx.SignData, err = txSign.Serialize()
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

func GetNonce(address string) []byte {
	sdk := nuls.NewNulsSdk("https://api.nuls.io/jsonrpc", "https://public1.nuls.io/", account.NULSChainId)
	status, err := sdk.GetBalance(address, int(account.NULSChainId), 1)
	if err != nil {
		return nil
	}
	if status == nil {
		return []byte{0, 0, 0, 0, 0, 0, 0, 0}
	}
	return status.Nonce
}

func init() {
	rootCmd.AddCommand(transferCmd)

	transferCmd.Flags().IntVarP(&m, "m", "m", 0, "发起交易的最小签名个数")
	transferCmd.Flags().StringVarP(&pks, "publickeys", "p", "", "多签地址的成员公钥，以','分隔不同的公钥")
	transferCmd.MarkFlagRequired("m")
	transferCmd.MarkFlagRequired("publickeys")

	transferCmd.Flags().StringVarP(&to, "to", "t", "", "转入地址")
	transferCmd.MarkFlagRequired("to")
	transferCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "金额，到账数量，以NULS为单位")
	transferCmd.MarkFlagRequired("amount")
	transferCmd.Flags().StringVarP(&remark, "remark", "r", "", "交易备注，可以为空")
}
