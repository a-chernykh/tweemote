package worker

import (
	"os"
	"sync"
	"time"

	"github.com/golang/glog"

	"bitbucket.org/andreychernih/tweemote/models"
)

const heartbeatPeriodSeconds = 5

type Worker struct {
	name        string
	worker_type string
	queue       chan *Job
	model       *models.Worker
	handler     Handler
}

func NewWorker(name string, wt string, queue chan *Job, handler Handler) *Worker {
	return &Worker{
		name:        name,
		worker_type: wt,
		queue:       queue,
		handler:     handler,
	}
}

func (w *Worker) GetID() uint {
	return w.model.ID
}

func (w *Worker) GetName() string {
	return w.name
}

func (w *Worker) start(wg *sync.WaitGroup, quitChan chan int) {
	defer wg.Done()

	if err := w.register(); err != nil {
		panic(err)
	}

	glog.Infof("Booted worker %s\n", w.name)

	go func() {
		ticker := time.NewTicker(heartbeatPeriodSeconds * time.Second)
		for {
			select {
			case <-ticker.C:
				w.ping()
			case <-quitChan:
				return
			}
		}
	}()

	for j := range w.queue {
		w.handler(w, j, quitChan)
	}

	if err := w.deregister(); err != nil {
		panic(err)
	}

	glog.Infof("Exited worker %s\n", w.name)
}

func (w *Worker) register() error {
	h, err := os.Hostname()
	if err != nil {
		return err
	}

	w.model = &models.Worker{Name: w.name,
		WorkerType:   w.worker_type,
		Hostname:     h,
		StartedAt:    time.Now(),
		LastActiveAt: time.Now()}

	db := models.Connect()
	if err := db.Create(&w.model).Error; err != nil {
		return err
	}

	return nil
}

func (w *Worker) ping() error {
	db := models.Connect()
	w.model.LastActiveAt = time.Now()
	return db.Save(&w.model).Error
}

func (w *Worker) deregister() error {
	db := models.Connect()
	return db.Delete(&w.model).Error
}
