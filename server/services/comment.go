package services

import (
	"fmt"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
)

func CreateComment(trackID string, comment models.Comment) (models.Comment, error) {

    track, err := GetOneTrack(trackID)
    if err != nil {
        return models.Comment{}, err
    }

    createdComment, err := repositories.CreateComment(comment)
    if err != nil {
        return models.Comment{}, err
    }

    track.Comments = append(track.Comments, createdComment.Id)
    _, err = UpdateTrack(track)
    if err != nil {
        fmt.Println(err)
    }
    return createdComment, nil
}
