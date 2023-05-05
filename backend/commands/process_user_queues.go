package commands

import (
	"time"

	"github.com/golang/glog"

	"bitbucket.org/andreychernih/tweemote/errors"
	"bitbucket.org/andreychernih/tweemote/mb"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/worker"
)

const ProcessUserQueuesDelay = 120 * time.Second

type actionsConsumer struct {
	worker.SyncContext

	broker           mb.MessageBroker
	twitterAccountID uint
	isRunning        bool
}

func (ac actionsConsumer) start() error {
	ac.WaitGroup.Add(1)

	queue := actionQueueName(ac.twitterAccountID)
	ch, err := ac.broker.Consume(queue, 1)
	if err != nil {
		return err
	}

	go func() {
		defer ac.WaitGroup.Done()

		for {
			select {
			case d := <-ch:
				ac.broker.Publish(mb.ActionsQueueName, d.GetMessage())
				d.Ack()
				time.Sleep(ProcessUserQueuesDelay)
			case <-ac.QuitChan:
				return
			}
		}
	}()

	return nil
}

type ProcessUserQueuesCommand struct {
	Meta

	worker.SyncContext

	broker    mb.MessageBroker
	consumers map[uint]actionsConsumer
}

func (cmd ProcessUserQueuesCommand) Synopsis() string {
	return "Consumes from all user action queues and add necessary delay before pushing it to global actions queue"
}

func (cmd ProcessUserQueuesCommand) Help() string {
	return cmd.Synopsis()
}

func (cmd ProcessUserQueuesCommand) syncQueues() {
	glog.Info("Syncing queues")

	db := models.Connect()
	var ids []uint
	db.Find(&models.TwitterAccount{}).Pluck("id", &ids)

	for _, id := range ids {
		_, ok := cmd.consumers[id]
		if ok {
		} else {
			cmd.consumers[id] = actionsConsumer{
				broker:           cmd.broker,
				twitterAccountID: id,
				SyncContext: worker.SyncContext{
					WaitGroup: cmd.WaitGroup,
					QuitChan:  cmd.QuitChan,
				},
			}
			err := cmd.consumers[id].start()
			errors.Check(err)
		}
	}
}

func (cmd ProcessUserQueuesCommand) Run(args []string) int {
	cmd.QuitChan = make(chan struct{})
	cmd.consumers = make(map[uint]actionsConsumer)
	cmd.Trap()

	broker, err := mb.NewConnection(persistThreads)
	errors.Check(err)
	cmd.broker = broker
	defer cmd.broker.Disconnect()

	ticker := time.NewTicker(60 * time.Second)
	cmd.syncQueues()

	cmd.WaitGroup.Add(1)
	go func() {
		defer cmd.WaitGroup.Done()

		for {
			select {
			case <-ticker.C:
				cmd.syncQueues()
			case <-cmd.QuitChan:
				glog.Info("Exiting")
				ticker.Stop()
				return
			}
		}
	}()

	cmd.WaitGroup.Wait()

	return 0
}
