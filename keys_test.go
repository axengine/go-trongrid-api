package go_trongrid_api

import (
	"fmt"
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
