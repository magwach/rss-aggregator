package main

import (
	"fmt"
	"net/http"

	"github.com/magwach/rss-aggregator/internal/auth"
	"github.com/magwach/rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) GetUserMiddleware(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			RespondWithError(w, 403, string(err.Error()))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			RespondWithError(w, 400, fmt.Sprintf("Error while getting the user: %v", err))
			return
		}

		handler(w, r, user)
	}

}
