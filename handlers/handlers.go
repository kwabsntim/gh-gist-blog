package handlers

import (
	"AuthGo/models"
	"AuthGo/services"
	"net/http"
)

// ---------------------------Structs----------------------------
type Handler struct {
	RegisterService services.RegisterInterface
	LoginService    services.LoginInterface
	FetchService    services.FetchUsersInterface
}

func NewHandlers(services *ServiceContainer) *Handler {
	return &Handler{
		RegisterService: services.RegisterService,
		LoginService:    services.LoginService,
		FetchService:    services.FetchService,
	}
}

//---------------------------Handler functions----------------------------

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	// Validate HTTP method

	var req models.User
	if !DecodeJSONRequest(w, r, &req) {
		return
	}
	// Call service - let it handle all business logic
	user, err := h.RegisterService.RegisterUser(req.Email, req.Username, req.Password, req.Role)
	if err != nil {
		http.Error(w, "Registration failed: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Create token
	token, err := CreateTokenForUser(user)
	if err != nil {
		http.Error(w, "Token creation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	// Success response
	SendAuthResponse(w, "User created successfully", data, http.StatusCreated)

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req models.User
	//decoding the json request
	if !DecodeJSONRequest(w, r, &req) {
		return
	}

	user, err := h.LoginService.LoginUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := CreateTokenForUser(user)
	if err != nil {
		http.Error(w, "Token creation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	// Success response
	SendAuthResponse(w, "Login successful", data, http.StatusCreated)
}
func (h *Handler) FetchAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.FetchService.FetchAllUsers()
	if err != nil {
		http.Error(w, "Could not fetch users", http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"users": users,
	}
	SendAuthResponse(w, "Fetch successful", data, http.StatusCreated)
}
