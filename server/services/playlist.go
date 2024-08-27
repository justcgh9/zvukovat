package services

import (
	"errors"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
)

func CreatePlaylist(playlist models.Playlist) (models.Playlist, error) {

	newPlaylist, err := repositories.CreatePlaylist(playlist)
	if err != nil {
		return models.Playlist{}, err
	}

	return newPlaylist, nil
}

func AddTrackToPlaylist(playlistId, trackId, owner string) error {

	_, err := repositories.GetOneTrack(trackId)
	if err != nil {
		return errors.New("failure. Ensure the validity of track identifier")
	}

	playlist, err := repositories.GetPlaylist(playlistId)
	if err != nil {
		return err
	}

    if playlist.Owner != owner {
        return errors.New("You don't have permissions to add tracks to this playlist")
    }

	return repositories.AddTrackToPlaylist(playlist, trackId)

}

func GetPlaylist(playlistId string) (models.Playlist, error) {
	return repositories.GetPlaylist(playlistId)
}

func GetMyPlaylists(userId string) ([]models.Playlist, error) {
   return repositories.GetMyPlaylists(userId)
}

func DeletePlaylist(playlistId, owner string) error {
    playlist, err := GetPlaylist(playlistId)
    if err != nil {
        return err
    }

    if playlist.Owner != owner {
        return errors.New("You don't have permissions to add tracks to this playlist")
    }
	return repositories.RemovePlaylist(playlistId)
}

func FlipVisibility(playlistId, owner string) error {
    playlist, err := GetPlaylist(playlistId)
    if err != nil {
        return err
    }
    if playlist.Owner != owner {
        return errors.New("You don't have permissions to add tracks to this playlist")
    }
    err = repositories.SetPlaylistVisibility(playlistId, !playlist.IsPrivate)
    return err
}
