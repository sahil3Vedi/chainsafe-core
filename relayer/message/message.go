// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package message

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nonceblox/elysium-chainsafe-core/types"
)

type TransferType string
type Metadata struct {
	Priority uint8
	Blob     []byte
}

const (
	FungibleTransfer    TransferType = "FungibleTransfer"
	NonFungibleTransfer TransferType = "NonFungibleTransfer"
	GenericTransfer     TransferType = "GenericTransfer"
)

type ProposalStatus struct {
	Status        uint8
	YesVotes      *big.Int
	YesVotesTotal uint8
	ProposedBlock *big.Int
}

const (
	ProposalStatusInactive uint8 = iota
	ProposalStatusActive
	ProposalStatusPassed // Ready to be executed
	ProposalStatusExecuted
	ProposalStatusCanceled
)

var (
	StatusMap = map[uint8]string{ProposalStatusInactive: "inactive", ProposalStatusActive: "active", ProposalStatusPassed: "passed", ProposalStatusExecuted: "executed", ProposalStatusCanceled: "canceled"}
)

type Message struct {
	Source       uint8  // Source where message was initiated
	Destination  uint8  // Destination chain of message
	DepositNonce uint64 // Nonce for the deposit
	ResourceId   types.ResourceID
	Payload      []interface{} // data associated with event sequence
	Metadata     Metadata      // Arbitrary data that will be most likely be used by the relayer
	Type         TransferType
}
type Message2 struct {
	Source             uint8  // Source where message was initiated
	Destination        uint8  // Destination chain of message
	DepositNonce       uint64 // Nonce for the deposit
	ResourceId         types.ResourceID
	Sourcehandler      common.Address
	Desthandler        common.Address
	DestBridgeAddress  common.Address
	SourceBrigeAddress common.Address
	SourceTokenAddress common.Address
	DestTokenAddress   common.Address
}

func NewMessage(
	source uint8,
	destination uint8,
	depositNonce uint64,
	resourceId types.ResourceID,
	transferType TransferType,
	payload []interface{},
	metadata Metadata,
) *Message {
	return &Message{
		source,
		destination,
		depositNonce,
		resourceId,
		payload,
		metadata,
		transferType,
	}
}

func NewMessage1(
	source uint8,
	destination uint8,
	depositNonce uint64,
	resourceId types.ResourceID,
	sourcehandler common.Address,
	Desthandler common.Address,
	DestBridgeAddress common.Address,
	SourceBrigeAddress common.Address,
	SourceTokenAddress common.Address,
	DestTokenAddress common.Address,
) *Message2 {
	return &Message2{
		source,
		destination,
		depositNonce,
		resourceId,
		sourcehandler,
		Desthandler,
		DestBridgeAddress,
		SourceBrigeAddress,
		SourceTokenAddress,
		DestTokenAddress,
	}
}
