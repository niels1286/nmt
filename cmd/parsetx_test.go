// @Title
// @Description
// @Author  Niels
package cmd

import "testing"

func TestParseTx(t *testing.T) {
	txHex = "0500e3aebc5e00574066a67bb829000000000000000000000000000000000000000000000000000001000384dfe765866a220db30753c66c6cf51e48e52ba4a92ac22f4e3686a594be6558f46400d889b3832fe776fbf5e705307005fa865a8c011701000384dfe765866a220db30753c66c6cf51e48e52ba4010001008073a97bb829000000000000000000000000000000000000000000000000000008000000000000000000011701000384dfe765866a220db30753c66c6cf51e48e52ba4010001004066a67bb8290000000000000000000000000000000000000000000000000000fffffffffffffffffde70103052103de80653c218a69b3f89d41f82797d6847afb3c217944edad21199ee6a3f831de20046ff585fca5282eb9c7479967890bb6cc5ff6ea2bce667283ce97a21227be622103b736b7ac6cd2cfe8c221926f8fc0e2b84f8cb96c079c800be09f682d8805097e2103a1f65c80936606df6185fe9bd808d7dd5201e1e88f2a475f6b2a70d81f7f52e421021226ecc609a1a054b760d30a8e1715524ddf22274b152a7f8a98b9b7da81d83321021226ecc609a1a054b760d30a8e1715524ddf22274b152a7f8a98b9b7da81d833463044022056f80aab8ea742b9136ff69c5f0b5340472f889532d7511a0da99c599e2a501d0220395a649c1ee57cd3daf4494e0400d3b98bcb8f66301ebf9d31ee065f87e0ac9c2103de80653c218a69b3f89d41f82797d6847afb3c217944edad21199ee6a3f831de4730450221009b8b5b450cf82f84b94e02d656c38b6591b716c610ba1c5fbc2c34b4ba60a35d02202f916d0df7fb9016b3713eba82fb25920b2ebb1d1fc379cbba49d89a07d1530b2103b736b7ac6cd2cfe8c221926f8fc0e2b84f8cb96c079c800be09f682d8805097e46304402206fa07e556e2d88bb9cbe33efb4f78081f965823401d0c2dc68f2a95a09bfdae5022043c6945c23a7927f77ee19c0dc853aa2659fac864216b3fe821f00707d19e974"
	parsetxCmd.Run(nil, nil)
}
