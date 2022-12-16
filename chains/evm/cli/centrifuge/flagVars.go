package centrifuge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nonceblox/elysium-chainsafe-core/crypto/secp256k1"
)

//flag vars
var (
	Hash    string
	Address string
)

//processed flag vars
var (
	StoreAddr common.Address
	ByteHash  [32]byte
)

// global flags
var (
	url           string
	gasPrice      *big.Int
	senderKeyPair *secp256k1.Keypair
	prepare       bool
)
