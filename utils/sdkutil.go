// @Title
// @Description
// @Author  Niels  2020/4/11
package utils

import (
	"github.com/niels1286/nmt/cfg"
	"github.com/niels1286/nuls-go-sdk"
)

func GetOfficalSdk() *nuls.NulsSdk {
	return nuls.NewNulsSdk(cfg.APIServeURL, cfg.PublicSercServeURL, cfg.DefaultChainId)
}
