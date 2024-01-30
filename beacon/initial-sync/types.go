// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package initialsync

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/itsdevbear/bolaris/beacon/execution"
	"github.com/itsdevbear/bolaris/types/state"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
)

// ethClient is an interface that wraps the ChainSyncReader from the go-ethereum package.
type ethClient interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*ethtypes.Header, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*ethtypes.Header, error)
}

// BeaconStateProvider defines an interface for providing a ForkChoiceStore.
type BeaconStateProvider interface {
	// ForkChoiceStore retrieves the ForkChoiceStore from the provided context.
	BeaconState(ctx context.Context) state.BeaconState
}

// executionService defines an interface for interacting with the execution client.
type executionService interface {
	// NotifyForkchoiceUpdate notifies the execution client of a forkchoice update.
	NotifyForkchoiceUpdate(
		ctx context.Context, slot primitives.Slot,
		arg *execution.NotifyForkchoiceUpdateArg, withAttrs, withRetry, async bool,
	) error
}
