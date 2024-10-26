package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zYasser/MyFitness/config"
	"github.com/zYasser/MyFitness/middleware"
	"github.com/zYasser/MyFitness/service"
)

// var logger = utils.GetLogger()
// type Router struct{

// }

func (app *Application) initRouter() {
	app.Router.Use(middleware.LoggingMiddleware)

	userRouter := app.Router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/register", app.register).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", app.login).Methods(http.MethodPost)
	exercise_router := app.Router.PathPrefix("/exercise").Subrouter()
	// exercise_router.Use(middleware.AuthorizationMiddleware(app.Redis))
	exercise_router.HandleFunc("", app.createExercise).Methods(http.MethodPost)
	exercise_router.HandleFunc("", app.fetchAllExercises).Methods(http.MethodGet)
	exercise_router.HandleFunc("/{id}", app.getExerciseFromId).Methods(http.MethodGet)
}

func InitApplication() *Application {
	application := &Application{Router: mux.NewRouter(), ApplicationConfig: &ApplicationConfig{
		Redis: config.InitRedis(),
		Db:    config.InitDatabase(),
	}}
	application.initRouter()
	service.Migration(application.ApplicationConfig.Db)

	return application
}
