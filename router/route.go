package router

import (
	"fmt"
	"net/http"
	"time"

	"repo/flags"
	"repo/response"

	"github.com/gorilla/mux"

	"repo/configs"
	"repo/handler"
	"repo/loggers"
	"repo/middleware"
)

// Logger
var log = loggers.Get()

// A Router extends gorilla mux.Router functionality to handle RESTFunc
type Router struct {
	*mux.Router
}

// Auth middleware
var Auth = middleware.Auth()

// Init start time
var startTime time.Time

// HandleREST authenticated request and execute RESTFunc type handler
func (r Router) HandleREST(path string, fn handler.RESTFunc, purpose string) *mux.Route {
	var h http.Handler

	if purpose != flags.ACLEveryone {
		h = Auth(fn, purpose)
	} else {
		h = fn
	}
	return r.NewRoute().Path(path).Handler(h)
}

// New creates new router instance and configure api routing by calling routeAPI() function
func New(start time.Time) Router {

	// Set start time
	startTime = start
	// Create new router
	r := Router{mux.NewRouter()}

	// Get base url
	var fs = http.FileServer(http.Dir(configs.MustGetString("file.images")))
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", fs)).Methods("GET")

	baseURL := configs.MustGetString("server.base_url")
	log.Infof("API Base URL: %s", baseURL)
	// Init api router
	a := Router{r.PathPrefix(baseURL).Subrouter()}
	routeAPI(a)
	// Set error handler
	r.NotFoundHandler = handler.RESTFunc(handler.NotFound)
	r.MethodNotAllowedHandler = handler.RESTFunc(handler.MethodNotAllowed)
	// Set main handler
	r.HandleREST(baseURL, GetAppStatus, flags.ACLEveryone).Methods("GET")
	// Return main router
	return r
}

// GetAppStatus ...
func GetAppStatus(_ *http.Request) (*response.Success, error) {
	body := AppStatus{
		BuildVersion: flags.AppVersion,
		Uptime:       fmt.Sprintf("%s", time.Since(startTime)),
	}
	return response.NewSuccess(&body), nil
}

// AppStatus ...
type AppStatus struct {
	BuildVersion string `json:"build_version"`
	Uptime       string `json:"uptime"`
}
