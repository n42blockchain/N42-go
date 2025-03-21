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

package tracers

//func BenchmarkTransactionTrace(b *testing.B) {
//	key, _ := crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
//	from := crypto.PubkeyToAddress(key.PublicKey)
//	gas := uint64(1000000) // 1M gas
//	to := common.HexToAddress("0x00000000000000000000000000000000deadbeef")
//	signer :=transaction.LatestSignerForChainID(big.NewInt(1337))
//	tx, err := transaction.SignNewTx(key, signer,
//		&transaction.LegacyTx{
//			Nonce:    1,
//			GasPrice: uint256.NewInt(500),
//			Gas:      gas,
//			To:       &to,
//		})
//	if err != nil {
//		b.Fatal(err)
//	}
//	txContext := evmtypes.TxContext{
//		Origin:   from,
//		GasPrice: tx.GasPrice(),
//	}
//	context := evmtypes.BlockContext{
//		CanTransfer: core.CanTransfer,
//		Transfer:    core.Transfer,
//		Coinbase:    common.Address{},
//		BlockNumber: uint64(5),
//		Time:        5,
//		Difficulty:  big.NewInt(0xffffffff),
//		GasLimit:    gas,
//		BaseFee:     uint256.NewInt(8),
//	}
//	var alloc []conf.Allocate
//	// The code pushes 'deadbeef' into memory, then the other params, and calls CREATE2, then returns
//	// the address
//	loop := []byte{
//		byte(vm.JUMPDEST), //  [ count ]
//		byte(vm.PUSH1), 0, // jumpdestination
//		byte(vm.JUMP),
//	}
//	alloc[] = conf.Allocate{
//		Address: common.HexToAddress("0x00000000000000000000000000000000deadbeef"),
//		Nonce:   1,
//		Code:    loop,
//		Balance: "1",
//	}
//	alloc[] = conf.Allocate{
//		Address: from,
//		Nonce:   1,
//		Code:    []byte{},
//		Balance: big.NewInt(500000000000000),
//	}
//	_, statedb := tests.MakePreState(rawdb.NewMemoryDatabase(), alloc, false)
//	// Create the tracer, the EVM environment and run it
//	tracer := logger.NewStructLogger(&logger.Config{
//		Debug: false,
//		//DisableStorage: true,
//		//EnableMemory: false,
//		//EnableReturnData: false,
//	})
//	evm := vm.NewEVM(context, txContext, statedb, params.AllEthashProtocolChanges, vm.Config{Debug: true, Tracer: tracer})
//	msg, err := core.TransactionToMessage(tx, signer, nil)
//	if err != nil {
//		b.Fatalf("failed to prepare transaction for tracing: %v", err)
//	}
//	b.ResetTimer()
//	b.ReportAllocs()
//
//	for i := 0; i < b.N; i++ {
//		snap := statedb.Snapshot()
//		st := core.NewStateTransition(evm, msg, new(common2.GasPool).AddGas(tx.Gas()))
//		_, err = st.TransitionDb()
//		if err != nil {
//			b.Fatal(err)
//		}
//		statedb.RevertToSnapshot(snap)
//		if have, want := len(tracer.StructLogs()), 244752; have != want {
//			b.Fatalf("trace wrong, want %d steps, have %d", want, have)
//		}
//		tracer.Reset()
//	}
//}
