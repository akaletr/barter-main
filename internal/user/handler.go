package user

import (
	"cmd/main/main.go/internal/handlers"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (handler *handler) Register(router *chi.Mux) {
	router.Get("/users/get", handler.getUser)
	router.Post("/users/create", handler.createUser)
}

func (handler *handler) getUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All users"))
}

func (handler *handler) createUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var userCandidate User
	err = json.Unmarshal(body, &userCandidate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(body)
	fmt.Println(userCandidate)

	w.Write([]byte(userCandidate.UserName))
}
