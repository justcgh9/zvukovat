package routers

import (
	"encoding/json"
	"fmt"
	"justcgh9/spotify_clone/server/models"
	"net/http"
)

func PostSignUp(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Could not process user data", http.StatusUnprocessableEntity)
    }
    fmt.Println(user)
    return
}
func PostSignIn(w http.ResponseWriter, r *http.Request) {
    return
}
func PostSignOut(w http.ResponseWriter, r *http.Request) {
    return
}
func GetActivation(w http.ResponseWriter, r *http.Request) {
    return
}
func GetRefreshedToken(w http.ResponseWriter, r *http.Request) {
    return
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
    return
}
