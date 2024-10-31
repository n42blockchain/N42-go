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

package api

import (
	"github.com/holiman/uint256"
	"github.com/n42blockchain/N42/common"
	"github.com/n42blockchain/N42/common/block"
	"github.com/n42blockchain/N42/common/hash"
	"github.com/n42blockchain/N42/common/hexutil"
	"github.com/n42blockchain/N42/common/transaction"
	"github.com/n42blockchain/N42/common/types"
	mvm_types "github.com/n42blockchain/N42/internal/avm/types"
	"math/big"
)

func RPCMarshalBlock(block block.IBlock, chain common.IBlockChain, inclTx bool, fullTx bool) (map[string]interface{}, error) {
	fields := RPCMarshalHeader(block.Header())

	if inclTx {
		formatTx := func(tx *transaction.Transaction) (interface{}, error) {
			hash := tx.Hash()
			return mvm_types.FromastHash(hash), nil
		}
		if fullTx {
			formatTx = func(tx *transaction.Transaction) (interface{}, error) {
				hash := tx.Hash()
				return newRPCTransactionFromBlockHash(block, hash), nil
			}
		}
		txs := block.Transactions()
		transactions := make([]interface{}, len(txs))
		var err error
		for i, tx := range txs {
			if transactions[i], err = formatTx(tx); err != nil {
				return nil, err
			}
		}
		fields["transactions"] = transactions

		// verifiers
		verifiers := make([]interface{}, len(block.Body().Verifier()))
		for i, verifier := range block.Body().Verifier() {
			verifiers[i] = verifier
		}
		fields["verifier"] = verifiers

		// reward todo
		type RPCReward struct {
			Address types.Address
			Amount  *uint256.Int
		}
		rewards := make([]*RPCReward, len(block.Body().Reward()))
		for i, reward := range block.Body().Reward() {
			rewards[i] = &RPCReward{
				reward.Address,
				reward.Amount,
			}
		}
		fields["rewards"] = rewards

		td := chain.GetTd(block.Hash(), block.Number64())
		if td == nil {
			td = new(uint256.Int)
		}
		fields["totalDifficulty"] = (*hexutil.Big)(td.ToBig())

	}
	// POA
	uncleHashes := make([]types.Hash, 0)
	fields["uncles"] = uncleHashes

	return fields, nil
}

// newRPCTransactionFromBlockHash returns a transaction that will serialize to the RPC representation.
func newRPCTransactionFromBlockHash(b block.IBlock, findHash types.Hash) *RPCTransaction {
	for idx, tx := range b.Transactions() {
		hash := tx.Hash()
		if hash == findHash {
			return newRPCTransactionFromBlockIndex(b, uint64(idx))
		}
	}
	return nil
}

// newRPCTransactionFromBlockIndex returns a transaction that will serialize to the RPC representation.
func newRPCTransactionFromBlockIndex(b block.IBlock, index uint64) *RPCTransaction {
	txs := b.Transactions()
	if index >= uint64(len(txs)) {
		return nil
	}
	return newRPCTransaction(txs[index], b.Hash(), b.Number64().Uint64(), index, big.NewInt(baseFee))
}

// RPCMarshalHeader converts the given header to the RPC output .
func RPCMarshalHeader(head block.IHeader) map[string]interface{} {
	header := head.(*block.Header)
	ethHeader := mvm_types.FromastHeader(head)

	result := map[string]interface{}{
		"number":           (*hexutil.Big)(head.Number64().ToBig()),
		"hash":             mvm_types.FromastHash(header.Hash()),
		"parentHash":       mvm_types.FromastHash(header.ParentHash),
		"nonce":            header.Nonce,
		"mixHash":          mvm_types.FromastHash(header.MixDigest),
		"sha3Uncles":       mvm_types.FromastHash(hash.EmptyUncleHash),
		"miner":            mvm_types.FromastAddress(&header.Coinbase),
		"difficulty":       (*hexutil.Big)(header.Difficulty.ToBig()),
		"extraData":        hexutil.Bytes(header.Extra),
		"size":             hexutil.Uint64(ethHeader.Size()),
		"gasLimit":         hexutil.Uint64(header.GasLimit),
		"gasUsed":          hexutil.Uint64(header.GasUsed),
		"timestamp":        hexutil.Uint64(header.Time),
		"transactionsRoot": mvm_types.FromastHash(header.TxHash),
		"receiptsRoot":     mvm_types.FromastHash(header.ReceiptHash),
		"logsBloom":        ethHeader.Bloom,
		"stateRoot":        mvm_types.FromastHash(header.Root),
		"signature":        header.Signature,
	}

	if header.BaseFee != nil {
		result["baseFeePerGas"] = (*hexutil.Big)(header.BaseFee.ToBig())
	}

	return result
}
