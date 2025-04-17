package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/mapper"
	"github.com/zYasser/MyFitness/service"
	"github.com/zYasser/MyFitness/utils"
)

func (app *Application) createExercise(w http.ResponseWriter, r *http.Request) {
	app.Logger.InfoLog.Println("Received Create Exercise Request")
	var params dto.Exercise
	if err := utils.FromJSON(&params, r.Body); err != nil {
		app.Logger.ErrorLog.Printf("Failed to serialize this object:%v", r.Body)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return

	}
	if errs := utils.Validate(&params, validate); len(errs) != 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, errs)
		return
	}

	exercise := mapper.DtoToExercise(params)
	service.InsertExercise(exercise, app.Db, app.Logger)
	utils.RespondWithJSON(w, http.StatusCreated, exercise)

}

func (app *Application) getExerciseFromId(w http.ResponseWriter, r *http.Request) {
	app.Logger.InfoLog.Println("Received Create Exercise Request")
	vars := mux.Vars(r)
	exerciseId := vars["id"]
	result, err := service.GetExerciseById(exerciseId, app.Db, app.Logger)
	if err == nil {
		utils.RespondWithJSON(w, http.StatusOK, result)

	} else {
		utils.RespondWithJSON(w, err.StatusCode, err)
	}

}

func (app *Application) fetchAllExercises(w http.ResponseWriter, r *http.Request) {
	app.Logger.InfoLog.Println("fetch All Exercises Request")
	exercises := service.GetAllExercise(app.Db, app.Logger)
	utils.RespondWithJSON(w, http.StatusOK, exercises)

}
