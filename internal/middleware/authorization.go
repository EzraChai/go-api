package middleware

import (
	"errors"
	"net/http"

	"github.com/EzraChai/goapi/api"
	"github.com/EzraChai/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("error: invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.URL.Query().Get("token")

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var err error
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetLoginDetails(username)

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
