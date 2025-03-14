// Package bls implements a go-wrapper around a library implementing the
// the BLS12-381 curve and signature scheme. This package exposes a public API for
// verifying and aggregating BLS signatures used by Ethereum.
package bls

import (
	"github.com/n42blockchain/N42/common/crypto/bls/blst"
	"github.com/n42blockchain/N42/common/crypto/bls/common"
)

// Initialize herumi temporarily while we transition to blst for ethdo.
func init() {
}

// SecretKeyFromBytes creates a BLS private key from a BigEndian byte slice.
func SecretKeyFromBytes(privKey []byte) (SecretKey, error) {
	return blst.SecretKeyFromBytes(privKey)
}

// PublicKeyFromBytes creates a BLS public key from a  BigEndian byte slice.
func PublicKeyFromBytes(pubKey []byte) (PublicKey, error) {
	return blst.PublicKeyFromBytes(pubKey)
}

// SignatureFromBytes creates a BLS signature from a LittleEndian byte slice.
func SignatureFromBytes(sig []byte) (Signature, error) {
	return blst.SignatureFromBytes(sig)
}

// MultipleSignaturesFromBytes creates a slice of BLS signatures from a LittleEndian 2d-byte slice.
func MultipleSignaturesFromBytes(sigs [][]byte) ([]Signature, error) {
	return blst.MultipleSignaturesFromBytes(sigs)
}

// AggregatePublicKeys aggregates the provided raw public keys into a single key.
func AggregatePublicKeys(pubs [][]byte) (PublicKey, error) {
	return blst.AggregatePublicKeys(pubs)
}

// AggregateMultiplePubkeys aggregates the provided decompressed keys into a single key.
func AggregateMultiplePubkeys(pubs []PublicKey) PublicKey {
	return blst.AggregateMultiplePubkeys(pubs)
}

// AggregateSignatures converts a list of signatures into a single, aggregated sig.
func AggregateSignatures(sigs []common.Signature) common.Signature {
	return blst.AggregateSignatures(sigs)
}

// AggregateCompressedSignatures converts a list of compressed signatures into a single, aggregated sig.
func AggregateCompressedSignatures(multiSigs [][]byte) (common.Signature, error) {
	return blst.AggregateCompressedSignatures(multiSigs)
}

// VerifyMultipleSignatures verifies multiple signatures for distinct messages securely.
func VerifyMultipleSignatures(sigs [][]byte, msgs [][32]byte, pubKeys []common.PublicKey) (bool, error) {
	return blst.VerifyMultipleSignatures(sigs, msgs, pubKeys)
}

// NewAggregateSignature creates a blank aggregate signature.
func NewAggregateSignature() common.Signature {
	return blst.NewAggregateSignature()
}

// RandKey creates a new private key using a random input.
func RandKey() (common.SecretKey, error) {
	return blst.RandKey()
}

// SecretKeyFromRandom32Byte creates a new private key using random 32 bytes
func SecretKeyFromRandom32Byte(ikm [32]byte) (common.SecretKey, error) {
	return blst.SecretKeyFromRandom32Byte(ikm)
}
