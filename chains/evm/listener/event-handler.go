package listener

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nonceblox/elysium-chainsafe-core/chains/evm/calls/events"
	"github.com/nonceblox/elysium-chainsafe-core/relayer/message"
	"github.com/nonceblox/elysium-chainsafe-core/types"
	"github.com/rs/zerolog/log"
)

type EventListener interface {
	FetchDeposits(ctx context.Context, address common.Address, startBlock *big.Int, endBlock *big.Int) ([]*events.Deposit, error)
}

type DepositHandler interface {
	HandleDeposit(sourceID, destID uint8, nonce uint64, resourceID types.ResourceID, calldata, handlerResponse []byte) (*message.Message, error)
}

type DepositEventHandler struct {
	eventListener  EventListener
	depositHandler DepositHandler
	bridgeAddress  common.Address
	domainID       uint8
}

func NewDepositEventHandler(eventListener EventListener, depositHandler DepositHandler, bridgeAddress common.Address, domainID uint8) *DepositEventHandler {
	return &DepositEventHandler{
		eventListener:  eventListener,
		depositHandler: depositHandler,
		bridgeAddress:  bridgeAddress,
		domainID:       domainID,
	}
}

func (eh *DepositEventHandler) HandleEvent(block *big.Int, msgChan chan *message.Message) error {
	deposits, err := eh.eventListener.FetchDeposits(context.Background(), eh.bridgeAddress, block, block)
	if err != nil {
		return fmt.Errorf("unable to fetch deposit events because of: %+v", err)
	}

	for _, d := range deposits {
		m, err := eh.depositHandler.HandleDeposit(eh.domainID, d.DestinationDomainID, d.DepositNonce, d.ResourceID, d.Data, d.HandlerResponse)
		if err != nil {
			log.Error().Str("block", block.String()).Uint8("domainID", eh.domainID).Msgf("%v", err)
			continue
		}
		log.Debug().Msgf("Resolved message %+v in block %s", m, block.String())
		msgChan <- m
	}
	log.Debug().Msgf("Queried block  %s", block.String())
	return nil
}
