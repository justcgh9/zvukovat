package services

import (
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
)

func CreateTrack(dto models.Track) (models.Track, error) {
    return repositories.AddTrack(dto)
}

func GetAllTracks() ([]models.Track, error) {
    return repositories.GetAllTracks()
}

func GetOneTrack(trackID string) (models.Track, error) {
    return repositories.GetOneTrack(trackID)
}

func DeleteTrack(trackID string) (models.Track, error) {
    track, err := repositories.GetOneTrack(trackID)
    if err != nil {
        return track, err
    }

    err = repositories.DeleteTrack(trackID)
    if err != nil {
        return models.Track{}, err
    }

    return track, nil
}
