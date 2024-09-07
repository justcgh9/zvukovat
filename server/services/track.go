package services

import (
	"fmt"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
)

func CreateTrack(dto models.Track) (models.Track, error) {
	return repositories.AddTrack(dto)
}

func UpdateTrack(dto models.Track) (models.Track, error) {
	return repositories.UpdateTrack(dto)
}

func GetAllTracks(params *models.TrackPaginationParams) ([]models.Track, error) {
	return repositories.GetAllTracks(params)
}

func FindTrack(name, artist string) ([]models.Track, error) {
	return repositories.SearchTrack(name, artist)
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

func LikeTrack(email, trackId string) ([]string, error) {
	fmt.Println("I")
	user, err := repositories.GetUser(email)
	if err != nil {
		return nil, err
	}
	fmt.Println("I")
	_, err = repositories.GetOneTrack(trackId)
	if err != nil {
		return nil, err
	}
	fmt.Println("I")
	if user.ContainsTrack(trackId) {
		return user.FavouriteTracks, nil
	}
	fmt.Println("I")
	user.FavouriteTracks = append(user.FavouriteTracks, trackId)
	err = repositories.UpdateFavourites(user)
	return user.FavouriteTracks, err
}

func UnlikeTrack(email, trackId string) ([]string, error) {
	user, err := repositories.GetUser(email)
	if err != nil {
		return nil, err
	}

	_, err = repositories.GetOneTrack(trackId)
	if err != nil {
		return nil, err
	}

	buff := []string{}

	for _, track := range user.FavouriteTracks {
		if track != trackId {
			buff = append(buff, track)
		}
	}

	user.FavouriteTracks = buff
	err = repositories.UpdateFavourites(user)
	return user.FavouriteTracks, err
}

func GetArtists() ([]string, error) {
    return repositories.GetArtists()
}
