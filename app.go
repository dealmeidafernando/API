package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

}

// Run ...
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/jobs", a.getJobs).Methods("GET")
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

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
