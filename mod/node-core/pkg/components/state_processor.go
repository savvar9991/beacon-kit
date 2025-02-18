// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package components

import (
	"cosmossdk.io/depinject"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/execution/pkg/engine"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components/metrics"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core"
)

// StateProcessorInput is the input for the state processor for the depinject
// framework.
type StateProcessorInput[
	LoggerT log.AdvancedLogger[LoggerT],
	ExecutionPayloadT ExecutionPayload[
		ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT,
	],
	ExecutionPayloadHeaderT ExecutionPayloadHeader[ExecutionPayloadHeaderT],
	DepositT Deposit[DepositT, *ForkData, WithdrawalCredentials],
	WithdrawalT Withdrawal[WithdrawalT],
	WithdrawalsT Withdrawals[WithdrawalT],
] struct {
	depinject.In
	Logger          LoggerT
	ChainSpec       common.ChainSpec
	ExecutionEngine *engine.Engine[
		ExecutionPayloadT,
		*engineprimitives.PayloadAttributes[WithdrawalT],
		PayloadID,
		WithdrawalsT,
	]
	DepositStore  DepositStore[DepositT]
	Signer        crypto.BLSSigner
	TelemetrySink *metrics.TelemetrySink
}

// ProvideStateProcessor provides the state processor to the depinject
// framework.
func ProvideStateProcessor[
	LoggerT log.AdvancedLogger[LoggerT],
	BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT, BeaconBlockHeaderT],
	BeaconBlockBodyT BeaconBlockBody[
		BeaconBlockBodyT, *AttestationData, DepositT,
		*Eth1Data, ExecutionPayloadT, *SlashingInfo,
	],
	BeaconBlockHeaderT BeaconBlockHeader[BeaconBlockHeaderT],
	BeaconStateT BeaconState[
		BeaconStateT, BeaconBlockHeaderT, BeaconStateMarshallableT,
		*Eth1Data, ExecutionPayloadHeaderT, *Fork, KVStoreT, *Validator,
		Validators, WithdrawalT,
	],
	BeaconStateMarshallableT any,
	DepositT Deposit[DepositT, *ForkData, WithdrawalCredentials],
	DepositStoreT DepositStore[DepositT],
	ExecutionPayloadT ExecutionPayload[
		ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT,
	],
	ExecutionPayloadHeaderT ExecutionPayloadHeader[ExecutionPayloadHeaderT],
	KVStoreT BeaconStore[
		KVStoreT, BeaconBlockHeaderT, *Eth1Data, ExecutionPayloadHeaderT,
		*Fork, *Validator, Validators, WithdrawalT,
	],
	WithdrawalsT Withdrawals[WithdrawalT],
	WithdrawalT Withdrawal[WithdrawalT],
](
	in StateProcessorInput[
		LoggerT,
		ExecutionPayloadT, ExecutionPayloadHeaderT,
		DepositT, WithdrawalT, WithdrawalsT,
	],
) *core.StateProcessor[
	BeaconBlockT, BeaconBlockBodyT, BeaconBlockHeaderT,
	BeaconStateT, *Context, DepositT, *Eth1Data, ExecutionPayloadT,
	ExecutionPayloadHeaderT, *Fork, *ForkData, KVStoreT, *Validator,
	Validators, WithdrawalT, WithdrawalsT, WithdrawalCredentials,
] {
	return core.NewStateProcessor[
		BeaconBlockT,
		BeaconBlockBodyT,
		BeaconBlockHeaderT,
		BeaconStateT,
		*Context,
		DepositT,
		*Eth1Data,
		ExecutionPayloadT,
		ExecutionPayloadHeaderT,
		*Fork,
		*ForkData,
		KVStoreT,
		*Validator,
		Validators,
		WithdrawalT,
		WithdrawalsT,
		WithdrawalCredentials,
	](
		in.Logger.With("service", "state-processor"),
		in.ChainSpec,
		in.ExecutionEngine,
		in.DepositStore,
		in.Signer,
		crypto.GetAddressFromPubKey,
		in.TelemetrySink,
	)
}
