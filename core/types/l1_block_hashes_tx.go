package types

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
)

// payload, RLP encoded
type L1BlockHashesTx struct {
	Gas    uint64
	To     *common.Address // can not be nil, we do not allow contract creation from L1
	Value  *big.Int
	Data   []byte
	Sender common.Address
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *L1BlockHashesTx) copy() TxData {
	cpy := &L1BlockHashesTx{
		Gas:    tx.Gas,
		To:     copyAddressPtr(tx.To),
		Value:  new(big.Int),
		Data:   common.CopyBytes(tx.Data),
		Sender: tx.Sender,
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	return cpy
}

// accessors for innerTx.
func (tx *L1BlockHashesTx) txType() byte           { return L1BlockHashesTxType }
func (tx *L1BlockHashesTx) chainID() *big.Int      { return common.Big0 }
func (tx *L1BlockHashesTx) accessList() AccessList { return nil }
func (tx *L1BlockHashesTx) data() []byte           { return tx.Data }
func (tx *L1BlockHashesTx) gas() uint64            { return tx.Gas }
func (tx *L1BlockHashesTx) gasFeeCap() *big.Int    { return new(big.Int) }
func (tx *L1BlockHashesTx) gasTipCap() *big.Int    { return new(big.Int) }
func (tx *L1BlockHashesTx) gasPrice() *big.Int     { return new(big.Int) }
func (tx *L1BlockHashesTx) value() *big.Int        { return tx.Value }
func (tx *L1BlockHashesTx) nonce() uint64          { return 0 }
func (tx *L1BlockHashesTx) to() *common.Address    { return tx.To }

func (tx *L1BlockHashesTx) rawSignatureValues() (v, r, s *big.Int) {
	return common.Big0, common.Big0, common.Big0
}

func (tx *L1BlockHashesTx) setSignatureValues(chainID, v, r, s *big.Int) {
	// this is a noop for l1 message transactions
}
