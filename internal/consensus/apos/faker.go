// Copyright 2023 The N42 Authors
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

package apos

import (
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/n42blockchain/N42/common/block"
	"github.com/n42blockchain/N42/common/transaction"
	"github.com/n42blockchain/N42/common/types"
	"github.com/n42blockchain/N42/internal/consensus"
	"github.com/n42blockchain/N42/modules/rpc/jsonrpc"
	"github.com/n42blockchain/N42/modules/state"
	"github.com/n42blockchain/N42/params"
)

type Faker struct{}

func (f Faker) Author(header block.IHeader) (types.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) VerifyHeader(chain consensus.ChainHeaderReader, header block.IHeader, seal bool) error {
	//TODO implement me
	panic("implement me")
}

func (f Faker) VerifyHeaders(chain consensus.ChainHeaderReader, headers []block.IHeader, seals []bool) (chan<- struct{}, <-chan error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) VerifyUncles(chain consensus.ChainReader, block block.IBlock) error {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Prepare(chain consensus.ChainHeaderReader, header block.IHeader) error {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Finalize(chain consensus.ChainHeaderReader, header block.IHeader, state *state.IntraBlockState, txs []*transaction.Transaction, uncles []block.IHeader) ([]*block.Reward, map[types.Address]*uint256.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header block.IHeader, state *state.IntraBlockState, txs []*transaction.Transaction, uncles []block.IHeader, receipts []*block.Receipt) (block.IBlock, []*block.Reward, map[types.Address]*uint256.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Rewards(tx kv.RwTx, header block.IHeader, state *state.IntraBlockState, setRewards bool) ([]*block.Reward, error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Seal(chain consensus.ChainHeaderReader, block block.IBlock, results chan<- block.IBlock, stop <-chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func (f Faker) SealHash(header block.IHeader) types.Hash {
	//TODO implement me
	panic("implement me")
}

func (f Faker) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent block.IHeader) *uint256.Int {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Type() params.ConsensusType {
	return params.Faker
}

func (f Faker) APIs(chain consensus.ChainReader) []jsonrpc.API {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Close() error {
	//TODO implement me
	panic("implement me")
}

func NewFaker() consensus.Engine {
	return &Faker{}
}

func (f Faker) IsServiceTransaction(sender types.Address, syscall consensus.SystemCall) bool {
	return false
}
