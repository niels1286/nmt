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
	"fmt"
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/spf13/cobra"
	"strings"
)

var m int
var pks string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a multi address",
	Long:  `create a multi address`,
	Run: func(cmd *cobra.Command, args []string) {
		if m < 1 || m > 15 {
			fmt.Println("m value valid")
			return
		}
		pkArray := strings.Split(pks, ",")
		if len(pkArray) < m {
			fmt.Println("Incorrect public keys")
			return
		}
		address := account.CreateMultiAddress(account.NULSChainId, uint8(m), pkArray, account.NULSPrefix)
		fmt.Println("Operation Successed.\naddress:", address)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().IntVarP(&m, "m", "m", 0, "发起交易的最小签名个数")
	createCmd.Flags().StringVarP(&pks, "publickeys", "p", "", "多签地址的成员公钥，以','分隔不同的公钥")
	createCmd.MarkFlagRequired("m")
	createCmd.MarkFlagRequired("publickeys")

}
