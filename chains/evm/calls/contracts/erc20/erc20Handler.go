package erc20

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nonceblox/elysium-chainsafe-core/chains/evm/calls"
	"github.com/nonceblox/elysium-chainsafe-core/chains/evm/calls/consts"
	"github.com/nonceblox/elysium-chainsafe-core/chains/evm/calls/contracts"
	"github.com/nonceblox/elysium-chainsafe-core/chains/evm/calls/transactor"
)

type ERC20HandlerContract struct {
	contracts.Contract
}

func NewERC20HandlerContract(
	client calls.ContractCallerDispatcher,
	erc20HandlerContractAddress common.Address,
	t transactor.Transactor,
) *ERC20HandlerContract {
	a, _ := abi.JSON(strings.NewReader(consts.ERC20HandlerABI))
	b := common.FromHex(consts.ERC20HandlerBin)
	return &ERC20HandlerContract{contracts.NewContract(erc20HandlerContractAddress, a, b, client, t)}
}
