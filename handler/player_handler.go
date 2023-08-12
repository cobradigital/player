package handler

import (
	"fmt"
	"net/http"
	"repo/request"
	"repo/response"
	"repo/services"
	"repo/util"

	"github.com/gorilla/context"
)

// Register ...
func GetAll(r *http.Request) (*response.Success, error) {

	q := r.URL.Query()

	limit, skip := Pagination(q)

	filters := Filter(q, services.Players.Filter())

	orderby := q.Get("order_by")
	havingDeposit := q.Get("deposit")

	result, err := services.Players.GetAll(filters, limit, skip, havingDeposit, orderby)
	if err != nil {
		return nil, err
	}

	return &response.Success{
		Result: result.Player,
		Metadata: &response.Metadata{
			Count: result.Meta.Total,
			Limit: result.Meta.Limit,
			Skip:  result.Meta.Skip,
		},
	}, nil
}

// Register ...
func Register(r *http.Request) (*response.Success, error) {

	var req request.Player
	err := parseJSON(r, &req)
	if err != nil {
		return nil, util.NewError("400")
	}

	result, err := services.Players.Register(req)
	if err != nil {
		return nil, err
	}

	return &response.Success{
		Result: result,
	}, nil
}

// Login ...
func Login(r *http.Request) (*response.Success, error) {

	var req request.Login
	err := parseJSON(r, &req)
	if err != nil {
		return nil, util.NewError("400")
	}

	result, err := services.Players.Login(req)
	if err != nil {
		return nil, err
	}

	return &response.Success{
		Result: result.Player,
		Header: result.Header,
	}, nil
}

// Login ...
func Logout(r *http.Request) (*response.Success, error) {

	jwt := fmt.Sprintf("%v", context.Get(r, "jwt"))
	result, err := services.Players.Logout(jwt)
	if err != nil {
		return nil, err
	}

	return &response.Success{
		Result: result,
	}, nil
}

// GetProfile ...
func GetProfile(r *http.Request) (*response.Success, error) {

	id := fmt.Sprintf("%v", context.Get(r, "id"))

	result, err := services.Players.Profile(id)
	if err != nil {
		return nil, err
	}

	return &response.Success{
		Result: result,
	}, nil
}

// PutProfile ...
func PutProfile(r *http.Request) (*response.Success, error) {

	id := fmt.Sprintf("%v", context.Get(r, "id"))

	var req request.Player
	err := parseJSON(r, &req)
	if err != nil {
		return nil, util.NewError("400")
	}

	result, err := services.Players.Update(id, req)
	if err != nil {
		return nil, err
	}

	return &response.Success{
		Result: result,
	}, nil
}

// PostPlayerDeposit ...
func PostPlayerDeposit(r *http.Request) (*response.Success, error) {

	var req request.PlayerDeposit
	err := parseJSON(r, &req)
	if err != nil {
		return nil, util.NewError("400")
	}

	result, err := services.Players.PlayerDeposit(req)
	if err != nil {
		return nil, err
	}

	return &response.Success{
		Result: result,
	}, nil
}
