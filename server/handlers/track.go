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

const uploadDir = "files/"

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

    err := r.ParseMultipartForm(10 << 20) // Max upload size set to 10MB
    if err != nil {
        http.Error(w, "Error parsing form data", http.StatusBadRequest)
        return
    }

    // Extract form values
    var createTrackDTO models.Track
    createTrackDTO.Name = r.FormValue("name")
    createTrackDTO.Artist = r.FormValue("artist")
    createTrackDTO.Text = r.FormValue("text")
    pictureFile, pictureHeader, err := r.FormFile("picture")
    if err == nil {
        createTrackDTO.Picture, err = services.SaveFile(pictureFile, pictureHeader, uploadDir)
        if err != nil {
            fmt.Println(err)
            http.Error(w, "Error saving picture file", http.StatusInternalServerError)
            return
        }
    } else {
        createTrackDTO.Picture = ""
    }

    audioFile, audioHeader, err := r.FormFile("audio")
    if err == nil {
        createTrackDTO.Audio, err = services.SaveFile(audioFile, audioHeader, uploadDir)
        if err != nil {
            http.Error(w, "Error saving audio file", http.StatusInternalServerError)
            return
        }
    } else {
        createTrackDTO.Audio = ""
    }

    createTrackDTO.Listens = 0
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

    err = services.DeleteFile(deletedTrack.Audio, uploadDir)
    if err != nil {
        http.Error(w, "Error deleting audio", http.StatusInternalServerError)
    }

    err = services.DeleteFile(deletedTrack.Picture, uploadDir)
    if err != nil {
        http.Error(w, "Error deleting audio", http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(deletedTrack)

    return
}
