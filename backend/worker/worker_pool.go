package worker

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/golang/glog"

	"bitbucket.org/andreychernih/tweemote/models"
)

const deadCheckSeconds = 10
const considerDeadInactiveSeconds = 120

type WorkerPool struct {
	name       string
	numWorkers int
	queue      chan *Job
	wg         sync.WaitGroup
	workersWg  sync.WaitGroup
	handler    Handler
	dispatcher Dispatcher
	quitChan   chan int
	stopped    bool
}

func NewWorkerPool(name string, numWorkers int, dFunc Dispatcher, wFunc Handler) *WorkerPool {
	return &WorkerPool{
		name:       name,
		numWorkers: numWorkers,
		queue:      make(chan *Job),
		quitChan:   make(chan int),
		handler:    wFunc,
		dispatcher: dFunc,
		stopped:    false,
	}
}

func (wp *WorkerPool) Start() {
	// Delete stale workers
	wp.wg.Add(1)
	go wp.reap()

	// Dispatch messages
	wp.wg.Add(1)
	go wp.dispatch()

	// Handle kill signals
	wp.wg.Add(1)
	go wp.trap()

	// Start workers
	for i := 0; i < wp.numWorkers; i++ {
		worker := NewWorker(fmt.Sprintf("%s-%d", wp.name, i), wp.name, wp.queue, wp.handler)

		wp.workersWg.Add(1)
		go worker.start(&wp.workersWg, wp.quitChan)
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.queue)
	close(wp.quitChan)
}

func (wp *WorkerPool) Wait() {
	wp.workersWg.Wait()
	wp.wg.Wait()
}

func (wp *WorkerPool) dispatch() {
	defer glog.Info("Exiting dispatcher")
	defer wp.wg.Done()

	glog.Info("Booted dispatcher")
	wp.dispatcher(wp.queue, wp.quitChan)
}

func (wp *WorkerPool) trap() {
	defer glog.Infof("Exiting trapper")
	defer wp.wg.Done()

	glog.Info("Booted trapper")

	// Trap signals and stop worker pools
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)

	hardKill := false

	for {
		select {
		case <-wp.quitChan:
			return
		case s := <-sc:
			if hardKill {
				glog.Infof("Received signal %s for a second time, exiting with exit code 1", s)
				os.Exit(1)
			} else {
				glog.Infof("Received signal %s, trying to exit gracefully", s)
				wp.Stop()
				hardKill = true
			}
		}
	}
}

func (wp *WorkerPool) reap() {
	defer wp.wg.Done()
	defer glog.Info("Exiting reaper")

	glog.Info("Booted reaper")

	ticker := time.NewTicker(deadCheckSeconds * time.Second)
	for {
		select {
		case <-wp.quitChan:
			return
		case <-ticker.C:
			db := models.Connect()
			db.Where("last_active_at < ?", time.Now().Add(-considerDeadInactiveSeconds*time.Second)).Delete(models.Worker{})
		}
	}

}
