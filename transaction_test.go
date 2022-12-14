package go_trongrid_api

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"testing"
)

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
