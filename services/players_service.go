package services

import (
	"fmt"
	"repo/components/redis"
	"repo/crypto"
	"repo/flags"
	"repo/helper"
	"repo/models"
	"repo/query"
	"repo/repositories"
	"repo/request"
	"repo/response"
	"repo/util"
	"time"
)

// UserService ...
type PlayersService interface {
	Login(req request.Login) (res response.Login, err error)
	Logout(jwt string) (res string, err error)
	Register(req request.Player) (res string, err error)
	Profile(id string) (res response.Player, err error)
	Update(id string, req request.Player) (res string, err error)
	GetAll(filter map[string]query.Filter, limit, skip int, havingDeposit, orderby string) (response.ListPlayer, error)
	PlayerDeposit(req request.PlayerDeposit) (res string, err error)
	Filter() map[string]query.Filter
}

type playersService struct {
	player  repositories.PlayerRepository
	deposit repositories.PlayerDepositRepository
}

func (s *playersService) Register(req request.Player) (res string, err error) {

	player, err := s.player.GetOne("email = ?", req.Email)
	if err != nil {
		return res, util.NewError("400")
	}

	if player.Username != "" {
		return "player existing", util.NewError("-1000")
	}

	hash, err := helper.HashPassword(req.Password)
	if err != nil {
		return "failed hash password", util.NewError("500")
	}

	if err = s.player.Create(models.Players{
		Username:  req.Username,
		Password:  hash,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		return "failed register player", util.NewError("500")
	}

	return "success register player", nil
}

func (s *playersService) Login(req request.Login) (res response.Login, err error) {

	player, err := s.player.GetOne("email = ?", req.Email)
	if err != nil {
		return res, util.NewError("400")
	}

	if !helper.CheckPasswordHash(req.Password, player.Password) {
		return res, util.NewError("498")
	}

	header, err := s.authTokenGenerator(player.ID.String())
	if err != nil {
		return res, util.NewError("500")
	}

	err = redis.Set(header["x-token"], player.ID.String(), 0)
	if err != nil {
		return res, util.NewError("500")
	}

	debit, err := s.deposit.GetOne("player_id = ? AND type = ?", player.ID, flags.DepositDebit)
	if err != nil {
		return res, util.NewError("400")
	}
	credit, err := s.deposit.GetOne("player_id = ? AND type = ?", player.ID, flags.DepositCredit)
	if err != nil {
		return res, util.NewError("400")
	}

	return response.Login{
		Header: header,
		Player: response.Player{
			ID:           player.ID.String(),
			Username:     player.Username,
			Email:        player.Email,
			Bank:         player.Bank,
			NamaRekening: player.NamaRekening,
			NoRekening:   player.NoRekening,
			Debit:        debit,
			Credit:       credit,
			Deposit:      debit - credit,
			CreatedAt:    player.CreatedAt,
			UpdatedAt:    player.UpdatedAt,
		},
	}, nil
}

func (s *playersService) Logout(jwt string) (res string, err error) {
	if _, err := redis.Del(jwt); err != nil {
		return "failed logout", util.NewError("500")
	}

	return "success logout", nil
}

func (s *playersService) Profile(id string) (res response.Player, err error) {
	player, err := s.player.GetOne("id = ?", id)
	if err != nil {
		return res, util.NewError("400")
	}

	debit, err := s.deposit.GetOne("player_id = ? AND type = ?", id, flags.DepositDebit)
	if err != nil {
		return res, util.NewError("400")
	}
	credit, err := s.deposit.GetOne("player_id = ? AND type = ?", id, flags.DepositCredit)
	if err != nil {
		return res, util.NewError("400")
	}

	return response.Player{
		ID:           player.ID.String(),
		Username:     player.Username,
		Email:        player.Email,
		Bank:         player.Bank,
		NamaRekening: player.NamaRekening,
		NoRekening:   player.NoRekening,
		Debit:        debit,
		Credit:       credit,
		Deposit:      debit - credit,
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}, nil
}

func (s *playersService) Update(id string, req request.Player) (res string, err error) {

	if err = s.player.Update(models.Players{
		Username:     req.Username,
		Email:        req.Email,
		Bank:         req.Bank,
		NamaRekening: req.NamaRekening,
		NoRekening:   req.NoRekening,
		UpdatedAt:    time.Now(),
	}, "id = ?", id); err != nil {
		return "failed update player", util.NewError("500")
	}

	return "success update player", nil
}

func (s *playersService) authTokenGenerator(userId string) (map[string]string, error) {
	// Initiate token
	token, _, expiredAt, err := crypto.NewJWT(userId, flags.ACLAuthenticatedUser)
	if err != nil {
		return nil, err
	}

	// Generate token header
	return response.NewToken(token, expiredAt), nil
}

func (s *playersService) GetAll(filter map[string]query.Filter, limit, skip int, havingDeposit, orderby string) (response.ListPlayer, error) {
	var list []response.Player
	data, count, _ := s.player.GetAll(filter, limit, skip, havingDeposit, orderby)

	for _, player := range data {
		list = append(list, response.Player{
			ID:           player.ID.String(),
			Username:     player.Username,
			Email:        player.Email,
			Bank:         player.Bank,
			NamaRekening: player.NamaRekening,
			NoRekening:   player.NoRekening,
			Debit:        player.Debit,
			Credit:       player.Credit,
			Deposit:      player.Deposit,
			CreatedAt:    player.CreatedAt,
			UpdatedAt:    player.UpdatedAt,
		})
	}

	return response.ListPlayer{
		Meta: response.Meta{
			Total: int(count),
			Limit: int8(limit),
			Skip:  int64(skip),
		},
		Player: list,
	}, nil
}

func (s *playersService) PlayerDeposit(req request.PlayerDeposit) (res string, err error) {

	if err = s.deposit.Create(models.PlayersDeposit{
		PlayerID:  req.PlayerID,
		Type:      req.Type,
		Nominal:   req.Nominal,
		CreatedAt: time.Now(),
	}); err != nil {
		return fmt.Sprintf("failed %s player", req.Type), util.NewError("500")
	}

	return fmt.Sprintf("success %s player", req.Type), nil
}

func (s *playersService) Filter() map[string]query.Filter {
	// Init filter from repository
	return s.player.Filter()
}
