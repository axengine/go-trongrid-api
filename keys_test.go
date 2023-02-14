package go_trongrid_api

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"testing"
)

func TestGenTronKey(t *testing.T) {
	addr, pub, priv, err := GenTronKey()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(addr)
	fmt.Println(pub)
	fmt.Println(priv)
}

func TestValidTronAddress(t *testing.T) {
	if !ValidTronAddress("TRZ9YRNpZwEjCE5ZSU4exwMhKswtKVCmcp") {
		t.Fatal("invalid")
	}
}

func TestAddressConvert(t *testing.T) {
	tron, _ := address.Base58ToAddress("TRZ9YRNpZwEjCE5ZSU4exwMhKswtKVCmcp")
	eth := Tron2ETH(tron)
	fmt.Println(eth.Hex())

	tron = ETH2Tron(eth)
	fmt.Println(tron.String())
}
