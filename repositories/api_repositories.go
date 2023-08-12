package repositories

var Players PlayerRepository
var Deposit PlayerDepositRepository

// Init ...
func Init() {
	Config()
	Players = initPlayerRepository()
	Deposit = initPlayerDepositRepository()
}
