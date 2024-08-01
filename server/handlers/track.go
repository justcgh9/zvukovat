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

func GetTrackHandler(w http.ResponseWriter, r *http.Request) {

    var trackID string
    trackID = mux.Vars(r)["track_id"]

    track, err := services.GetOneTrack(trackID)
    if err != nil {
        switch {
        case errors.Is(err, mongo.ErrNoDocuments):
            http.Error(w, "No Track with this id", http.StatusNotFound)
        default:
            http.Error(w, "Error fetching track", http.StatusInternalServerError)
        }
        return

    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(track)

    return
}

func GetTracksHandler(w http.ResponseWriter, r *http.Request) {

    tracks, err := services.GetAllTracks()
    if err != nil {
        http.Error(w, "Error fetching tracks", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tracks)
    return
}

func PostTrack(w http.ResponseWriter, r *http.Request) {

    var createTrackDTO models.Track
    err := json.NewDecoder(r.Body).Decode(&createTrackDTO)
    if err != nil {
        http.Error(w, "Invalid Request Payload", http.StatusUnprocessableEntity)
        return
    }

    if createTrackDTO.Comments == nil {
        createTrackDTO.Comments = make([]string, 1)
    }

    newTrack, err := services.CreateTrack(createTrackDTO)
    if err != nil {
        http.Error(w, "Error when creating track", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newTrack)

    return
}

func DeleteTrack(w http.ResponseWriter, r *http.Request) {

    var trackID string
    trackID = mux.Vars(r)["track_id"]
    deletedTrack, err := services.DeleteTrack(trackID)
    if err != nil {
        switch {
        case errors.Is(err, mongo.ErrNoDocuments):
            http.Error(w, "No Track with this id", http.StatusNotFound)
        default:
            http.Error(w, "Error fetching track", http.StatusInternalServerError)
    }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(deletedTrack)

    return
}
