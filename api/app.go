package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/test", app.test).Methods("GET")
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(200)
	w.Write([]byte("Test successfull!"))
}
