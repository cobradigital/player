package request

import (
	"net/http"

	"repo/flags"
)

func JWTSubject(r *http.Request) string {
	return r.Header.Get(flags.HeaderKeyCOBRASubject)
}
