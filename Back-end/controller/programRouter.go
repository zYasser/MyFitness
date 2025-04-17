package controller

import (
	"net/http"
	"strconv"

	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/middleware"
	"github.com/zYasser/MyFitness/service"
	"github.com/zYasser/MyFitness/utils"
)

func (app *Application) createProgram(w http.ResponseWriter, r *http.Request) {
	con := r.Context()
	logger := middleware.FromContext(con)
	logger.InfoLog.Println("Received Create Exercise Request")
	var params dto.Program
	if err := utils.FromJSON(&params, r.Body); err != nil {
		logger.ErrorLog.Printf("Failed to serialize this object:%v", r.Body)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	if errs := utils.Validate(&params, validate); len(errs) != 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, errs)
		return
	}
	err := service.InsertProgram(app.Db, logger, params)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, nil)

}
func (app *Application) insertWorkoutToProgram(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) getAllPrograms(w http.ResponseWriter, r *http.Request) {
	con := r.Context()
	logger := middleware.FromContext(con)

	logger.InfoLog.Println("Received Get All Program Request")
	result := service.GetAllProgram(app.Db, logger)
	utils.RespondWithJSON(w, http.StatusOK, result)

}
func (app *Application) getProgramById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	if query == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "id query can not be empty"})
		return
	}
	con := r.Context()
	logger := middleware.FromContext(con)

	id, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		logger.ErrorLog.Printf("Failed to parse getProgramById Query error:%v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "id query can not be empty"})
		return

	}

	result, _ := service.GetProgramById(app.Db, logger, id)
	utils.RespondWithJSON(w, http.StatusOK, result)

}
