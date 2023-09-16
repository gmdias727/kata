package main

import (
	"fmt"
	"net/http"
	"strings"
)

// This interface is used to define a contract for any type that wishes to act as a player score storage.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// This struct represents the web server that will handle player score requests.
type PlayerServer struct {
	store PlayerStore
}

// This method is used to handle incoming HTTP requests.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Inside this method, it extracts the player's name from the request URL (after removing the "/players/" prefix).
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
