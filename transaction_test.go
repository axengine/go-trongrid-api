package go_trongrid_api

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"testing"
)

func TestGetAddress(t *testing.T) {
	account, err := conn.GetAccountResource("TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG")
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println(address.Address(account.GetAddress()))
	JsonPretty(account)
}

func TestGetTransactionInfoByID(t *testing.T) {
	tx, err := conn.GetTransactionInfoByID("9926803bb63a4d5598c6d2c645c16d3456efc13a6f7a15873d11d179e97e218d")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tx.Result)
	JsonPretty(tx)
}

func TestTransferTRX(t *testing.T) {
	tx, err := conn.Transfer("TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG", "TV9uEbtAmAW8e9sdsa8k6aZ2ziz9siGqBj", 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(common.Bytes2Hex(tx.GetTxid()))
	signedTx, err := conn.SignTx(tx.Transaction, "--")
	if err != nil {
		t.Fatal(err)
	}

	ret, err := conn.Broadcast(signedTx)
	if err != nil {
		t.Fatal(err)
	}
	JsonPretty(ret)
}
