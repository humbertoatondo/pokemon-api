package api // import "github.com/humbertoatondo/pokemon-api/api"


TYPES

type App struct {
        Router *mux.Router
}
    App contains the necessary components for runnign the server. In this case
    it just stores the router.

func (app *App) Initialize()
    Initialize inits the app router and the routes.

func (app *App) Run(port string)
    Run starts listening and serving in a given port.