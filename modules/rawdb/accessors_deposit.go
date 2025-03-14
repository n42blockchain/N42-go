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

package rawdb

import (
	"fmt"
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/n42blockchain/N42/common/crypto/bls"
	"github.com/n42blockchain/N42/common/types"
	"github.com/n42blockchain/N42/modules"
)

//// PutDeposit
//func PutDeposit(db kv.Putter, key []byte, val []byte) error {
//	return db.Put(modules.Deposit, key, val)
//}

func PutDeposit(db kv.Putter, addr types.Address, pub types.PublicKey, amount uint256.Int) error {

	data := make([]byte, types.PublicKeyLength+amount.ByteLen())
	copy(data[:types.PublicKeyLength], pub.Bytes())
	copy(data[types.PublicKeyLength:], amount.Bytes())
	//
	if err := db.Put(modules.Deposit, addr[:], data); err != nil {
		return fmt.Errorf("failed to store address Deposit: %w", err)
	}
	return nil
}

// GetDeposit
func GetDeposit(db kv.Getter, addr types.Address) (types.PublicKey, *uint256.Int, error) {
	valBytes, err := db.GetOne(modules.Deposit, addr[:])
	if err != nil {
		return types.PublicKey{}, nil, err
	}
	if len(valBytes) < types.PublicKeyLength {
		return types.PublicKey{}, nil, fmt.Errorf("the data length wrong")
	}
	_, err = bls.PublicKeyFromBytes(valBytes[:types.PublicKeyLength])
	if err != nil {
		return types.PublicKey{}, nil, fmt.Errorf("cannot unmarshal pubkey from bytes")
	}
	pubkey := new(types.PublicKey)
	if err = pubkey.SetBytes(valBytes[:types.PublicKeyLength]); err != nil {
		return types.PublicKey{}, nil, fmt.Errorf("cannot unmarshal pubkey from bytes")
	}
	amount := uint256.NewInt(0).SetBytes(valBytes[types.PublicKeyLength:])

	return *pubkey, amount, nil
}

// DeleteDeposit removes Deposit data associated with an address.
func DeleteDeposit(db kv.Deleter, addr types.Address) error {
	return db.Delete(modules.Deposit, addr[:])
}

// IsDeposit is deposit account
func IsDeposit(db kv.Getter, addr types.Address) bool {
	is, _ := db.Has(modules.Deposit, addr[:])
	return is
}

func DepositNum(tx kv.Tx) (uint64, error) {
	cur, err := tx.Cursor(modules.Deposit)
	if nil != err {
		return 0, err
	}
	defer cur.Close()
	return cur.Count()
}
