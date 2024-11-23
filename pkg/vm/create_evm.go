// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package vm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanche-cli/pkg/application"
	"github.com/ava-labs/avalanche-cli/pkg/binutils"
	"github.com/ava-labs/avalanche-cli/pkg/models"
	"github.com/ava-labs/avalanche-cli/pkg/teleporter"
	icmgenesis "github.com/ava-labs/avalanche-cli/pkg/teleporter/genesis"
	"github.com/ava-labs/avalanche-cli/pkg/validatormanager"
	blockchainSDK "github.com/ava-labs/avalanche-cli/sdk/blockchain"
	"github.com/ava-labs/subnet-evm/core"
	"github.com/ava-labs/subnet-evm/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (
	// 600 AVAX: to deploy teleporter contract, registry contract, and fund
	// starting relayer operations
	teleporterBalance = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(600))
	// 1000 AVAX: to deploy teleporter contract, registry contract, fund
	// starting relayer operations, and deploy bridge contracts
	externalGasTokenBalance = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1000))
)

func CreateEvmSidecar(
	sc *models.Sidecar,
	app *application.Avalanche,
	subnetName string,
	subnetEVMVersion string,
	tokenSymbol string,
	getRPCVersionFromBinary bool,
	sovereign bool,
) (*models.Sidecar, error) {
	var (
		err        error
		rpcVersion int
	)

	if sc == nil {
		sc = &models.Sidecar{}
	}

	if getRPCVersionFromBinary {
		_, vmBin, err := binutils.SetupSubnetEVM(app, subnetEVMVersion)
		if err != nil {
			return nil, fmt.Errorf("failed to install subnet-evm: %w", err)
		}
		rpcVersion, err = GetVMBinaryProtocolVersion(vmBin)
		if err != nil {
			return nil, fmt.Errorf("unable to get RPC version: %w", err)
		}
	} else {
		rpcVersion, err = GetRPCProtocolVersion(app, models.SubnetEvm, subnetEVMVersion)
		if err != nil {
			return nil, err
		}
	}

	sc.Name = subnetName
	sc.VM = models.SubnetEvm
	sc.VMVersion = subnetEVMVersion
	sc.RPCVersion = rpcVersion
	sc.Subnet = subnetName
	sc.TokenSymbol = tokenSymbol
	sc.TokenName = tokenSymbol + " Token"
	sc.Sovereign = sovereign
	return sc, nil
}

func CreateEVMGenesis(
	params SubnetEVMGenesisParams,
	teleporterInfo *teleporter.Info,
	addICMRegistryToGenesis bool,
	proxyOwner string,
	rewardBasisPoints uint64,
) ([]byte, error) {
	feeConfig := getFeeConfig(params)

	// Validity checks on the parameter settings.
	if params.enableTransactionPrecompile {
		if someoneWasAllowed(params.transactionPrecompileAllowList) &&
			!someAllowedHasBalance(params.transactionPrecompileAllowList, params.initialTokenAllocation) {
			return nil, errors.New("none of the addresses in the transaction allow list precompile have any tokens allocated to them. Currently, no address can transact on the network. Allocate some funds to one of the allow list addresses to continue")
		}
	}
	if (params.UseTeleporter || params.UseExternalGasToken) && !params.enableWarpPrecompile {
		return nil, fmt.Errorf("a teleporter enabled blockchain was requested but warp precompile is disabled")
	}
	if (params.UseTeleporter || params.UseExternalGasToken) && teleporterInfo == nil {
		return nil, fmt.Errorf("a teleporter enabled blockchain was requested but no teleporter info was provided")
	}

	// Add the teleporter deployer to the initial token allocation if necessary.
	if params.UseTeleporter || params.UseExternalGasToken {
		balance := teleporterBalance
		if params.UseExternalGasToken {
			balance = externalGasTokenBalance
		}
		if params.initialTokenAllocation == nil {
			params.initialTokenAllocation = core.GenesisAlloc{}
		}
		params.initialTokenAllocation[common.HexToAddress(teleporterInfo.FundedAddress)] = core.GenesisAccount{
			Balance: balance,
		}
		icmgenesis.AddICMMessengerContractToAllocations(params.initialTokenAllocation)
		if addICMRegistryToGenesis {
			// experimental
			if err := icmgenesis.AddICMRegistryContractToAllocations(params.initialTokenAllocation); err != nil {
				return nil, err
			}
		}
	}

	if params.UsePoAValidatorManager && params.UsePoSValidatorManager {
		return nil, fmt.Errorf("blockchain can not be both PoA and PoS")
	}
	if params.UsePoAValidatorManager {
		validatormanager.AddPoAValidatorManagerContractToAllocations(params.initialTokenAllocation)
		validatormanager.AddTransparentProxyContractToAllocations(params.initialTokenAllocation, proxyOwner)
	} else if params.UsePoSValidatorManager {
		validatormanager.AddPoSValidatorManagerContractToAllocations(params.initialTokenAllocation)
		validatormanager.AddTransparentProxyContractToAllocations(params.initialTokenAllocation, proxyOwner)
		validatormanager.AddRewardCalculatorToAllocations(params.initialTokenAllocation, rewardBasisPoints)
		params.enableNativeMinterPrecompile = true
	}

	if params.UseExternalGasToken {
		params.enableNativeMinterPrecompile = true
		params.nativeMinterPrecompileAllowList.AdminAddresses = append(
			params.nativeMinterPrecompileAllowList.AdminAddresses,
			common.HexToAddress(teleporterInfo.FundedAddress),
		)
	}

	genesisBlock0Timestamp := utils.TimeToNewUint64(time.Now())
	precompiles := getPrecompiles(params, genesisBlock0Timestamp)

	if params.UseTeleporter || params.UseExternalGasToken {
		addTeleporterAddressesToAllowLists(
			&precompiles,
			teleporterInfo.FundedAddress,
			teleporterInfo.MessengerDeployerAddress,
			teleporterInfo.RelayerAddress,
		)
	}

	subnetConfig, err := blockchainSDK.New(
		&blockchainSDK.SubnetParams{
			SubnetEVM: &blockchainSDK.SubnetEVMParams{
				ChainID:     new(big.Int).SetUint64(params.chainID),
				FeeConfig:   feeConfig,
				Allocation:  params.initialTokenAllocation,
				Precompiles: precompiles,
				Timestamp:   genesisBlock0Timestamp,
			},
			Name: "TestSubnet",
		})
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, subnetConfig.Genesis, "", "    ")
	if err != nil {
		return nil, err
	}

	return prettyJSON.Bytes(), nil
}

func someoneWasAllowed(allowList AllowList) bool {
	addrs := append(append(allowList.AdminAddresses, allowList.ManagerAddresses...), allowList.EnabledAddresses...)
	return len(addrs) > 0
}

func someAllowedHasBalance(allowList AllowList, allocations core.GenesisAlloc) bool {
	addrs := append(append(allowList.AdminAddresses, allowList.ManagerAddresses...), allowList.EnabledAddresses...)
	for _, addr := range addrs {
		// we can break at the first address that has a non-zero balance
		if bal, ok := allocations[addr]; ok &&
			bal.Balance != nil &&
			bal.Balance.Uint64() > uint64(0) {
			return true
		}
	}
	return false
}
