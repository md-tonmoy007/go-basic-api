package middleware

import (
	"errors"
	"net/http"

	"github.com/md-tonmoy007/go-basic-api/api"
	"github.com/md-tonmoy007/go-basic-api/internal/tools"
	log "github.com/sirupsen/logrus"

)



var UnauthorizedError = errors.New("Invalid username or token. ")


func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("Authorization")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}


		var loginDetails *tools.LoginDetails
		
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}
	})
}