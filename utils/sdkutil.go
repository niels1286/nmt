// @Title
// @Description
// @Author  Niels  2020/4/11
package utils

import (
	"github.com/niels1286/nuls-go-sdk"
	"github.com/niels1286/nuls-go-sdk/account"
)

func GetOfficalSdk() *nuls.NulsSdk {
	return nuls.NewNulsSdk("https://api.nuls.io/jsonrpc/", "https://public1.nuls.io", account.NULSChainId)
}
