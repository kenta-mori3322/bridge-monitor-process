package types

import "github.com/humansdotai/bridge-monitor-process/common"

type ErrataBlock struct {
	Height int64
	Txs    []ErrataTx
}

type ErrataTx struct {
	TxID  common.TxID
	Chain common.Chain
}
