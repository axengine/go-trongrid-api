package go_trongrid_api

import (
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/pkg/errors"
	"math/big"
)

func (cli *Client) TRC20Transfer(key string, from, to, contract string, amount *big.Int, feeLimit int64) (string, error) {
	tx, err := cli.TRC20Send(from, to, contract, amount, feeLimit)
	if err != nil {
		return "", err
	}
	hash := common.Bytes2Hex(tx.Txid)
	signedTx, err := cli.SignTx(tx.Transaction, key)
	if err != nil {
		return hash, err
	}
	ret, err := cli.Broadcast(signedTx)
	if err != nil {
		return hash, err
	}
	if ret.GetCode() != 0 {
		return hash, errors.Errorf("code:%d msg:%s", ret.GetCode(), ret.GetMessage())
	}

	return hash, nil
}

func (cli *Client) TRC20BalanceOf(addr, contract string) (*big.Int, error) {
	return cli.TRC20ContractBalance(addr, contract)
}
