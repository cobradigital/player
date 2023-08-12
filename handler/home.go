package handler

import (
	"net/http"

	"repo/configs"
	"repo/response"
)

// Health ...
func Health(r *http.Request) (*response.Success, error) {
	result := make(map[string]interface{}, 1)
	result["version"] = configs.MustGetString("server.version")
	return response.NewSuccess(result, nil), nil
}
