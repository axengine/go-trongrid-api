package go_trongrid_api

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func (cli *Client) SignTx(tx *core.Transaction, key string) (*core.Transaction, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)
	return tx, nil
}

func (cli *Client) SendTx(signedTx *core.Transaction) error {
	_, err := cli.Broadcast(signedTx)
	return err
}

func (cli *Client) SendTRX(key string, from, toAddress string, amount int64) (string, error) {
	tx, err := cli.Transfer(from, toAddress, amount)
	if err != nil {
		return "", errors.Wrap(err, "Transfer")
	}
	hash := common.Bytes2Hex(tx.Txid)
	signedTx, err := cli.SignTx(tx.Transaction, key)
	if err != nil {
		return hash, errors.Wrap(err, "SignTx")
	}
	ret, err := cli.Broadcast(signedTx)
	if err != nil {
		return hash, errors.Wrap(err, "Broadcast")
	}
	if ret.GetCode() != 0 {
		return hash, errors.Errorf("code:%d msg:%s", ret.GetCode(), ret.GetMessage())
	}
	return hash, nil
}
