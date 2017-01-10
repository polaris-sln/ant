package app


type Worker struct {
	id      int
	jobChan chan Job
	sigChan chan SIG
	logger  *Logger
	quit    chan bool
}

func NewWorker(app *App, id int) Worker {
	return Worker{
		id:      id,
		jobChan: app.jobChan,
		sigChan: app.sigChan,
		logger:  app.logger,
	        quit:    make(chan bool),
	}
}

func (worker *Worker)Start() {
	for {
		select {
			case job := <-worker.jobChan:
				job.Do()
			case sig := <-worker.sigChan:
				worker.SigHandler(sig)
			case <-worker.quit:
				return
		}
	}
}

func (worker *Worker)Stop() {
	worker.quit <- true
}

func (worker *Worker)Log(logType int, msg string) {
	worker.jobChan <- NewLogMsg(logType, msg, worker.logger)
}

func (worker *Worker)SigHandler(signal SIG) {
	switch signal {
	case SIGQUIT:
		worker.Stop()
	default:
		worker.Log(LOGERR, "Unrecognized signal")
	}
}

