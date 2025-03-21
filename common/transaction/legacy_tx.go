// Copyright 2022 The N42 Authors
// This file is part of the N42 library.
//
// The N42 library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The N42 library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the N42 library. If not, see <http://www.gnu.org/licenses/>.

package transaction

import (
	"github.com/holiman/uint256"
	"github.com/n42blockchain/N42/common/hash"
	"github.com/n42blockchain/N42/common/types"
)

// LegacyTx is the transaction data of regular Ethereum transactions.
type LegacyTx struct {
	Nonce    uint64         // nonce of sender account
	GasPrice *uint256.Int   // wei per gas
	Gas      uint64         // gas limit
	To       *types.Address `rlp:"nil"` // nil means contract creation
	From     *types.Address `rlp:"nil"` // nil means contract creation
	Value    *uint256.Int   // wei amount
	Data     []byte         // contract invocation input data
	V, R, S  *uint256.Int   // signature values
	Sign     []byte
}

// NewTransaction creates an unsigned legacy transaction.
// Deprecated: use NewTx instead.
func NewTransaction(nonce uint64, from types.Address, to *types.Address, amount *uint256.Int, gasLimit uint64, gasPrice *uint256.Int, data []byte) *Transaction {
	return NewTx(&LegacyTx{
		Nonce:    nonce,
		To:       to,
		From:     &from,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})
}

// NewContractCreation creates an unsigned legacy transaction.
// Deprecated: use NewTx instead.
func NewContractCreation(nonce uint64, amount *uint256.Int, gasLimit uint64, gasPrice *uint256.Int, data []byte) *Transaction {
	return NewTx(&LegacyTx{
		Nonce:    nonce,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *LegacyTx) copy() TxData {
	cpy := &LegacyTx{
		Nonce: tx.Nonce,
		To:    copyAddressPtr(tx.To),
		From:  copyAddressPtr(tx.From),
		Data:  types.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are initialized below.
		Value:    new(uint256.Int),
		GasPrice: new(uint256.Int),
		V:        new(uint256.Int),
		R:        new(uint256.Int),
		S:        new(uint256.Int),
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	if tx.sign() != nil {
		copy(cpy.Sign, tx.Sign)
	}

	return cpy
}

// accessors for innerTx.
func (tx *LegacyTx) txType() byte { return LegacyTxType }
func (tx *LegacyTx) chainID() *uint256.Int {
	return DeriveChainId(tx.V)
}
func (tx *LegacyTx) accessList() AccessList  { return nil }
func (tx *LegacyTx) data() []byte            { return tx.Data }
func (tx *LegacyTx) gas() uint64             { return tx.Gas }
func (tx *LegacyTx) gasPrice() *uint256.Int  { return tx.GasPrice }
func (tx *LegacyTx) gasTipCap() *uint256.Int { return tx.GasPrice }
func (tx *LegacyTx) gasFeeCap() *uint256.Int { return tx.GasPrice }
func (tx *LegacyTx) value() *uint256.Int     { return tx.Value }
func (tx *LegacyTx) nonce() uint64           { return tx.Nonce }
func (tx *LegacyTx) to() *types.Address      { return tx.To }
func (tx *LegacyTx) from() *types.Address    { return tx.From }
func (tx *LegacyTx) sign() []byte            { return tx.Sign }

func (tx *LegacyTx) rawSignatureValues() (v, r, s *uint256.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *LegacyTx) setSignatureValues(chainID, v, r, s *uint256.Int) {
	tx.V, tx.R, tx.S = v, r, s
}

func (tx *LegacyTx) hash() types.Hash {
	hash := hash.RlpHash([]interface{}{
		tx.Nonce,
		tx.GasPrice,
		tx.Gas,
		tx.To,
		tx.Value,
		tx.Data,
		tx.V, tx.R, tx.S,
	})
	return hash
}
