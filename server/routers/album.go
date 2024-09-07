package routers

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

func PostAlbum(w http.ResponseWriter, r *http.Request) {

    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Error parsing form data", http.StatusBadRequest)
        return
    }

    var createAlbumDTO models.Album
    createAlbumDTO.Tracks = make([]string, 0)
    createAlbumDTO.Name = r.FormValue("name")
    createAlbumDTO.Artist = r.FormValue("artist")
    pictureFile, pictureHeader, err := r.FormFile("picture")
    if err == nil {
        createAlbumDTO.Picture, err = services.SaveFile(pictureFile, pictureHeader, uploadDir, "picture")
        if err != nil {
            http.Error(w, "Error saving picture file", http.StatusInternalServerError)
            return
        }
    } else {
        createAlbumDTO.Picture = ""
    }

    track, err := services.CreateAlbum(createAlbumDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(track)


}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
    var albumID string
    albumID = mux.Vars(r)["album_id"]

    album, err := services.GetAlbum(albumID)
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
    json.NewEncoder(w).Encode(album)

    return

}

func PostToAlbum(w http.ResponseWriter, r *http.Request) {

    AllowOrigin(w)
    var albumID string
    albumID = mux.Vars(r)["album_id"]

    var addTrackDTO models.AddToAlbumDTO
    err := json.NewDecoder(r.Body).Decode(&addTrackDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

    err = services.AddTrackToAlbum(albumID, addTrackDTO.TrackId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {

    AllowOrigin(w)
    var albumID string
    albumID = mux.Vars(r)["album_id"]

    album, err := services.GetAlbum(albumID)

    err = services.DeleteAlbum(albumID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    err = services.DeleteFile(album.Picture, uploadDir)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Album with id %s deleted successfully", albumID)
}
