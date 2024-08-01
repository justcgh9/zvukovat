package handlers

import (
	"encoding/json"
	"errors"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/services"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {

    var trackID string
    trackID = mux.Vars(r)["track_id"]

    var comment models.Comment
    err := json.NewDecoder(r.Body).Decode(&comment)
    if err != nil {
        http.Error(w, "Could not process comment", http.StatusUnprocessableEntity)
        return
    }
    newComment, err := services.CreateComment(trackID, comment)
    if err != nil {
        switch {
        case errors.Is(err, mongo.ErrNoDocuments):
            http.Error(w, "No Track with this id", http.StatusNotFound)
        default:
            http.Error(w, "Error Creating Comment", http.StatusInternalServerError)
        }
        return

    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newComment)

    return
}

func EditComment(w http.ResponseWriter, r *http.Request) {
    return
}

func GetComments(w http.ResponseWriter, r *http.Request) {
    return
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
    return
}
