// @Title
// @Description
// @Author  Niels
package cmd

import "testing"

func TestParseTx(t *testing.T) {
	txHex = "02007a6d915e13e4b8ade59bbde5ad9768616861686168616861007501170100038d5fb45ff2a3053447d25828ba64e21fb89a041301000100a06a0d540200000000000000000000000000000000000000000000000000000008549abf8dc487748d0001000100010000e40b5402000000000000000000000000000000000000000000000000000000000000000000000000"
	parsetxCmd.Run(nil, nil)
}
