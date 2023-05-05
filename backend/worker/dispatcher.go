package worker

type Dispatcher func(jobsChan chan *Job, quitChan chan int)
