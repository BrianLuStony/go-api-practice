package handlers

import (
	"encoding/json"
	"example/api/internal/http/requests"
	"example/api/internal/http/responses"
	"example/api/internal/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) GetUserWithId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := handler.service.GetUserByID(id)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := responses.NewUserResponse(user.ID, user.Email, "")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (handler *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req requests.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			log.Printf("Validation error for field '%s', tag '%s' ", fieldErr.Field(), fieldErr.Tag())
		}

		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}
	user, token, err := handler.service.Login(req.Email, req.Password)
	if err != nil {
		log.Printf("Login error: %v", err)
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}
	response := responses.NewUserResponse(user.ID, user.Email, token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the GetUser handler"))
}
