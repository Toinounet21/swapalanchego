// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"github.com/Toinounet21/swapalanchego/ids"
	"github.com/Toinounet21/swapalanchego/snow/consensus/snowman"
	"github.com/Toinounet21/swapalanchego/vms/proposervm/indexer"
)

var _ indexer.BlockServer = &VM{}

// GetFullPostForkBlock implements BlockServer interface
// Note: this is a contention heavy call that should be avoided
// for frequent/repeated indexer ops
func (vm *VM) GetFullPostForkBlock(blkID ids.ID) (snowman.Block, error) {
	vm.ctx.Lock.Lock()
	defer vm.ctx.Lock.Unlock()

	return vm.getPostForkBlock(blkID)
}

// Commit implements BlockServer interface
func (vm *VM) Commit() error {
	vm.ctx.Lock.Lock()
	defer vm.ctx.Lock.Unlock()

	return vm.db.Commit()
}
