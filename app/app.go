package app

import (
	"fmt"
)

type App struct {
	routes    Route
	addr      string
	maxJob    int
	maxWorker int
	jobChan   chan Job
	sigChan   chan SIG
	workerIds []int{}
	infoFile  string
	errFile   string
	logger    *Logger
}

func NewApp() App {
	return App{
		addr:      ":8080",
		maxJob:    1000,
		maxWorker: 1,
		infoFile:  "",
		errFilei:  "",
		jobChan:   make(Job chan, 1000),
		routes:    nil,
		logger:    nil,
	}
}


func (app *App)InitJobChan() {
	app.jobChan = make(chan Job, app.maxJob)
}

func (app *App)SetLogFile(fileName ...string) {
	paraLen := len(fileName)
	if paraLen >= 1 {
		app.infoFile = fileName[0]
	}

	if paraLen == 2 {
		app.errFile = fileName[1]
	}

}

func (app *App)SetMaxJob(amount int) {
	app.maxJob = amount
}

func (app *App)SetMaxWorker(amount int) {
	app.maxWorker = amount
}

func (app *App)SetAddr(addr string) {
	app.addr = addr
}

func (app *App)SetRoutes(routes Route) {
	app.routes = routes
}

func (app *App)log(lt int, msg string) {
	app.jobChan <-NewLogMsg(LOGINFO, fmt.Sprint(msg...) app.logger)
}

func (app *App)LogInfo(msg ...string) {
	app.log(LOGINFO, fmt.Sprint(msg...))
}

func (app *App)LogErr(msg ...string) {
	app.log(LOGERR, fmt.Sprint(msg...))
}

func (app *App)newWorker() {
	for i := 0; i < app.maxWorker; i++ {
		worker := NewWorker(app, i)
		app.workerIds = append(app.workerIds, i)
		go worker.Start()
	}

}

func (app *App)dispatch(job Job) {
	app.jobChan <- job
}

func (app *App)Run() {

}
