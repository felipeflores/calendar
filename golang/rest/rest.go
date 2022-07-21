package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SendJSON(w http.ResponseWriter, payload interface{}) error {
	return WriteJSON(context.Background(), w, payload)
}

func WriteJSON(ctx context.Context, w http.ResponseWriter, payload interface{}, httpStatusCode ...int) error {
	w.Header().Set("Content-Type", "application/json")

	if len(httpStatusCode) > 0 {
		w.WriteHeader(httpStatusCode[0])
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return errors.New("errou")
	}
	return nil
}

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet)
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("favicon.ico"))
	}).Methods(http.MethodGet)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("barra"))
	}).Methods(http.MethodGet)
}

func GetInt(r *http.Request, param string) (int, error) {
	p, ok := mux.Vars(r)[param]
	if !ok {
		return 0, errors.New("parameter not found")
	}
	return strconv.Atoi(p)
}

func GetStr(r *http.Request, param string) (string, error) {
	p, ok := mux.Vars(r)[param]
	if !ok {
		return "", errors.New("parameter not found")
	}
	return p, nil
}
