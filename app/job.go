package app

type Job interface {
	Do()
	SetApp(app *App)
}

