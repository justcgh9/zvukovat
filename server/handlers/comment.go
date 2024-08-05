package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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
    var trackID, commentID string
    trackID = mux.Vars(r)["track_id"]
    commentID = mux.Vars(r)["comment_id"]
    var comment models.Comment
    err := json.NewDecoder(r.Body).Decode(&comment)
    if err != nil {
        http.Error(w, "Could not process comment", http.StatusUnprocessableEntity)
        return
    }

    if comment.Id == "" {
        comment.Id = commentID
    }

    if comment.Id != commentID || comment.Track_id != trackID {
        fmt.Printf("comment.Id %s, commentID %s\ncomment.Track_id %s, trackID %s\n", comment.Id, commentID, comment.Track_id, trackID)
        http.Error(w, "Incorrect track or comment identifier", http.StatusUnprocessableEntity)
        return
    }

    newComment, err := services.UpdateComment(comment)
    if err != nil {
        http.Error(w, "Error updating comment", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newComment)
    return
}

func GetComments(w http.ResponseWriter, r *http.Request) {

    var trackID string
    trackID = mux.Vars(r)["track_id"]
    comments, err := services.GetComments(trackID)
    if err != nil {
        switch {
        case errors.Is(err, mongo.ErrNoDocuments):
            http.Error(w, "The track with given ID is not found", http.StatusNotFound)
        default:
            http.Error(w, "Error fetching comment", http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comments)
    return
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {

    var trackID, commentID string
    trackID = mux.Vars(r)["track_id"]
    commentID = mux.Vars(r)["comment_id"]
    comment, err := services.DeleteComment(commentID, trackID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comment)

    return
}

