// @Title
// @Description
// @Author  Niels  2020/4/11
package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/niels1286/nmt/cfg"
	"github.com/niels1286/nuls-go-sdk"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"math/big"
	"time"
)

func GetOfficalSdk() *nuls.NulsSdk {
	return nuls.NewNulsSdk(cfg.APIServeURL, cfg.PublicSercServeURL, cfg.DefaultChainId)
}

func AssembleTransferTx(m int, pkArray []string, amount float64, remark string, to string, fromLocked byte, toLockValue uint64, nonce []byte) *txprotocal.Transaction {
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
		return nil
	}

	from := CreateAddress(m, pkArray)
	if "" == from {
		fmt.Println("")
		return nil
	}
	if nonce == nil {
		nonce = GetNonce(from)
	}
	from1 := txprotocal.CoinFrom{
		Coin: txprotocal.Coin{
			Address:       account.AddressStrToBytes(from),
			AssetsChainId: cfg.DefaultChainId,
			AssetsId:      1,
			Amount:        fromVal,
		},
		Nonce:  nonce,
		Locked: fromLocked,
	}
	to1 := txprotocal.CoinTo{
		Coin: txprotocal.Coin{
			Address:       account.AddressStrToBytes(to),
			AssetsChainId: cfg.DefaultChainId,
			AssetsId:      1,
			Amount:        val,
		},
		LockValue: toLockValue,
	}
	coinData := txprotocal.CoinData{
		Froms: []txprotocal.CoinFrom{from1},
		Tos:   []txprotocal.CoinTo{to1},
	}
	var err error
	tx.CoinData, err = coinData.Serialize()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	publicKeys := [][]byte{}
	for _, pk := range pkArray {
		bytes, err := hex.DecodeString(pk)
		if err != nil {
			fmt.Println("public key not right.")
			return nil
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
		return nil
	}
	return &tx
}

func GetNonce(address string) []byte {
	sdk := GetOfficalSdk()
	status, err := sdk.GetBalance(address, int(cfg.DefaultChainId), 1)
	if err != nil {
		return nil
	}
	if status == nil {
		return []byte{0, 0, 0, 0, 0, 0, 0, 0}
	}
	return status.Nonce
}

func CreateAddress(m int, pks []string) string {
	address := account.CreateMultiAddress(cfg.DefaultChainId, uint8(m), pks, cfg.DefaultAddressPrefix)
	return address
}
