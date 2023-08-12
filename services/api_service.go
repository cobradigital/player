package services

import (
	"repo/loggers"
	"repo/repositories"
)

// Logger
var log = loggers.Get()

var Auth AuthService
var Players PlayersService

// Init ...
func Init() {
	Auth = &authService{repositories.Players}
	Players = &playersService{repositories.Players, repositories.Deposit}
}
