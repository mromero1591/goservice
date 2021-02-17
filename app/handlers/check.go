package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type check struct {
	log *log.Logger
}

//func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
func (c check) readiness(w http.ResponseWriter, r *http.Request) {

	status := struct {
		Status string
	}{
		Status: "ok",
	}
	log.Println(r, status)

	json.NewEncoder(w).Encode(status)

}
