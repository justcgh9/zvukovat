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

func PostPlaylist(w http.ResponseWriter, r *http.Request) {

    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Error parsing form data", http.StatusBadRequest)
        return
    }

    var createPlaylistDTO models.Playlist
    createPlaylistDTO.Tracks = make([]string, 0)
    createPlaylistDTO.Name = r.FormValue("name")
    //CHANGEME
    createPlaylistDTO.Owner = r.FormValue("owner")
    pictureFile, pictureHeader, err := r.FormFile("picture")
    if err == nil {
        createPlaylistDTO.Picture, err = services.SaveFile(pictureFile, pictureHeader, uploadDir, "picture")
        if err != nil {
            http.Error(w, "Error saving picture file", http.StatusInternalServerError)
            return
        }
    } else {
        createPlaylistDTO.Picture = ""
    }

    track, err := services.CreatePlaylist(createPlaylistDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(track)


}

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)
	user := r.Context().Value("user").(*models.UserClaims)
    var playlistID string
    playlistID = mux.Vars(r)["playlist_id"]

    playlist, err := services.GetPlaylist(playlistID)
    if err != nil {
        switch {
        case errors.Is(err, mongo.ErrNoDocuments):
            http.Error(w, "No Playlist with this id", http.StatusNotFound)
        default:
            http.Error(w, "Error fetching track", http.StatusInternalServerError)
        }
        return

    }

    if playlist.IsPrivate && user.Payload.Id != playlist.Owner {
        http.Error(w, "You don't have access to this playlist", http.StatusForbidden)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(playlist)

    return

}

func GetMyPlaylists(w http.ResponseWriter, r *http.Request) {

    AllowOrigin(w)
    user := r.Context().Value("user").(*models.UserClaims)
    playlists, err := services.GetMyPlaylists(user.Payload.Id)
    if err != nil {
        http.Error(w, "Could not get playlists", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(playlists)
}

func GetPublicPlaylists(w http.ResponseWriter, r *http.Request) {

    AllowOrigin(w)
    playlists, err := services.GetPublicPlaylists()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(playlists)

}

func PostToPlaylist(w http.ResponseWriter, r *http.Request) {

    AllowOrigin(w)
    var playlistID string
    playlistID = mux.Vars(r)["playlist_id"]
	user := r.Context().Value("user").(*models.UserClaims)

    var addTrackDTO models.AddToPlaylistDTO
    err := json.NewDecoder(r.Body).Decode(&addTrackDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

    err = services.AddTrackToPlaylist(playlistID, addTrackDTO.TrackId, user.Payload.Id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


}

func DeletePlaylist(w http.ResponseWriter, r *http.Request) {

    AllowOrigin(w)
    var playlistID string
    playlistID = mux.Vars(r)["playlist_id"]
	user := r.Context().Value("user").(*models.UserClaims)

    err := services.DeletePlaylist(playlistID, user.Payload.Id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
    }

    fmt.Fprintf(w, "Playlist with id %s deleted successfully", playlistID)
}

func ToggleVisibility(w http.ResponseWriter, r *http.Request) {
    AllowOrigin(w)

	user := r.Context().Value("user").(*models.UserClaims)
    var playlistID string
    playlistID = mux.Vars(r)["playlist_id"]

    err := services.FlipVisibility(playlistID, user.Payload.Id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    fmt.Fprintf(w, "The visibility of playlist %s was changed successfully", playlistID)
}
