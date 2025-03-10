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

package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/n42blockchain/N42/common"
	"github.com/n42blockchain/N42/common/block"
	"github.com/n42blockchain/N42/common/crypto"
	"github.com/n42blockchain/N42/common/crypto/bls"
	"github.com/n42blockchain/N42/common/crypto/bls/blst"
	"github.com/n42blockchain/N42/common/types"
	"github.com/n42blockchain/N42/contracts/deposit"
	"github.com/n42blockchain/N42/internal/consensus"
	"github.com/n42blockchain/N42/log"
	event "github.com/n42blockchain/N42/modules/event/v2"
	"github.com/n42blockchain/N42/modules/rawdb"
	"github.com/n42blockchain/N42/modules/state"
	"golang.org/x/crypto/sha3"
)

var sigChannel = make(chan AggSign, 10)

var validVerifers = map[string]string{
	//"astb9e94477f5f88b5e8da2e97e8506d6e4fcf04e5b": "2c02dd3cf600af9a8567e5cc5ff158c1b89e1f3ea21bff61f505d141a96a60ee",
}

//type WithCodeAndHash struct {
//	CodeIndex []byte `json:"codeIndex"`
//	Code      []byte `json:"code"`
//	Hash      []byte `json:"hash"`
//}

//func ExportCodeAndHash(ctx context.Context, db kv.RwDB) (WithCodeAndHash, error) {
//	var result WithCodeAndHash
//	var err error
//	errs := make(chan error, 1)
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//	// export header hash
//	go func(ctx context.Context) {
//		defer wg.Done()
//		rtx, err := db.BeginRo(ctx)
//		if nil != err {
//			errs <- err
//			return
//		}
//		defer rtx.Rollback()
//
//		buf := new(bytes.Buffer)
//		hashW := zlib.NewWriter(buf)
//		defer hashW.Close()
//
//		cur, err := rtx.Cursor(modules.HeaderCanonical)
//		if nil != err {
//			errs <- err
//			return
//		}
//		defer cur.Close()
//
//		select {
//		case <-ctx.Done():
//			return
//		default:
//			for k, v, err := cur.First(); k != nil; k, v, err = cur.Next() {
//				if nil != err {
//					errs <- err
//					return
//				}
//				//b, _ := modules.DecodeBlockNumber(k)
//				//h := types.Hash{}
//				//h.SetBytes(v)
//				//log.Tracef("read hash, %d, %v", b, h)
//				hashW.Write(v)
//			}
//
//			if err := hashW.Flush(); nil != err {
//				errs <- err
//				return
//			}
//			result.Hash = buf.Bytes()
//		}
//	}(ctx)
//
//	// export code
//	go func(ctx context.Context) {
//		defer wg.Done()
//		rtx, err := db.BeginRo(ctx)
//		if nil != err {
//			errs <- err
//			return
//		}
//		defer rtx.Rollback()
//
//		cur, err := rtx.Cursor(modules.Code)
//		if nil != err {
//			errs <- err
//			return
//		}
//		defer cur.Close()
//
//		indBuf := new(bytes.Buffer)
//		indW := zlib.NewWriter(indBuf)
//		defer indW.Close()
//		codeBuf := new(bytes.Buffer)
//		codeW := zlib.NewWriter(codeBuf)
//		defer codeW.Close()
//		index := uint64(0)
//
//		select {
//		case <-ctx.Done():
//			return
//		default:
//			for k, v, err := cur.First(); k != nil; k, v, err = cur.Next() {
//				if nil != err {
//					errs <- err
//					return
//				}
//				indW.Write(k)
//				indW.Write(modules.EncodeBlockNumber(index))
//				index += uint64(len(v))
//				indW.Write(modules.EncodeBlockNumber(index))
//				codeW.Write(v)
//			}
//			result.CodeIndex = indBuf.Bytes()
//			result.Code = codeBuf.Bytes()
//		}
//	}(ctx)
//
//	select {
//	case e := <-errs:
//		err = e
//		cancel()
//	default:
//		wg.Wait()
//	}
//	close(errs)
//	log.Tracef("export code and hash: %+v", result)
//	return result, err
//}

type AggSign struct {
	Number    uint64          `json:"number"`
	StateRoot types.Hash      `json:"stateRoot"`
	Sign      types.Signature `json:"sign"`
	Address   types.Address   `json:"address"`
	PublicKey types.PublicKey `json:"-"`
}

func (s *AggSign) Check(root types.Hash) bool {
	if s.StateRoot != root {
		return false
	}
	sig, err := bls.SignatureFromBytes(s.Sign[:])
	if nil != err {
		return false
	}

	pub, err := bls.PublicKeyFromBytes(s.PublicKey[:])
	if nil != err {
		return false
	}
	return sig.Verify(pub, s.StateRoot[:])
}

func DepositInfo(db kv.RwDB, key types.Address) *deposit.Info {
	var info *deposit.Info
	_ = db.View(context.Background(), func(tx kv.Tx) error {
		info = deposit.GetDepositInfo(tx, key)
		return nil
	})
	return info
}

func IsDeposit(db kv.RwDB, addr types.Address) (bool, error) {
	tx, err := db.BeginRo(context.Background())
	if nil != err {
		return false, err
	}
	defer tx.Rollback()

	return rawdb.IsDeposit(tx, addr), nil
}

func SignMerge(ctx context.Context, header *block.Header, depositNum uint64) (types.Signature, []*block.Verify, error) {
	aggrSigns := make([]bls.Signature, 0)
	verifiers := make([]*block.Verify, 0)
	uniq := make(map[types.Address]struct{})

LOOP:
	for {
		select {
		case s := <-sigChannel:
			log.Tracef("accept sign, %+v", s)
			if s.Number != header.Number.Uint64() {
				log.Tracef("discard sign: need block number %d, get %d", header.Number.Uint64(), s.Number)
				continue
			}

			if _, ok := uniq[s.Address]; ok {
				continue
			}

			if !s.Check(header.Root) {
				log.Tracef("discard sign: sign check failed! %v", s)
				continue
			}
			sig, err := bls.SignatureFromBytes(s.Sign[:])
			if nil != err {
				return types.Signature{}, nil, err
			}

			aggrSigns = append(aggrSigns, sig)
			verifiers = append(verifiers, &block.Verify{
				Address:   s.Address,
				PublicKey: s.PublicKey,
			})
			uniq[s.Address] = struct{}{}
		case <-ctx.Done():
			break LOOP
		}
	}
	// todo enough sigs check
	// 1
	// uint64(len(aggrSigns)) < depositNum/2
	// uint64(len(aggrSigns)) < 7
	if uint64(len(aggrSigns)) < 3 {
		return types.Signature{}, nil, consensus.ErrNotEnoughSign
	}

	aggS := blst.AggregateSignatures(aggrSigns)
	var aggSign types.Signature
	copy(aggSign[:], aggS.Marshal())
	return aggSign, verifiers, nil
}

func MachineVerify(ctx context.Context) error {
	entire := make(chan common.MinedEntireEvent)
	blocksSub := event.GlobalEvent.Subscribe(entire)
	defer blocksSub.Unsubscribe()

	errs := make(chan error)
	defer close(errs)

	for {
		select {
		case b := <-entire:
			log.Tracef("machine verify accept entire, number: %d", b.Entire.Entire.Header.Number.Uint64())
			for k, s := range validVerifers {
				go func(seckey string, address string) {
					// recover private key
					sByte, err := hex.DecodeString(seckey)
					if nil != err {
						errs <- err
						return
					}
					var addr types.Address
					if !addr.DecodeString(address) {
						errs <- fmt.Errorf("unvalid address")
						return
					}

					// before state verify
					var hash types.Hash
					hasher := sha3.NewLegacyKeccak256()
					state.EncodeBeforeState(hasher, b.Entire.Entire.Snap.Items, b.Entire.Codes)
					_, err = hasher.(crypto.KeccakState).Read(hash[:])
					if err != nil {
						return
					}
					if b.Entire.Entire.Header.MixDigest != hash {
						log.Warn("misMatch before state hash", "want:", b.Entire.Entire.Header.MixDigest, "get:", hash, b.Entire.Entire.Header.Number.Uint64())
						return
					}

					// publicKey
					var bs [32]byte
					copy(bs[:], sByte)
					pri, err := bls.SecretKeyFromRandom32Byte(bs)
					if nil != err {
						errs <- err
						return
					}

					// Signature
					sign := pri.Sign(b.Entire.Entire.Header.Root[:])
					tmp := AggSign{Number: b.Entire.Entire.Header.Number.Uint64()}
					copy(tmp.StateRoot[:], b.Entire.Entire.Header.Root[:])
					copy(tmp.Sign[:], sign.Marshal())
					copy(tmp.PublicKey[:], pri.PublicKey().Marshal())
					tmp.Address = addr
					// send res
					sigChannel <- tmp
					//log.Tracef("send verify sign, %+v", tmp)
				}(s, k)
			}
		case <-ctx.Done():
			return nil
		case err := <-errs:
			return err
		}
	}
}
