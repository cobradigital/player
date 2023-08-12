package router

import (
	"repo/flags"
	"repo/handler"
)

// routeAPI configure request routing in API. Handlers must be defined in handler package
func routeAPI(r Router) {

	r.HandleREST("/", handler.Health, flags.ACLEveryone).Methods("GET")

	// Player / User endpoint
	r.HandleREST("/register", handler.Register, flags.ACLEveryone).Methods("POST")
	r.HandleREST("/login", handler.Login, flags.ACLEveryone).Methods("POST")
	r.HandleREST("/logout", handler.Logout, flags.ACLAuthenticatedUser).Methods("GET")
	r.HandleREST("/profile", handler.GetProfile, flags.ACLAuthenticatedUser).Methods("GET")
	r.HandleREST("/profile", handler.PutProfile, flags.ACLAuthenticatedUser).Methods("PUT")

	// Admin Endpoint
	r.HandleREST("/player/list", handler.GetAll, flags.ACLAuthenticatedAdmin).Methods("GET")
	r.HandleREST("/player/deposit", handler.PostPlayerDeposit, flags.ACLAuthenticatedAdmin).Methods("POST")

}
