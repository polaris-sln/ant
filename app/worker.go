package app


type Worker struct {
	id      int
	jobChan chan Job
	sigChan chan SIG
	quit    chan bool
}

func NewWorker(app App, id int) Worker {
	return Worker{
		id:      id
		jobChan: app.JobChan,
		sigChan: app.sigChan,
	        quit:    make(chan bool)
	}
}

func (worker *Worker)Start() {
	for {
		select {
			case job := <-worker.jobChan:
				job.Do()
			case sig := <-worker.sigChan:
				worker.SigHandler(sig)
			case <-quit:
				return
		}
	}
}

func (worker *Worker)Stop() {
	worker.quit <- true
}

func (worker *Worker)SigHandler(signal int) {
	switch sig {
	case SIGQUIT:
		worker.Stop()
	case default:
		fmt.FPrint(os.stderr, "Unrecognized signal\n")
	}
}

