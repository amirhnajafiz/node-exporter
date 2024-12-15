package pbft

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/config"
	"github.com/F24-CSE535/2pcbyz/cluster/internal/pbft/memory"

	"go.uber.org/zap"
)

// StateMachine runs PBFT protocol.
type StateMachine struct {
	Cluster   string
	Ipt       *config.IPTable
	Consensus chan context.Context
	Queue     chan context.Context
	Logger    *zap.Logger

	sequence  int
	block     bool
	byzantine bool
	memory    *memory.SharedMemory
}

func (sm *StateMachine) Start() {
	sm.block = false
	sm.byzantine = false
	sm.sequence = 0
	sm.memory = memory.NewSharedMemory()

	for {
		// get context messages from queue
		ctx := <-sm.Consensus
		payload := ctx.Value("request")

		// create an error variable for handlers result
		var err error

		// map of method to handler
		switch ctx.Value("method").(string) {
		case "request":
			if sm.block || sm.byzantine {
				continue
			}
			err = sm.request(payload)
		case "input":
			if sm.block || sm.byzantine {
				continue
			}
			err = sm.input(payload)
		case "preprepare":
			if sm.block || sm.byzantine {
				continue
			}
			err = sm.prePrepare(payload)
		case "pbft.ackpreprepare":
			if sm.block || sm.byzantine {
				continue
			}
			err = sm.ackPrePrepare(payload)
		case "pbft.prepare":
			if sm.block || sm.byzantine {
				continue
			}
			err = sm.prepare(payload)
		case "pbft.ackprepare":
			if sm.block || sm.byzantine {
				continue
			}
			err = sm.ackPrepare(payload)
		case "pbft.commit":
			if sm.block || sm.byzantine {
				continue
			}
			err = sm.commit(payload)
		case "pbft.timeout":
			err = sm.timeout(payload)
		case "block":
			sm.block = true
		case "unblock":
			sm.block = false
		case "byzantine":
			sm.byzantine = true
		case "nonbyzantine":
			sm.byzantine = false
		default:
			sm.Queue <- ctx
		}

		// check error
		if err != nil {
			sm.Logger.Warn("state-machine error", zap.Error(err))
		}
	}
}
