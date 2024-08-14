package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/zYasser/MyFitness/config"
	"github.com/zYasser/MyFitness/middleware"
	"github.com/zYasser/MyFitness/service"
	"gorm.io/gorm"
)

// var logger = utils.GetLogger()
// type Router struct{

// }

func (app *Application) initRouter() {
	app.Router.Use(middleware.LoggingMiddleware)
	userRouter := app.Router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/register" , app.register).Methods(http.MethodPost)
	userRouter.HandleFunc("/login" , app.login).Methods(http.MethodPost)
	exercise_router := app.Router.PathPrefix("/exercise").Subrouter()
	exercise_router.Use(middleware.AuthorizationMiddleware(app.Redis))
	exercise_router.HandleFunc( "",app.createExercise).Methods(http.MethodPost)
	exercise_router.HandleFunc( "",app.fetchAllExercises).Methods(http.MethodGet)
}	


func InitApplication()*Application{
	application:=&Application{Db: config.InitDatabase() , Router:mux.NewRouter() ,Redis: config.InitRedis()}
	application.initRouter()
	service.Migration(application.Db)

	return application
}

type Application struct{
	Router *mux.Router
	Db *gorm.DB
	Redis *redis.Client
}

