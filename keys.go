package go_trongrid_api

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
)

// GenTronKey 生成波场地址信息
// 返回:地址(Base58),公钥(hex),私钥(hex),error
func GenTronKey() (string, string, string, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return "", "", "", err
	}

	buff := make([]byte, 32)
	copy(buff[32-len(priv.D.Bytes()):], priv.D.Bytes())
	return address.PubkeyToAddress(priv.PublicKey).String(),
		common.Bytes2Hex(crypto.CompressPubkey(&priv.PublicKey)),
		common.Bytes2Hex(buff),
		err
}

// ValidTronAddress 校验是否合法的Tron地址(Base58)
func ValidTronAddress(addr string) bool {
	_, err := common.DecodeCheck(addr)
	if err != nil {
		return false
	}
	return true
}

func ETH2Tron(ethAddr ethcmn.Address) address.Address {
	addressTron := make([]byte, 0)
	addressTron = append(addressTron, byte(0x41))
	addressTron = append(addressTron, ethAddr[:]...)
	return addressTron
}

func Tron2ETH(addressTron address.Address) ethcmn.Address {
	var addressETH ethcmn.Address
	copy(addressETH[:], addressTron[1:])
	return addressETH
}
