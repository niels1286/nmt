// @Title
// @Description
// @Author  Niels  2020/4/23
package cmd

import (
	"github.com/niels1286/nmt/cfg"
	"github.com/niels1286/nuls-go-sdk/account"
	"testing"
)

func TestDeposit(t *testing.T) {

	ap := "3e73f764492e95362cf325bd7168d145110a75e447510c927612586c06b23e91"
	bp := "6d10f3aa23018de6bc7d1ee52badd696f0db56082c62826ba822978fdf3a59fa"
	cp := "f7bb391ab82ba9ec7a552955b2fe50d79eea085d7571e5e2480d1777bc171f5e"

	a, _ := account.GetAccountFromPrkey(ap, cfg.DefaultChainId, cfg.DefaultAddressPrefix)
	b, _ := account.GetAccountFromPrkey(bp, cfg.DefaultChainId, cfg.DefaultAddressPrefix)
	c, _ := account.GetAccountFromPrkey(cp, cfg.DefaultChainId, cfg.DefaultAddressPrefix)

	m = 2
	pks = a.GetPubKeyHex(true) + "," + b.GetPubKeyHex(true) + "," + c.GetPubKeyHex(true)
	agentHash = "fd8ffdf4fdda19db761ddd5c5e7ecc1d9f7be540706b1f5d22353683ac585b9e"
	amount = 2001

	depositCmd.Run(nil, nil)
}
