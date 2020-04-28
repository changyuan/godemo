package main

//noinspection ALL
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
)

type App struct {
	Router *mux.Router
}

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

type shortlinkResp struct {
	Shortlink string `json:"shortlink"`
}

//initialize app

func (a *App) Initialize() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/shorten", a.createShortlink).Methods("POST")
	a.Router.HandleFunc("/api/info", a.getShortlinkInfo).Methods("GET")
	a.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,11}}", a.redirect).Methods("GET")
}

func (a *App) getShortlinkInfo(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	s := vals.Get("shortlink")
	fmt.Printf("%s\n", s)
}

func (a *App) createShortlink(writer http.ResponseWriter, request *http.Request) {
	var req shortenReq
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		return
	}
	err = validator.Validate(req)
	if err != nil {
		return
	}
	defer request.Body.Close()
	fmt.Printf("%v\n", req)
}

func (a *App) redirect(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Printf("%s\n", vars["shortlink"])
}

//Run

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
