package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App :)
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// InitializeConection :)
func (a *App) InitializeConection(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

// Run ...
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/jobs", a.getJobs).Methods("GET")
	a.Router.HandleFunc("/jobs/{id:[0-9]+}", a.getJobID).Methods("GET")
	a.Router.HandleFunc("/jobs", a.createJob).Methods("POST")
	a.Router.HandleFunc("/jobs/{id:[0-9]+}", a.updateJob).Methods("PUT")
	a.Router.HandleFunc("/jobs/{id:[0-9]+}", a.deleteJob).Methods("DELETE")
}

func (a *App) getJobs(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	jobs, err := getJobs(a.DB, start, count)
	if err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, http.StatusOK, jobs)
}

func (a *App) getJobID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid job ID")
		return
	}

	u := job{ID: id}
	if err := u.getJobID(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Job not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) createJob(w http.ResponseWriter, r *http.Request) {
	var u job
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		log.Fatal(err)
		return
	}
	defer r.Body.Close()

	if err := u.createJob(a.DB); err != nil {
		log.Fatal(err)
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) updateJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid job ID")
		return
	}
	var u job
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	u.ID = id

	if err := u.updateJob(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) deleteJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Job ID")
		return
	}

	u := job{ID: id}
	if err := u.deleteJob(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
