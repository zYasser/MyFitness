package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/zYasser/MyFitness/config"
	"github.com/zYasser/MyFitness/middleware"
	"github.com/zYasser/MyFitness/models"
	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
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

	program_router := app.Router.PathPrefix("/programs").Subrouter()
	program_router.HandleFunc("", app.createProgram).Methods(http.MethodPost)
	program_router.HandleFunc("/all", app.getAllPrograms).Methods(http.MethodGet)
	program_router.HandleFunc("", app.getProgramById).Queries("id", "{id}")

}

func InitApplication() *Application {
	application := &Application{Db: config.InitDatabase(), Router: mux.NewRouter(), Redis: config.InitRedis()}
	application.initRouter()
	models.Migration(application.Db)

	return application
}

type Application struct {
	Router *mux.Router
	Db     *gorm.DB
	Redis  *redis.Client
	Logger *utils.Logger
}
