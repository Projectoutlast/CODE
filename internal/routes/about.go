package routes

import (
	"fmt"
	"net/http"
)

func ourHistory(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Our History, URL: %s", r.URL.Path[1:])
	if err != nil {
		return
	}
}

func team(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Team, URL: %s", r.URL.Path[1:])
	if err != nil {
		return
	}
}

func missionAndValues(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Team, URL: %s", r.URL.Path[1:])
	if err != nil {
		return
	}
}
