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
	"github.com/spf13/cobra"
)

var agentHash string

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "委托",
	Long:  `委托一笔资产到节点上`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(depositCmd)
	depositCmd.Flags().IntVarP(&m, "m", "m", 0, "发起交易的最小签名个数")
	depositCmd.MarkFlagRequired("m")
	depositCmd.Flags().StringVarP(&pks, "publickeys", "p", "", "多签地址的成员公钥，以','分隔不同的公钥")
	depositCmd.MarkFlagRequired("publickeys")
	depositCmd.Flags().StringVarP(&pks, "agenthash", "h", "", "节点位移标识")
	depositCmd.MarkFlagRequired("publickeys")
	depositCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "委托金额")
	depositCmd.MarkFlagRequired("amount")

}
