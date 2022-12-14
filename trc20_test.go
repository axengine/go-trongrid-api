package go_trongrid_api

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/fatih/structs"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/ptypes"
	"math/big"
	"strings"
	"testing"
)

var abiIns abi.ABI

func TestParseTRC20Transfer(t *testing.T) {
	abiIns, _ = abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"}]`))

	block, err := conn.GetBlockByNum(29589966)
	if err != nil {
		t.Fatal(err)
	}
	for _, tx := range block.Transactions {
		JsonPretty(tx)
		if tx.Result.GetCode() != api.Return_SUCCESS {
			continue
		}
		if tx.Transaction.RawData != nil {
			for _, contract := range tx.Transaction.RawData.Contract {
				if contract.GetType() == core.Transaction_Contract_TriggerSmartContract {
					var c core.TriggerSmartContract
					if err = ptypes.UnmarshalAny(contract.GetParameter(), &c); err != nil {
						t.Fatal(err)
					}
					fmt.Println(address.Address(c.OwnerAddress).String())
					fmt.Println(address.Address(c.ContractAddress).String())
					JsonPretty(structs.Map(c))

					methodId := new(big.Int).SetBytes(c.Data[:4]).Int64()
					if methodId != 0xa9059cbb { //transfer
						continue
					}
					unpack, err := abiIns.MethodById(c.Data[:4])
					if err != nil {
						t.Fatal(err)
					}
					params := make(map[string]interface{})
					inputs := unpack.Inputs
					err = inputs.UnpackIntoMap(params, c.Data[4:])
					if err != nil {
						t.Fatal(err)
					}
					fmt.Println(params)
				}
			}
		}

	}
	//JsonPretty(block)
}

func TestGetTRC20Bal(t *testing.T) {
	bal, err := conn.TRC20BalanceOf("TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG", "TGT6r3rS1tyA77B8LT1gE1pfMFxK9f2XJT")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bal)
}

func TestTRC20Transfer(t *testing.T) {
	hash, err := conn.TRC20Transfer("--",
		"TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG",
		"TV9uEbtAmAW8e9sdsa8k6aZ2ziz9siGqBj",
		"TGT6r3rS1tyA77B8LT1gE1pfMFxK9f2XJT",
		new(big.Int).Mul(big.NewInt(12), big.NewInt(1e17)),
		10000000)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hash)
}
