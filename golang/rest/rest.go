package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"iot/pkg/ferrors"
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

func DeserializeJSON(r *http.Request, payload interface{}) error {
	return ReadJSON(context.Background(), r, payload)
}

// ReadJSON decode JSON from body to payload
func ReadJSON(ctx context.Context, r *http.Request, payload interface{}) error {
	fmt.Println(fmt.Printf("BODY: %s", r.Body))

	// b, err := io.ReadAll(r.Body)
	b, err := ioutil.ReadAll(r.Body) //Go.1.15 and earlier
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	err = json.Unmarshal(b, payload)
	if err != nil {
		fmt.Println("passou aqui e errou")
		return ferrors.NewBadRequest(err)
	}
	// fmt.Println(fmt.Printf("sera q foi %v", payload))
	// err = json.NewDecoder(r.Body).Decode(payload)
	// if err != nil {
	// 	return ferrors.NewBadRequest(err)
	// }

	return nil
}
