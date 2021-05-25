package blockchain

import (
	"bytes"
	"github.com/Lozovoi-Rodion/golang-blockchain/wallet"
)

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

func NewTxOutput(value int, address string) *TxOutput {
	txo := &TxOutput{
		Value:      value,
		PubKeyHash: nil,
	}
	txo.Lock([]byte(address))

	return txo
}

func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4] // remove version byte and checksum
	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) isLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}
