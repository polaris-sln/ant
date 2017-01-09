package app

type app struct {
	routes    Route
	addr      string
	maxJob    int
	maxWorker int
}

func (app *App)Init() {

}

func (app *App)Default() {

}

func (app *App)SetMaxJob(amount int) {
	App.maxJob = amount
}

func (app *App)SetMaxWorker(amount int) {
	App.maxWorker = amount
}


func (app *App)SetAddr(addr string) {
	App.addr = addr
}

func (app *App)SetRoutes(routes Route) {
	App.routes = routes;
}
