package repositories

import (
	"repo/models"
)

// PlayerDepositeRepository ...
type PlayerDepositRepository interface {
	GetOne(where string, args ...interface{}) (float64, error)
	Create(data models.PlayersDeposit) error
}

// playerDepositRepository ...
type playerDepositRepository struct{}

func initPlayerDepositRepository() PlayerDepositRepository {
	// Prepare statements
	var r playerDepositRepository
	return &r
}

func (r *playerDepositRepository) GetOne(where string, args ...interface{}) (float64, error) {
	var data float64
	err := DBConnect.Table("players_deposit").Select("COALESCE(SUM(nominal), 0) as nominal").Where(where, args...).Find(&data).Error
	return data, err
}

func (r *playerDepositRepository) Create(data models.PlayersDeposit) error {
	data.ID = GenerateUUID()
	err := DBConnect.Table("players_deposit").Create(&data).Error
	return err
}
