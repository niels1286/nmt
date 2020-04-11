// @Title
// @Description
// @Author  Niels  2020/4/11
package cmd

import "testing"

func TestTransfer(t *testing.T) {
	from = "NULSd6HhA3cJftK4YZhRGM4DfsgZhUkFKkF4j"
	to = "NULSd6HhA3cJftK4YZhRGM4DfsgZhUkFKkF4j"
	amount = 100
	remark = "中国字hahahahaha"
	transferCmd.Run(nil, nil)

}
