package user

import (
	"encoding/json"
	"go/types"
	"net/http"

	"gihtub.com/go-backend/go-ecommerce/repotypes"
	"github.com/gorilla/mux"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}


func (h *Handler) handleLogin(w http.ResponseWriter,r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter,r *http.Request) {

     var payload *repotypes.RegisterUserPayload

	 if r.Body == nil {
		
	 }

	 err := json.NewDecoder(r.Body).Decode(&payload);
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login",h.handleLogin).Methods("POST")
	router.HandleFunc("/register",h.handleRegister).Methods("POST")
}