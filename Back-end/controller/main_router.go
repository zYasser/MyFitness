package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zYasser/MyFitness/repository"
	"github.com/zYasser/MyFitness/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var logger=utils.GetLogger()
// type Router struct{

// }

func initRouter() *mux.Router{
	router:=mux.NewRouter()

	router.HandleFunc("/test" , func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello"))	
	})	
	return router
}
func initDatabase()*gorm.DB{
	DATABASE_USERNAME:=utils.GetEnv("DATABASE_USER")
	DATABASE_PASSWORD:=utils.GetEnv("DATABASE_PASSWORD")
	fmt.Printf(DATABASE_PASSWORD, "\n " , DATABASE_USERNAME)
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=MyFitness port=5432 sslmode=disable" ,DATABASE_USERNAME,DATABASE_PASSWORD)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if(err !=nil){
		logger.ErrorLog.Fatalf("Connection Failed : %v" , err)
	}
	logger.InfoLog.Println("Connection has been established")
	
	return db
	
}

func InitApplication()*Application{
	application:=&Application{db: initDatabase() , Router:initRouter() }
	repository.Migration(application.db)
	return application
}

type Application struct{
	Router *mux.Router
	db *gorm.DB
}

