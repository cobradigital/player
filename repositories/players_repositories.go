package repositories

import (
	"repo/flags"
	"repo/models"
	"repo/query"
)

// PlayerRepository ...
type PlayerRepository interface {
	GetAll(filters map[string]query.Filter, limit, skip int, havingDeposit, orderby string) ([]models.PlayersList, int64, error)
	GetOne(where string, args ...interface{}) (models.Players, error)
	Create(data models.Players) error
	Update(data models.Players, where string, args ...interface{}) error
	Filter() map[string]query.Filter
}

// playerRepository ...
type playerRepository struct{}

func initPlayerRepository() PlayerRepository {
	// Prepare statements
	var r playerRepository
	return &r
}

func (r *playerRepository) GetAll(filters map[string]query.Filter, limit, skip int, havingDeposit, orderby string) ([]models.PlayersList, int64, error) {
	var (
		count int64
		list  []models.PlayersList
	)

	b := query.NewFilterBuilder(filters)
	filterQuery, filterArgs := b.Build()

	join := `
	left join
	(SELECT 
	player_id, 
	(SELECT COALESCE(SUM(nominal), 0) FROM players_deposit WHERE player_id = player_id AND type = 'debit' ) as debit, 
	(SELECT COALESCE(SUM(nominal), 0) FROM players_deposit WHERE player_id = player_id AND type = 'credit' ) as credit
	FROM players_deposit GROUP BY player_id)
	as pd on p.id = pd.player_id
	`
	query := DBConnect.Table("players as p").Joins(join).Group("p.id, pd.debit, pd.credit")
	query = query.Select("p.id,p.username,p.password,p.email,p.bank,p.nama_rekening,p.no_rekening, pd.debit, pd.credit, SUM(pd.debit-pd.credit) as deposit, p.created_at,p.updated_at")
	queryCount := DBConnect.Table("players as p").Joins(join).Group("p.id, pd.debit, pd.credit")

	if len(filterArgs) > 0 {
		query = query.Where(filterQuery, filterArgs...)
		queryCount = queryCount.Where(filterQuery, filterArgs...)
	}

	if orderby != "" {
		query = query.Order(orderby)
	}

	if havingDeposit != "" {
		query = query.Having("SUM(pd.debit-pd.credit) > ?", havingDeposit)
		queryCount = queryCount.Having("SUM(pd.debit-pd.credit) > ?", havingDeposit)
	}

	err := query.Limit(limit).Offset(skip).Find(&list).Error
	if err != nil {
		return list, count, err
	}

	err = queryCount.Count(&count).Error
	if err != nil {
		return list, count, err
	}

	return list, count, err
}

func (r *playerRepository) GetOne(where string, args ...interface{}) (models.Players, error) {
	var data models.Players
	err := DBConnect.Table("players").Where(where, args...).Find(&data).Error
	return data, err
}

func (r *playerRepository) Create(data models.Players) error {
	data.ID = GenerateUUID()
	err := DBConnect.Table("players").Create(&data).Error
	return err
}

func (r *playerRepository) Update(data models.Players, where string, args ...interface{}) error {
	err := DBConnect.Model(models.Players{}).Where(where, args...).Updates(data).Error
	return err
}

func (r *playerRepository) Filter() map[string]query.Filter {
	// Init available filters for menu discovery
	return map[string]query.Filter{
		// Data Vendor Status
		flags.FilterUsername:       query.NewFilter(query.FilterStringLike, query.ClauseWhere, "username LIKE ?"),
		flags.FilterBank:           query.NewFilter(query.FilterStringMatch, query.ClauseWhere, "bank = ?"),
		flags.FilterNamaRekening:   query.NewFilter(query.FilterStringLike, query.ClauseWhere, "nama_rekening LIKE ?"),
		flags.FilterNoRekening:     query.NewFilter(query.FilterStringMatch, query.ClauseWhere, "no_rekening = ?"),
		flags.FilterStartCreatedAt: query.NewFilter(query.FilterStringMatch, query.ClauseWhere, "to_date(p.created_at::TEXT,'YYYY-MM-DD') > ?"),
		flags.FilterEndCreatedAt:   query.NewFilter(query.FilterStringMatch, query.ClauseWhere, "to_date(p.created_at::TEXT,'YYYY-MM-DD') < ?"),
	}
}
