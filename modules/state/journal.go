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

package state

import (
	"github.com/holiman/uint256"
	"github.com/n42blockchain/N42/common/types"
)

// journalEntry is a modification entry in the state change journal that can be
// reverted on demand.
type journalEntry interface {
	// revert undoes the changes introduced by this journal entry.
	revert(*IntraBlockState)

	// dirtied returns the Ethereum address modified by this journal entry.
	dirtied() *types.Address
}

// journal contains the list of state modifications applied since the last state
// commit. These are tracked to be able to be reverted in case of an execution
// exception or revertal request.
type journal struct {
	entries []journalEntry        // Current changes tracked by the journal
	dirties map[types.Address]int // Dirty accounts and the number of changes
}

// newJournal create a new initialized journal.
func newJournal() *journal {
	return &journal{
		dirties: make(map[types.Address]int),
	}
}

// append inserts a new modification entry to the end of the change journal.
func (j *journal) append(entry journalEntry) {
	j.entries = append(j.entries, entry)
	if addr := entry.dirtied(); addr != nil {
		j.dirties[*addr]++
	}
}

// revert undoes a batch of journalled modifications along with any reverted
// dirty handling too.
func (j *journal) revert(statedb *IntraBlockState, snapshot int) {
	for i := len(j.entries) - 1; i >= snapshot; i-- {
		// Undo the changes made by the operation
		j.entries[i].revert(statedb)

		// Drop any dirty tracking induced by the change
		if addr := j.entries[i].dirtied(); addr != nil {
			if j.dirties[*addr]--; j.dirties[*addr] == 0 {
				delete(j.dirties, *addr)
			}
		}
	}
	j.entries = j.entries[:snapshot]
}

// dirty explicitly sets an address to dirty, even if the change entries would
// otherwise suggest it as clean. This method is an ugly hack to handle the RIPEMD
// precompile consensus exception.
func (j *journal) dirty(addr types.Address) {
	j.dirties[addr]++
}

// length returns the current number of entries in the journal.
func (j *journal) length() int {
	n := len(j.entries)
	return n
}

type (
	// Changes to the account trie.
	createObjectChange struct {
		account *types.Address
	}
	resetObjectChange struct {
		account *types.Address
		prev    *stateObject
	}
	selfdestructChange struct {
		account     *types.Address
		prev        bool // whether account had already selfdestructed
		prevbalance uint256.Int
	}

	// Changes to individual accounts.
	balanceChange struct {
		account *types.Address
		prev    uint256.Int
	}
	balanceIncrease struct {
		account  *types.Address
		increase uint256.Int
	}
	balanceIncreaseTransfer struct {
		bi *BalanceIncrease
	}
	nonceChange struct {
		account *types.Address
		prev    uint64
	}
	storageChange struct {
		account  *types.Address
		key      types.Hash
		prevalue uint256.Int
	}
	fakeStorageChange struct {
		account  *types.Address
		key      types.Hash
		prevalue uint256.Int
	}
	codeChange struct {
		account  *types.Address
		prevcode []byte
		prevhash types.Hash
	}

	// Changes to other state values.
	refundChange struct {
		prev uint64
	}
	addLogChange struct {
		txhash types.Hash
	}
	touchChange struct {
		account *types.Address
	}
	// Changes to the access list
	accessListAddAccountChange struct {
		address *types.Address
	}
	accessListAddSlotChange struct {
		address *types.Address
		slot    *types.Hash
	}
)

func (ch createObjectChange) revert(s *IntraBlockState) {
	delete(s.stateObjects, *ch.account)
	delete(s.stateObjectsDirty, *ch.account)
}

func (ch createObjectChange) dirtied() *types.Address {
	return ch.account
}

func (ch resetObjectChange) revert(s *IntraBlockState) {
	s.setStateObject(*ch.account, ch.prev)
}

func (ch resetObjectChange) dirtied() *types.Address {
	return nil
}

func (ch selfdestructChange) revert(s *IntraBlockState) {
	obj := s.getStateObject(*ch.account)
	if obj != nil {
		obj.selfdestructed = ch.prev
		obj.setBalance(&ch.prevbalance)
	}
}

func (ch selfdestructChange) dirtied() *types.Address {
	return ch.account
}

var ripemd = types.HexToAddress("0000000000000000000000000000000000000003")

func (ch touchChange) revert(s *IntraBlockState) {
}

func (ch touchChange) dirtied() *types.Address {
	return ch.account
}

func (ch balanceChange) revert(s *IntraBlockState) {
	s.getStateObject(*ch.account).setBalance(&ch.prev)
}

func (ch balanceChange) dirtied() *types.Address {
	return ch.account
}

func (ch balanceIncrease) revert(s *IntraBlockState) {
	if bi, ok := s.balanceInc[*ch.account]; ok {
		bi.increase.Sub(&bi.increase, &ch.increase)
		bi.count--
		if bi.count == 0 {
			delete(s.balanceInc, *ch.account)
		}
	}
}

func (ch balanceIncrease) dirtied() *types.Address {
	return ch.account
}

func (ch balanceIncreaseTransfer) dirtied() *types.Address {
	return nil
}

func (ch balanceIncreaseTransfer) revert(s *IntraBlockState) {
	ch.bi.transferred = false
}
func (ch nonceChange) revert(s *IntraBlockState) {
	s.getStateObject(*ch.account).setNonce(ch.prev)
}

func (ch nonceChange) dirtied() *types.Address {
	return ch.account
}

func (ch codeChange) revert(s *IntraBlockState) {
	s.getStateObject(*ch.account).setCode(ch.prevhash, ch.prevcode)
}

func (ch codeChange) dirtied() *types.Address {
	return ch.account
}

func (ch storageChange) revert(s *IntraBlockState) {
	s.getStateObject(*ch.account).setState(&ch.key, ch.prevalue)
}

func (ch storageChange) dirtied() *types.Address {
	return ch.account
}

func (ch fakeStorageChange) revert(s *IntraBlockState) {
	s.getStateObject(*ch.account).fakeStorage[ch.key] = ch.prevalue
}

func (ch fakeStorageChange) dirtied() *types.Address {
	return ch.account
}

func (ch refundChange) revert(s *IntraBlockState) {
	s.refund = ch.prev
}

func (ch refundChange) dirtied() *types.Address {
	return nil
}

func (ch addLogChange) revert(s *IntraBlockState) {
	logs := s.logs[ch.txhash]
	if len(logs) == 1 {
		delete(s.logs, ch.txhash)
	} else {
		s.logs[ch.txhash] = logs[:len(logs)-1]
	}
	s.logSize--
}

func (ch addLogChange) dirtied() *types.Address {
	return nil
}

func (ch accessListAddAccountChange) revert(s *IntraBlockState) {
	/*
		One important invariant here, is that whenever a (addr, slot) is added, if the
		addr is not already present, the add causes two journal entries:
		- one for the address,
		- one for the (address,slot)
		Therefore, when unrolling the change, we can always blindly delete the
		(addr) at this point, since no storage adds can remain when come upon
		a single (addr) change.
	*/
	s.accessList.DeleteAddress(*ch.address)
}

func (ch accessListAddAccountChange) dirtied() *types.Address {
	return nil
}

func (ch accessListAddSlotChange) revert(s *IntraBlockState) {
	s.accessList.DeleteSlot(*ch.address, *ch.slot)
}

func (ch accessListAddSlotChange) dirtied() *types.Address {
	return nil
}
