package routers

import (
	"encoding/json"
	"fmt"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
	"justcgh9/spotify_clone/server/services"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func PostSignUp(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Could not process user data", http.StatusUnprocessableEntity)
		return
	}

	tokens, user, err := services.Registration(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = ""

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    tokens["refreshToken"],
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	}

	w.Header().Add("Set-Cookie", fmt.Sprintf("%s;Partitioned", cookie.String()))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := make(map[string]interface{})
	response["tokens"] = tokens
	response["user"] = user
	json.NewEncoder(w).Encode(response)

	return
}

func PostSignIn(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userData, err := services.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    userData["refreshToken"],
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	}

	w.Header().Add("Set-Cookie", fmt.Sprintf("%s;Partitioned", cookie.String()))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
	fmt.Println(userData)

	return
}
func PostSignOut(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	user := r.Context().Value("user").(*models.UserClaims)
    err := services.Logout(user.Payload.Id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
        MaxAge: -1,
	}
	w.Header().Add("Set-Cookie", fmt.Sprintf("%s;Partitioned", cookie.String()))
	return
}
func GetActivation(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	activationLink := mux.Vars(r)["link"]
	user, err := repositories.ActivateUser(activationLink)
	fmt.Println(activationLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	Tmpl.Render(w, "activate", user)

	return
}
func GetRefreshedToken(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	refreshCookie, err := r.Cookie("refreshToken")
	if err != nil {
		http.Error(w, "Ivalid refreshToken cookie", http.StatusForbidden)
		return
	}

	userData, err := services.Refresh(*refreshCookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    userData["refreshToken"],
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	}

	w.Header().Add("Set-Cookie", fmt.Sprintf("%s;Partitioned", cookie.String()))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
	fmt.Println(userData)

	return
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	users, err := repositories.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
	return
}
func GetUser(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	user_id := mux.Vars(r)["user_id"]
	user, err := repositories.GetUser(user_id)
	fmt.Println(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	user := r.Context().Value("user").(*models.UserClaims)
	fmt.Fprintf(w, "Hello, %s!", user.Payload.Email)
}
