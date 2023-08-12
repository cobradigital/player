package middleware

import (
	"net/http"

	"repo/configs"
	"repo/flags"
	"repo/handler"
	"repo/response"
	"repo/services"
	"repo/util"

	"github.com/gorilla/context"
)

// Middleware ...
type Middleware func(handler.RESTFunc, ...string) handler.RESTFunc

func Auth() Middleware {
	// Create middleware
	m := func(next handler.RESTFunc, args ...string) handler.RESTFunc {
		// Define new handler
		h := func(r *http.Request) (*response.Success, error) {
			// Get purpose
			var purpose string
			if len(args) > 0 {
				purpose = args[0]
			}

			if purpose == flags.ACLAuthenticatedAdmin {
				user, pass, ok := r.BasicAuth()
				if !ok {
					return nil, util.NewError("401")
				}
				ok = validateBasicAuth(user, pass)
				if !ok {
					return nil, util.NewError("401")
				}

				return next(r)
			}
			// Call Validate token service
			id, jwt, err := services.Auth.ValidateAccessToken(r, purpose)
			if err != nil {
				return nil, util.NewError("401")
			}

			context.Set(r, "id", id)
			context.Set(r, "jwt", jwt)

			return next(r)
		}
		return h
	}
	// Return middleware
	return m
}

func validateBasicAuth(user, pass string) bool {

	if user != configs.MustGetString("admin.user") || pass != configs.MustGetString("admin.pass") {
		return false
	}

	return true
}
