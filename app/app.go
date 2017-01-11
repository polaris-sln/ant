package app

import (
	"fmt"
)

type Runner interface {
	Run()
	SetApp(app *App)
}

type App struct {
	maxJob    int
	maxWorker int
	jobChan   chan Job
	sigChan   chan SIG
	workerIds []int
	infoFile  string
	errFile   string
	logger    *Logger
}

func NewApp() *App {
	return &App{
		maxJob:     1000,
		maxWorker:  1,
		infoFile:   "",
		errFile:    "",
	        workerIds: []int{},
		logger:     nil,
		jobChan:    make(chan Job, 1000),
		sigChan:    make(chan SIG, 10),
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

func (app *App)log(lt int, msg string) {
	app.jobChan <-NewLogMsg(lt, msg, app.logger)
}

func (app *App)LogInfo(msg ...interface{}) {
	app.log(LOGINFO, fmt.Sprint(msg...))
}

func (app *App)LogErr(msg ...interface{}) {
	app.log(LOGERR, fmt.Sprint(msg...))
}

func (app *App)newWorker() {
	for i := 0; i < app.maxWorker; i++ {
		worker := NewWorker(app, i)
		app.workerIds = append(app.workerIds, i)
		go worker.Start()
	}

}

func (app *App)initLogger() {
	app.logger = NewLogger(app.infoFile, app.errFile)
}

func (app *App)Dispatch(job Job) {
	app.jobChan <- job
}

func (app *App)Run(runner Runner) {
	app.initLogger()
	go app.newWorker()
	runner.SetApp(app)
	runner.Run()
}
