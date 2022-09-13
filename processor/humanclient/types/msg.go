package types

import (
	stypes "github.com/humansdotai/bridge-monitor-process/x/humans/types"
)

type Msg struct {
	Type  string                    `json:"type"`
	Value stypes.MsgObservationVote `json:"value"`
}
