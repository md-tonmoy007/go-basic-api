package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/md-tonmoy007/go-basic-api/api"
	"github.com/md-tonmoy007/go-basic-api/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)




func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder  *schema.Decoder = schema.NewDecoder()
	var err error


	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error("Failed to decode query parameters: ", err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error("Failed to connect to database: ", err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoinDetails(params.Username)
	if tokenDetails == nil {
		log.Error("No coin details found for user: ", params.Username)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: (*&tokenDetails).Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Failed to encode response: ", err)
		api.InternalErrorHandler(w)
		return
	}

}