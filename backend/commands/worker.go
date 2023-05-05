package commands

import (
	"fmt"

	"bitbucket.org/andreychernih/tweemote/mb"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/worker"
)

// Creates dispatcher func which will be dispatching messages from the given Queue
func CreateDispatchFunc(mb mb.MessageBroker, queue string, prefetchCount int) (worker.Dispatcher, error) {
	ch, err := mb.Consume(queue, prefetchCount)
	if err != nil {
		return nil, err
	}

	return func(jobsChan chan *worker.Job, quitChan chan int) {
		for {
			select {
			case d := <-ch:
				//glog.Info("CONSUME: %s", string(d.GetMessage().Serialize()))
				j := worker.NewJob("job", d)
				jobsChan <- j
			case <-quitChan:
				return
			}
		}
	}, nil
}

// Creates dispatcher func which will be dispatching every single twitter account from the database
func CreateTwitterAccountsDispatcherFunc() worker.Dispatcher {
	return func(jobsChan chan *worker.Job, quitChan chan int) {
		var accounts []models.TwitterAccount
		models.GetAllActiveTwitterAccounts(&accounts)

		for _, account := range accounts {
			j := worker.NewJob(fmt.Sprintf("twitter-account-%s", account.TwitterUsername), account)
			jobsChan <- j
		}

		// Stop pool after processing
		close(jobsChan)
		close(quitChan)
	}
}
