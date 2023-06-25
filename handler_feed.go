package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nanashi10211/rssaggregator/internal/database"
)


func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string	`json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feed, db_err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url: 	   params.URL,
		UserID: user.ID,
	})
	if db_err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn;t create a feed: %s", db_err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}



func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	

	feeds, db_err := apiCfg.DB.GetFeeds(r.Context())
	if db_err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed: %s", db_err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}

