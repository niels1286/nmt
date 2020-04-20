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
	"github.com/niels1286/nmt/cfg"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/tx/txdata"
	"math/big"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var alias string

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "set alias",
	Long:  `Set alias for quick transfer and node name display`,
	Run: func(cmd *cobra.Command, args []string) {
		amount = 1
		to = cfg.BlackHoleAddress

		tx := txprotocal.Transaction{
			TxType:   txprotocal.TX_TYPE_ACCOUNT_ALIAS,
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
				AssetsChainId: cfg.DefaultChainId,
				AssetsId:      1,
				Amount:        fromVal,
			},
			Nonce:  nonce,
			Locked: 0,
		}
		to1 := txprotocal.CoinTo{
			Coin: txprotocal.Coin{
				Address:       account.AddressStrToBytes(to),
				AssetsChainId: cfg.DefaultChainId,
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
		aliasData := txdata.Alias{
			Address: account.AddressStrToBytes(from),
			Alias:   alias,
		}
		tx.Extend, err = aliasData.Serialize()
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
	rootCmd.AddCommand(aliasCmd)
	aliasCmd.Flags().IntVarP(&m, "m", "m", 0, "发起交易的最小签名个数")
	aliasCmd.Flags().StringVarP(&pks, "publickeys", "p", "", "多签地址的成员公钥，以','分隔不同的公钥")
	aliasCmd.MarkFlagRequired("m")
	aliasCmd.MarkFlagRequired("publickeys")
	aliasCmd.Flags().StringVarP(&alias, "alias", "a", "", "别名，只允许小写字母和下划线")
	aliasCmd.MarkFlagRequired("alias")
}
