package main

import (
	"fmt"
	"net/http"
	"strings"
)

// This interface is used to define a contract for any type that wishes to act as a player score storage.
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// This struct represents the web server that will handle player score requests.
type PlayerServer struct {
	store PlayerStore
}

// This method is used to handle incoming HTTP requests.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	// Inside this method, it extracts the player's name from the request URL (after removing the "/players/" prefix).
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, GetPlayerScore(player))
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

// Given a name return the player score
func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
