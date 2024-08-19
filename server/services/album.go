package services

import (
	"errors"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
)

func CreateAlbum(album models.Album) (models.Album, error) {

	newAlbum, err := repositories.CreateAlbum(album)
	if err != nil {
		return models.Album{}, err
	}

	return newAlbum, nil
}

func AddTrackToAlbum(albumId, trackId string) error {

	_, err := repositories.GetOneTrack(trackId)
	if err != nil {
		return errors.New("failure. Ensure the validity of track identifier")
	}

	album, err := repositories.GetAlbum(albumId)
	if err != nil {
		return err
	}

	return repositories.AddTrackToAlbum(album, trackId)

}

func GetAlbum(albumId string) (models.Album, error) {
	return repositories.GetAlbum(albumId)
}

func DeleteAlbum(albumId string) error {
	return repositories.RemoveAlbum(albumId)
}
