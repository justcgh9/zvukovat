package routers

import (
	"encoding/json"
	"errors"
	"fmt"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

const uploadDir = "files/"

func AllowOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func HandleCORS(w http.ResponseWriter, r *http.Request) {

	AllowOrigin(w)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-type, Cookie, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}

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

	queryParams := r.URL.Query()
	countStr := queryParams.Get("count")
	offsetStr := queryParams.Get("offset")

	if (countStr == "") != (offsetStr == "") {
		http.Error(w, "Either both count and offset must be set or none of them", http.StatusBadRequest)
		return
	}

	var getParams *models.TrackPaginationParams
	if countStr != "" {
		count, err := strconv.Atoi(countStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		getParams = &models.TrackPaginationParams{
			Offset: offset,
			Count:  count,
		}
	}
	tracks, err := services.GetAllTracks(getParams)
	if err != nil {
		http.Error(w, "Error fetching tracks", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
	return
}

func SearchTrack(w http.ResponseWriter, r *http.Request) {
	AllowOrigin(w)
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	artist := queryParams.Get("artist")
	tracks, err := services.FindTrack(name, artist)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error fetching tracks", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
	return
}

func PostTrack(w http.ResponseWriter, r *http.Request) {
	AllowOrigin(w)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	var createTrackDTO models.Track
	createTrackDTO.Name = r.FormValue("name")
	createTrackDTO.Artist = r.FormValue("artist")
	createTrackDTO.Text = r.FormValue("text")
	pictureFile, pictureHeader, err := r.FormFile("picture")
	if err == nil {
		createTrackDTO.Picture, err = services.SaveFile(pictureFile, pictureHeader, uploadDir, "picture")
		if err != nil {
			http.Error(w, "Error saving picture file", http.StatusInternalServerError)
			return
		}
	} else {
		createTrackDTO.Picture = ""
	}

	audioFile, audioHeader, err := r.FormFile("audio")
	if err == nil {
		createTrackDTO.Audio, err = services.SaveFile(audioFile, audioHeader, uploadDir, "audio")
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

func LikeTrack(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.UserClaims)
	trackId := mux.Vars(r)["track_id"]
	fmt.Println(user.Payload.Email, trackId)
	trackIds, err := services.LikeTrack(user.Payload.Email, trackId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trackIds)
}

func UnlikeTrack(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.UserClaims)
	trackId := mux.Vars(r)["track_id"]
	trackIds, err := services.UnlikeTrack(user.Payload.Email, trackId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trackIds)
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := services.GetArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}
