package worker

import (
	"os"
	"os/signal"
	"sync"
)

type SyncContext struct {
	WaitGroup sync.WaitGroup
	QuitChan  chan struct{}
}

func (sc *SyncContext) Trap() {
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)

		_ = <-ch
		close(sc.QuitChan)
	}()
}
