package worker

type Handler func(w *Worker, j *Job, quitChan chan int)
