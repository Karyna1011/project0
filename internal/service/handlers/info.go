package handlers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/ape"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	message := "It's our database"
	if err := json.NewEncoder(w).Encode(message); err != nil {
		Log(r).WithError(err).Error("error during writing message")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ape.Render(w, message)
}
