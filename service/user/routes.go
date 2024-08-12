package user

import (
	"net/http"

	"github.com/edbertjk/server-api/types"
	"github.com/edbertjk/server-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//check if the user exists
	//we doesnt make user

}
